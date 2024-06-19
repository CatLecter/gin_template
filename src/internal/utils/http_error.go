package utils

type HTTPError struct {
	Result string `json:"result"`
	Msg    string `json:"msg"`
}

func NewHTTPError(msg string) *HTTPError { return &HTTPError{Result: "error", Msg: msg} }
