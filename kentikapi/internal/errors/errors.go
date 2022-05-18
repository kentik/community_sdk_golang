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
	AuthError          Code = 1
	InvalidRequest     Code = 2
	NotFound           Code = 3
	RateLimitExhausted Code = 4
	Temporary          Code = 5
	Timeout            Code = 6
	Unavailable        Code = 7
	Unknown            Code = 8
	InvalidResponse    Code = 9
)

type KentikError struct {
	code Code
	msg  string
}

func New(code Code, msg string) KentikError {
	return KentikError{
		code: code,
		msg:  msg,
	}
}

func (ke KentikError) Error() string {
	return fmt.Sprintf("code: %v, message: %s", ke.code, ke.msg)
}

func (ke KentikError) Code() Code {
	if ke.code == 0 {
		return 0
	}
	return ke.code
}

func KentikErrorFromHTTP(response *http.Response, err error) error {
	if err == nil {
		return nil
	}
	ke := KentikError{msg: err.Error()}
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
	if ke.code == 0 {
		ke.code = Unknown
	}
	return ke
}

func KentikErrorFromGRPC(err error) error {
	if err == nil {
		return nil
	}
	ke := KentikError{msg: err.Error()}
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
	} else {
		ke.code = Unknown
	}
	return ke
}
