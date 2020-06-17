package fetch

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/odair-pedro/gourl/logging"
	"io"
	"io/ioutil"
	"net/http"
)

var logger = logging.GetLoggerInstance()

type RequestParams struct {
	Url   string `validate:"required"`
	Count int    `validate:"default=1"`
}

func NewRequestParams(url string, count int) *RequestParams {
	return &RequestParams{Url: url, Count: count}
}

func Request(requestParams *RequestParams) (*Result, error) {
	validate := validator.Validate{}
	err := validate.Struct(requestParams)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Invalid request: %s", err))
	}

	ch := make(chan *Response)
	defer close(ch)

	result := NewResult(requestParams)
	result.Start()
	defer result.Complete()

	for i := 0; i < requestParams.Count; i++ {
		logger.Info("Calling fetchUrl")
		go fetchUrl(requestParams.Url, ch)
	}
	for i := 0; i < requestParams.Count; i++ {
		logger.Info("Waiting for fetchUrl result")
		request := <-ch
		logger.Info(request)
		result.responses = append(result.responses, request)
	}

	logger.Info("Returning FetchResult")
	return result, nil
}

func fetchUrl(url string, ch chan<- *Response) {
	urlRequest := &Response{}
	urlRequest.Start()
	defer urlRequest.Complete()

	resp, errGet := http.Get(url)
	if errGet != nil {
		urlRequest.err = errGet
		return
	}

	length, errCopy := io.Copy(ioutil.Discard, resp.Body)
	defer closeHttpResponse(resp)

	if errCopy != nil {
		urlRequest.err = errCopy
		return
	}

	urlRequest.length = length
	ch <- urlRequest
}

func closeHttpResponse(response *http.Response) {
	if response.Body != nil {
		response.Body.Close()
	}
	response.Close = true
}

//func (urlRequest *UrlRequest) start() {
//	logging.Println("Starting UrlRequest")
//	urlRequest.StartedAt = time.Now()
//}
//
//func (urlRequest *UrlRequest) complete() {
//	logging.Println("Completing UrlRequest")
//	if urlRequest.StartedAt.IsZero() {
//		panic("Request not started")
//	}
//	urlRequest.CompletedAt = time.Now()
//	urlRequest.Duration = urlRequest.CompletedAt.Sub(urlRequest.StartedAt)
//}
