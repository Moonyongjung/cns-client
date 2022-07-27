package types

type HttpResponseStruct struct {
	ResCode int `json:"resCode"`
	ResMsg string `json:"resMsg"`
	ResData string `json:"resData"`
}