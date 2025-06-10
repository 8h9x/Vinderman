package request

import (
	"encoding/json"
	"errors"
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

type EpicErrorResponse struct {
	ErrorCode          string   `json:"errorCode"`
	ErrorMessage       string   `json:"errorMessage"`
	MessageVars        []string `json:"messageVars"`
	NumericErrorCode   int      `json:"numericErrorCode"`
	OriginatingService string   `json:"originatingService"`
	Intent             string   `json:"intent"`
}

func ResponseParser[T any](resp *http.Response) (Response[T], error) {
	defer resp.Body.Close()

	if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		epicError := EpicErrorResponse{}

		err := json.NewDecoder(resp.Body).Decode(&epicError)
		if err != nil {
			return Response[T]{}, Error[T]{
				StatusCode: resp.StatusCode,
				Message:    "failed to decode response error body",
				Err:        err,
			}
		}

		return Response[T]{}, Error[EpicErrorResponse]{
			StatusCode: resp.StatusCode,
			Message:    epicError.ErrorMessage,
			Err:        errors.New(epicError.ErrorMessage),
		}
	}

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
