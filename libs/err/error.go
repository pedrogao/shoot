package err

import "fmt"

type HTTPErr struct {
	ErrorCode int         `json:"error_code"` // specify error code
	Msg       interface{} `json:"msg"`        // error message, always is a string, but a map in some times
	HTTPCode  int         `json:"-"`
}

func NewHTTPErr(errorCode, httpCode int, msg interface{}) *HTTPErr {
	return &HTTPErr{
		ErrorCode: errorCode,
		Msg:       msg,
		HTTPCode:  httpCode,
	}
}

func (h HTTPErr) Error() string {
	return fmt.Sprintf("%v", h.Msg)
}

func (h *HTTPErr) SetHTTPCode(httpCode int) *HTTPErr {
	h.HTTPCode = httpCode
	return h
}

func (h *HTTPErr) Clear() *HTTPErr {
	h.Msg = nil
	return h
}

func (h *HTTPErr) Set(msg interface{}) *HTTPErr {
	h.Msg = msg
	return h
}

func (h *HTTPErr) Setf(format string, args ...interface{}) *HTTPErr {
	h.Msg = " " + fmt.Sprintf(format, args...)
	return h
}
