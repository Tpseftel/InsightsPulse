package apiclients

import "github.com/go-resty/resty/v2"

type ApiClient interface {
	GetClient() *resty.Client
	IsClientOk() bool
	CheckRequestsLimits(resp *resty.Response)
}
