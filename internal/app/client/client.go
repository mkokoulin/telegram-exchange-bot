package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type GetConvertResponse struct {
	Date       string `json:"date"`
	Historical string `json:"historical"`
	Info       struct {
		Rate      float64 `json:"rate"`
		Timestamp int     `json:"timestamp"`
	} `json:"info"`
	Query struct {
		Amount int    `json:"amount"`
		From   string `json:"from"`
		To     string `json:"to"`
	} `json:"query"`
	Result  float64 `json:"result"`
	Success bool    `json:"success"`
}

type client struct {
	host       string
	apiKey     string
	httpClient *http.Client
}

func New(host, apiKey string) *client {
	return &client{
		host:       host,
		apiKey:     apiKey,
		httpClient: &http.Client{},
	}
}

func (c *client) Convert(from, to, amount string) (*GetConvertResponse, error) {
	query := url.Values{}

	query.Set("to", to)
	query.Set("from", from)
	query.Set("amount", amount)

	endpoint := fmt.Sprintf("%s/convert?%s", c.host, query.Encode())

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("apikey", c.apiKey)

	response, err := c.httpClient.Do(request)

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var data GetConvertResponse

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
