package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
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
	err := client.readConfig()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func (c *Client) StoreToken(token string) error {
	c.config.Token = token

	return c.storeConfig()
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

	return c.storeConfig()
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

func (c *Client) SendRequest(req types.Request) (*http.Response, error) {
	cl := &http.Client{
		// Timeout: 2000,
	}

	reqBody, err := req.GetBody()
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequestWithContext(context.Background(), req.Method, req.GetURL(c.config.BaseURL), reqBody)
	if err != nil {
		return nil, err
	}

	if req.AuthorizationRequired {
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
	}

	for n, v := range req.Header {
		r.Header.Set(n, v)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Accept", "application/json")

	return cl.Do(r)
}

func (c *Client) checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	return !errors.Is(err, os.ErrNotExist)
}

func (c *Client) storeConfig() error {
	file, err := os.OpenFile(configFIlePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, configFilePermission)
	if err != nil {
		return fmt.Errorf("can't create or open file, ERR: %w", err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			println(err)
		}
	}()
	j, err := json.Marshal(c.config)
	if err != nil {
		return fmt.Errorf("can't marshal json %w", err)
	}
	_, wErr := file.Write(j)
	if wErr != nil {
		return fmt.Errorf("can't write to the file %w", wErr)
	}

	return nil
}

func (c *Client) initConfig() error {
	c.config = Config{
		Token:   "",
		BaseURL: "http://manager.ormus.local",
	}

	return c.storeConfig()
}

func (c *Client) readConfig() error {
	if !c.checkFileExists(configFIlePath) {
		return c.initConfig()
	}
	file, err := os.OpenFile(configFIlePath, os.O_CREATE|os.O_RDONLY, configFilePermission)
	if err != nil {
		return fmt.Errorf("can't create or open file, ERR: %w", err)
	}
	j, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("can't read to the file %w", err)
	}
	config := Config{}
	err = json.Unmarshal(j, &config)
	if err != nil {
		return fmt.Errorf("can't unmarshal json %w", err)
	}
	c.config = config

	return nil
}
