package apiclients

import (
	"time"

	"insights-pulse/src/config"
	"insights-pulse/src/logger"

	"github.com/go-resty/resty/v2"
)

const (
	TIMEOUT = 60 * time.Second
)

type ApiFootballClient struct {
	client *resty.Client
}

func NewApiFootballClientImp() *ApiFootballClient {
	log := logger.GetLogger()
	config, err := config.GetConfig()
	if err != nil {
		log.Error("Cannot load Configuration variables: " + err.Error())
		panic(err)
	}

	// INFO: Setup Client Connection
	client := resty.New().
		SetTimeout(TIMEOUT).
		SetHeader("x-apisports-key", config.ApiKeyApiFootball).
		SetBaseURL(config.UrlApiFootball)

	return &ApiFootballClient{
		client: client,
	}
}

func (c *ApiFootballClient) GetClient() *resty.Client {
	return c.client
}

func (c *ApiFootballClient) IsClientOk() bool {
	log := logger.GetLogger()
	endpoint := "/status"

	resp, err := c.client.R().
		Get(endpoint)
	if err != nil {
		return false
	}
	if resp.IsError() {
		log.Warn("api respose code is >=400")
		return false
	}

	return resp.StatusCode() == 200
}

// TODO: Implement me
func (c *ApiFootballClient) HasClientAvailableRequests() bool {
	// INFO: Get date from status request responce
	return true
}
