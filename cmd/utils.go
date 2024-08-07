package cmd

import (
	"net/http"
	"time"
)

func returnResp(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrCommandNotFound
	}

	return resp, nil
}
