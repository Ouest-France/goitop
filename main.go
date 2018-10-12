package goitop

import (
	"net/http"
)

type Client struct {
	Client   *http.Client
	Address  string
	User     string
	Password string
}

type APIResponse struct {
	Code    int
	Message string
	Objects map[string]APIObject
}

type APIObject struct {
	Code    int
	Message string
	Class   string
	Key     string
	Fields  map[string]interface{}
}

func NewClient(address, user, password string) *Client {
	return &Client{
		Client:   &http.Client{},
		Address:  address,
		User:     user,
		Password: password,
	}
}
