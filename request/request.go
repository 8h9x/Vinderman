package request

import (
	"encoding/json"
	"net/http"
)

type Response[T any] struct {
	*http.Response
	Body T
}

type Error[T any] struct {
	StatusCode int
	Message    string
	Raw        T
	Err        error
}

func (e Error[T]) Error() string {
	return e.Message
}

func (e Error[T]) Unwrap() error {
	return e.Err
}

func ResponseParser[T any](resp *http.Response) (Response[T], error) {
	defer resp.Body.Close()

	var body T
	err := json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return Response[T]{}, Error[T]{
			StatusCode: resp.StatusCode,
			Message:    "failed to decode response body",
			Err:        err,
		}
	}

	return Response[T]{
		Response: resp,
		Body:     body,
	}, nil
}
