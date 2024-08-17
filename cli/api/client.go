package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/ormushq/ormus/cli/api/destination"
	"github.com/ormushq/ormus/cli/api/project"
	"github.com/ormushq/ormus/cli/api/source"
	"github.com/ormushq/ormus/cli/api/types"
	"github.com/ormushq/ormus/cli/api/user"
)

const (
	configFIlePath       = "./cli/config.json"
	configFilePermission = 0o644
)

type Client struct {
	User        user.Client
	Destination destination.Client
	Source      source.Client
	Project     project.Client
	config      Config
}

type Config struct {
	Token   string `json:"token"`
	BaseURL string `json:"base_url"`
}

func New() Client {
	client := Client{
		User:        user.New(),
		Destination: destination.New(),
		Source:      source.New(),
		Project:     project.New(),
	}
	client.readConfig()

	return client
}

func (c *Client) StoreToken(token string) {
	c.config.Token = token
	c.storeConfig()
}

func (c *Client) ReadToken() string {
	return c.config.Token
}

func (c *Client) SetConfig(key, value string) error {
	switch key {
	case "token":
		c.config.Token = value
	case "base_url":
		c.config.BaseURL = value
	default:
		return fmt.Errorf("key is invalid %s", key)
	}
	c.storeConfig()

	return nil
}

func (c *Client) GetConfig(key string) (string, error) {
	switch key {
	case "token":
		return c.config.Token, nil
	case "base_url":
		return c.config.BaseURL, nil
	default:
		return "", fmt.Errorf("key is invalid %s", key)
	}
}

func (c *Client) ListConfig() (map[string]string, error) {
	j, err := json.Marshal(c.config)
	if err != nil {
		return nil, err
	}
	var result map[string]string
	err = json.Unmarshal(j, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) getURL(path string) string {
	return fmt.Sprintf("%s/%s", c.config.BaseURL, path)
}

func (c *Client) SendRequest(req types.Request) (*http.Response, error) {
	cl := &http.Client{
		// Timeout: 2000,
	}
	var respBody []byte

	r, err := http.NewRequestWithContext(context.Background(), req.Method, c.getURL(req.Path), bytes.NewBuffer(respBody))
	if err != nil {
		panic(err)
	}
	if req.AuthorizationRequired {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
	}

	return cl.Do(r)
}

func (c *Client) checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !errors.Is(err, os.ErrNotExist)
}

func (c *Client) storeConfig() {
	file, err := os.OpenFile(configFIlePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, configFilePermission)
	if err != nil {
		panic(fmt.Sprintf("can't create or open file, ERR: %s", err))
	}
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			println(err)
		}
	}(file)
	j, err := json.Marshal(c.config)
	if err != nil {
		panic(fmt.Sprintf("can't marshal json %v\n", err))
	}
	_, wErr := file.Write(j)
	if wErr != nil {
		panic(fmt.Sprintf("can't write to the file %v\n", wErr))
	}
}

func (c *Client) initConfig() {
	c.config = Config{
		Token:   "",
		BaseURL: "http://manager.ormus.local",
	}
	c.storeConfig()
}

func (c *Client) readConfig() {
	if !c.checkFileExists(configFIlePath) {
		c.initConfig()

		return
	}
	file, err := os.OpenFile(configFIlePath, os.O_CREATE|os.O_RDONLY, configFilePermission)
	if err != nil {
		panic(fmt.Sprintf("can't create or open file, ERR: %s", err))
	}
	j, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Sprintf("can't read to the file %v\n", err))
	}
	config := Config{}
	err = json.Unmarshal(j, &config)
	if err != nil {
		panic(fmt.Sprintf("can't unmarshal json %v\n", err))
	}
	c.config = config
}
