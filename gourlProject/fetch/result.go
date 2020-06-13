package fetch

import (
	"time"
)

type Result struct {
	*RequestParams
	*Metrics
	meanDuration time.Duration
	responses    []*Response
}

func NewResult(requestParams *RequestParams) *Result {
	return &Result{RequestParams: requestParams}
}

func (result *Result) MeanDuration() time.Duration {
	return result.meanDuration
}

func (result *Result) Responses() []*Response {
	return result.responses
}

func (result *Result) Start() {
	result.Metrics.Start()
	result.responses = []*Response{}
}

func (result *Result) Complete() {
	result.Metrics.Complete()
	result.meanDuration = time.Duration(result.Duration().Milliseconds() / int64(len(result.responses)))
}

//func (result *Result) calcMeanDuration() {
//	//logging.Println("Calculating result duration")
//	//for _, request := range result.Requests {
//	//	result.TotalDuration += request.Duration
//	//}
//	result.MeanDuration = time.Duration(result.TotalDuration.Milliseconds() / int64(len(result.Requests)))
//}
