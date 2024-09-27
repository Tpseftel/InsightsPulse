package apiclients

import (
	"fmt"
	"time"

	"insights-pulse/src/config"
	"insights-pulse/src/logger"

	"github.com/go-resty/resty/v2"
)

const (
	TIMEOUT                       = 60 * time.Second
	MINUTELY_LIMIT                = "X-RateLimit-Limit"
	REMAINING_REQUESTS_PER_MINUTE = "X-RateLimit-Remaining"
	DAILY_LIMIT                   = "x-ratelimit-requests-limit"
	REMAINING_REQUESTS_PER_DAY    = "x-ratelimit-requests-remaining"
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
	endpoint := "/status"

	resp, err := c.client.R().
		Get(endpoint)
	if err != nil {
		return false
	}
	if resp.IsError() {
		logger.GetLogger().Warn("api respose code is >=400")
		return false
	}

	return resp.StatusCode() == 200
}

func (c *ApiFootballClient) CheckRequestsLimits(resp *resty.Response) {
	if !hasAvailableRequestsMinutely(resp) {
		fmt.Println("Header feedback:Go for sleep")
		logger.GetLogger().Info("Maximum Requests per minute reached: Going for sleep!!")
		time.Sleep(1 * time.Minute)
	}
	if !hastAvailableRequestsDaily(resp) {
		panic("Daily requests limit reached!!!")
	}

}

func hasAvailableRequestsMinutely(resp *resty.Response) bool {
	limit := resp.Header().Get(REMAINING_REQUESTS_PER_MINUTE)
	fmt.Printf("Minute Remaining: %v \n", limit)
	return limit != "0"
}

func hastAvailableRequestsDaily(resp *resty.Response) bool {
	limit := resp.Header().Get(REMAINING_REQUESTS_PER_DAY)
	fmt.Printf("Day Remaining: %v \n", limit)
	return limit != "0"
}
