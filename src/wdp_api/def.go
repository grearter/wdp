package wdp_api

type FailResponse struct {
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
}
