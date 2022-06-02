package errors

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code int

const (
	Unknown Code = iota
	AuthError
	InvalidRequest
	InvalidResponse
	NotFound
	RateLimitExhausted
	Timeout
	Unavailable
)

type StatusError struct {
	code Code
	msg  string
}

func New(code Code, msg string) StatusError {
	return StatusError{
		code: code,
		msg:  msg,
	}
}

func (ke StatusError) Error() string {
	return fmt.Sprintf("code: %v, message: %s", ke.code, ke.msg)
}

func (ke StatusError) Code() Code {
	return ke.code
}

func StatusErrorFromHTTP(response *http.Response, err error) error {
	if err == nil {
		return nil
	}
	ke := StatusError{msg: err.Error()}
	if response != nil {
		switch response.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			ke.code = AuthError
		case http.StatusBadRequest, http.StatusConflict:
			ke.code = InvalidRequest
		case http.StatusNotFound:
			ke.code = NotFound
		case http.StatusTooManyRequests:
			ke.code = RateLimitExhausted
		case http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			ke.code = Unavailable
		}
	}
	if errors.Is(err, context.DeadlineExceeded) {
		ke.code = Timeout
	}
	return ke
}

func StatusErrorFromGRPC(err error) error {
	if err == nil {
		return nil
	}
	ke := StatusError{msg: err.Error()}
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.DeadlineExceeded:
			ke.code = Timeout
		case codes.NotFound:
			ke.code = NotFound
		case codes.PermissionDenied, codes.Unauthenticated:
			ke.code = AuthError
		case codes.ResourceExhausted:
			ke.code = RateLimitExhausted
		case codes.Unavailable:
			ke.code = Unavailable
		case codes.InvalidArgument, codes.AlreadyExists, codes.FailedPrecondition, codes.OutOfRange:
			ke.code = InvalidRequest
		default:
			ke.code = Unknown
		}
	}
	return ke
}
