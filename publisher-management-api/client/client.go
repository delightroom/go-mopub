package client

import "fmt"

func SayHi(name string) string {
	return fmt.Sprintf("hello~~~~~~~!!!!! %s", name)
}

type Client interface {
	GetLineItem(aa string) (string, error)
}

type ApiClient struct {
	ApiKey string `json:"api_key"`
}

func NewApiClient(apiKey string) ApiClient {
	return ApiClient{ApiKey: apiKey}
}

func (a ApiClient) GetLineItem(aa string) (string, error) {
	return aa, nil
}
