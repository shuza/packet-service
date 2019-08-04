package service

import (
	"net/http"
	"time"
)

func NewHttpClient() http.Client {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	return client
}
