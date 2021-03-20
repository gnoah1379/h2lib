package herr

import "errors"

type Error interface {
	Error() string
	Cause() error
	Set(field string, value interface{}) Error
	Field(name string) interface{}
	Detail() ErrorDetail
}

type ErrorDetail struct {
	Error  string                 `json:"error" xml:"error" yaml:"error" mapstructure:"error"`
	Detail map[string]interface{} `json:"detail" xml:"detail" yaml:"detail" mapstructure:"detail"`
}

type HErr struct {
	err    error
	detail map[string]interface{}
}

func New(msg string) Error {
	return &HErr{
		err:    errors.New(msg),
		detail: make(map[string]interface{}),
	}
}

func Wrap(err error) Error {
	return &HErr{
		err:    err,
		detail: make(map[string]interface{}),
	}
}

func (e HErr) Error() string {
	return e.err.Error()
}

func (e HErr) Cause() error {
	return e.err
}

func (e *HErr) Set(field string, value interface{}) Error {
	e.detail[field] = value
	return e
}

func (e HErr) Field(name string) interface{} {
	return e.detail[name]
}

func (e HErr) Detail() ErrorDetail {
	return ErrorDetail{
		Error:  e.Error(),
		Detail: e.detail,
	}
}
