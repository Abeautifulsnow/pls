package cmd

import "net/http"

func returnResp(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, ErrCommandNotFound
	}

	return resp, nil
}
