package main

import (
	"github.com/go-chi/render"
	"net/http"
)

type ErrorResponse struct {
	Err        error  `json:"-"`
	StatusCode int    `json:"-"`
	StatusText string `json:"status_text"`
	Message    string `json:"message"`
}

var (
	ErrMethodNotAllowed = &ErrorResponse{StatusCode: 204, Message: "Method not allowed"}
	ErrNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found"}
	ErrBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request"}
	Code201             = &ErrorResponse{StatusCode: 201, Message: ""}
)

func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}
func ErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 404,
		StatusText: "Bad request 404",
		Message:    err.Error(),
	}
}
func ServerErrorRenderer(err error) *ErrorResponse {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "Internal server error",
		Message:    err.Error(),
	}
}

//func Server201Renderer(err error) *http.Response {
//	return &http.Response{
//		Status:           "201",
//		StatusCode:       201,
//		Proto:            "",
//		ProtoMajor:       0,
//		ProtoMinor:       0,
//		Header:           nil,
//		Body:             nil,
//		ContentLength:    0,
//		TransferEncoding: nil,
//		Close:            false,
//		Uncompressed:     false,
//		Trailer:          nil,
//		Request:          nil,
//		TLS:              nil,
//	}
//}
