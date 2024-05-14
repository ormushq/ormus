package rbbitmqchannel

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/ormushq/ormus/destination/dconfig"
	"github.com/ormushq/ormus/pkg/channel"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type testCase struct {
	name           string
	numWorkers     int
	numMessages    int
	expectedMsg    int
	receiveTimeout time.Duration
}
type message struct {
	WorkerId  int `json:"worker_id"`
	MessageId int `json:"message_id"`
}

func setup(t testing.TB) {
	cli, err := client.NewClientWithOpts()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx,
		&container.Config{
			Image:        "rabbitmq:3-management-alpine",
			ExposedPorts: nat.PortSet{"5672": struct{}{}, "15672": struct{}{}},
		},
		&container.HostConfig{
			PortBindings: map[nat.Port][]nat.PortBinding{
				nat.Port("5672"):  {{HostIP: "127.0.0.1", HostPort: "5672"}},
				nat.Port("15672"): {{HostIP: "127.0.0.1", HostPort: "15672"}},
			},
		},
		nil,
		nil, "rabbitmq-test")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}
	t.Cleanup(func() {
		t.Log("Under cleanup")
		timeout := time.Second
		err = cli.ContainerStop(ctx, resp.ID, &timeout)
		if err != nil {
			panic(err)
		}

		err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
		if err != nil {
			panic(err)
		}
	})
	time.Sleep(time.Second * 30)
}

func TestRabbitmqChannel(t *testing.T) {
	setup(t)
	cases := []testCase{
		{
			name:           "small test",
			numWorkers:     1,
			numMessages:    10,
			expectedMsg:    10,
			receiveTimeout: 30 * time.Second,
		},
		{
			name:           "huge test",
			numWorkers:     10,
			numMessages:    1000,
			expectedMsg:    10000,
			receiveTimeout: 120 * time.Second,
		},
	}
	done := make(chan bool)
	wg := &sync.WaitGroup{}
	config := dconfig.RabbitMQConsumerConnection{
		User:            "guest",
		Password:        "guest",
		Host:            "127.0.0.1",
		Port:            5672,
		Vhost:           "/",
		ReconnectSecond: 1,
	}
	bufferSize := 100
	numberInstants := 10
	maxRetryPolicy := 5

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			workerwg := &sync.WaitGroup{}

			inputAdapter := New(done, wg, config)
			inputAdapter.NewChannel(tc.name, channel.InputOnlyMode, bufferSize, numberInstants, maxRetryPolicy)
			inputChannel, err := inputAdapter.GetInputChannel(tc.name)
			if err != nil {
				panic(err)
			}

			for workerId := 0; workerId < tc.numWorkers; workerId++ {
				workerwg.Add(1)
				go func(wi int, ic chan<- []byte) {
					defer workerwg.Done()
					for messageId := 0; messageId < tc.numMessages; messageId++ {
						msg := message{
							WorkerId:  wi,
							MessageId: messageId,
						}
						m, err := json.Marshal(msg)
						if err != nil {
							panic(err)
						}
						inputChannel <- m
					}
				}(workerId, inputChannel)
			}
			time.Sleep(time.Second * tc.receiveTimeout / 2)
			outputAdapter := New(done, wg, config)
			outputAdapter.NewChannel(tc.name, channel.OutputOnly, bufferSize, numberInstants, maxRetryPolicy)
			outputChannel, err := outputAdapter.GetOutputChannel(tc.name)
			if err != nil {
				panic(err)
			}
			msgReceivedCount := atomic.Int32{}

			receivedMessages := &sync.Map{}

			workerwg.Add(1)
			go func() {
				defer workerwg.Done()
				timeoutChan := time.After(tc.receiveTimeout)
				for {
					select {
					case <-timeoutChan:
						t.Logf("Consume timeout \n")
						return
					case msg := <-outputChannel:
						workerwg.Add(1)
						go func(msg channel.Message) {
							defer workerwg.Done()
							m := message{}
							err := json.Unmarshal(msg.Body, &m)
							if err != nil {
								panic(err)
							}

							if value, ok := receivedMessages.Load(m.WorkerId); ok {
								value.(*sync.Map).Store(m.MessageId, true)
								receivedMessages.Store(m.WorkerId, value)
							} else {
								smap := &sync.Map{}
								smap.Store(m.MessageId, true)
								receivedMessages.Store(m.WorkerId, smap)
							}
							err = msg.Ack()
							if err != nil {
								panic(err)
							}
							msgReceivedCount.Add(1)
						}(msg)
					default:
						time.Sleep(time.Millisecond * 100)
					}
					if int(msgReceivedCount.Load()) == tc.expectedMsg {
						return
					}

				}
			}()
			t.Log("Before final worker wait")
			workerwg.Wait()
			for i := 0; i < tc.numWorkers; i++ {
				for j := 0; j < tc.numWorkers; j++ {
					v, ok := receivedMessages.Load(i)
					if !ok {
						t.Errorf("Worker %d not found", i)
						t.Fail()
					}
					_, ok = v.(*sync.Map).Load(j)
					if !ok {
						t.Errorf("Message from worker %d with message id %d not found", i, j)
						t.Fail()
					}

				}
			}
			if int(msgReceivedCount.Load()) != tc.expectedMsg {
				t.Errorf("Received message count %d but excepted %d", msgReceivedCount.Load(), tc.expectedMsg)
				t.Fail()
			}
			t.Log("Done successfully")

		})
	}
	t.Log("Before close done channel")
	close(done)
	t.Log("Before final wait")
	wg.Wait()
}
