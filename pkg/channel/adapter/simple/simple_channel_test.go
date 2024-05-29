package simple

import (
	"encoding/json"
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

func TestSimpleChannel(t *testing.T) {

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

	bufferSize := 100
	numberInstants := 10
	maxRetryPolicy := 5

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			workerwg := &sync.WaitGroup{}

			adapter := New(done, wg)
			adapter.NewChannel(tc.name, channel.BothMode, bufferSize, numberInstants, maxRetryPolicy)
			inputChannel, err := adapter.GetInputChannel(tc.name)
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

			outputChannel, err := adapter.GetOutputChannel(tc.name)
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
