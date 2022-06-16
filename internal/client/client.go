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

type GetConvertErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
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

type ConvertRequest struct {
	From   string
	To     string
	Amount string
}

type ErrorWithConvert struct {
	Err   error
	Title string
}

func (err *ErrorWithConvert) Error() string {
	return fmt.Sprintf("%v", err.Err)
}

func (err *ErrorWithConvert) Unwrap() error {
	return err.Err
}

func NewErrorWithConvert(err error, title string) error {
	return &ErrorWithConvert{
		Err:   err,
		Title: title,
	}
}

func (c *client) Convert(convert ConvertRequest) (*GetConvertResponse, error) {
	query := url.Values{}

	query.Set("from", convert.From)
	query.Set("to", convert.To)
	query.Set("amount", convert.Amount)

	endpoint := fmt.Sprintf("%s/convert?%s", c.host, query.Encode())

	request, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("apikey", c.apiKey)

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

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

	if !data.Success {
		return nil, NewErrorWithConvert(err, "InvalidRequestFormat")
	}

	return &data, nil
}
