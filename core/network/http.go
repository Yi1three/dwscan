package network

import (
	"github.com/corpix/uarand"
	"io"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"time"
)

type Request struct {
	Url      string
	Method   string
	Data     string
	Redirect bool
	Headers  map[string]string
}

type Response struct {
	Status     string
	Header     http.Header
	Body       string
	Location   string
	RequestUrl string
}

func Send(request *Request) (*Response, error) {
	//TODO 支持代理

	client := &http.Client{
		Timeout: time.Duration(5) * time.Second,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	if request.Redirect {
		jar, _ := cookiejar.New(nil)
		client = &http.Client{
			Timeout: time.Duration(5) * time.Second,
			Jar:     jar,
		}
	}

	sendRequest, err := http.NewRequest(strings.ToUpper(request.Method), request.Url, strings.NewReader(request.Data))
	if err != nil {
		return nil, err
	}

	//default header set
	if request.Headers["Content-Type"] == "" {
		sendRequest.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	}
	sendRequest.Header.Set("User-Agent", uarand.GetRandom())

	//user defined header set
	for header, value := range request.Headers {
		sendRequest.Header.Set(header, value)
	}

	response, err := client.Do(sendRequest)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	var location string
	var body string
	tmpBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	body = string(tmpBody)

	tmpLocation, err := response.Location()
	if err != nil {
		return nil, err
	}
	location = tmpLocation.String()

	return &Response{response.Status, response.Header, body, location, request.Url}, nil
}
