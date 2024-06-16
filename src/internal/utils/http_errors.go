package utils

type HTTPError struct {
	Result string `json:"result"`
	Detail string `json:"detail"`
}

func (e *HTTPError) Error() string {
	return e.Detail
}

func NewError(msg string) *HTTPError {
	return &HTTPError{Result: "error", Detail: msg}
}
