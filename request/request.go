package request

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type Response[T any] struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           http.Header
	Body             T
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Uncompressed     bool
	Trailer          http.Header
	Request          *http.Request
	TLS              *tls.ConnectionState

	Cookies []*http.Cookie
}

type Error[T any] struct {
	StatusCode int
	Message    string
	Raw        T
}

func (e Error[T]) Error() string {
	return e.Message
}

func ResponseParser[T any](resp *http.Response) (Response[T], error) {
	defer resp.Body.Close()

	var res T
	err := json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return Response[T]{}, err
	}

	return Response[T]{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMajor:       resp.ProtoMajor,
		ProtoMinor:       resp.ProtoMinor,
		Header:           resp.Header,
		Body:             res,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Close:            resp.Close,
		Uncompressed:     resp.Uncompressed,
		Trailer:          resp.Trailer,
		Request:          resp.Request,
		TLS:              resp.TLS,
	}, nil
}

func Getf[T any](url string, a ...any) (res Response[T], err error) {
	resp, err := http.Get(fmt.Sprintf(url, a...))
	if err != nil {
		return
	}

	return ResponseParser[T](resp)
}
