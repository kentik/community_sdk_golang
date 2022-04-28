package kentik_errors

import (
	"errors"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Code int

const (
	AuthError          Code = 0
	InvalidRequest     Code = 1
	NotFound           Code = 2
	RateLimitExhausted Code = 3
	Temporary          Code = 4
	Timeout            Code = 5
)

type KentikError struct {
	Codes []Code
	Msg   string
}

func (ke KentikError) Error() string {
	return fmt.Sprintf("category: %v, message: %s", ke.Codes, ke.Msg)
}

func (ke KentikError) GetKentikError() *KentikError { return &ke }

func GetCodes(err error) ([]Code, bool) {
	var tErr interface {
		GetKentikError() *KentikError
	}
	if ok := errors.As(err, &tErr); ok {
		ktError := tErr.GetKentikError()
		return ktError.Codes, true
	}
	return []Code{}, false
}

func KentikErrorUpdateMsg(msg string, err error) error {
	var tErr interface {
		GetKentikError() *KentikError
	}
	if ok := errors.As(err, &tErr); ok {
		ktError := tErr.GetKentikError()
		ktError.Msg = msg
		return ktError
	}
	return err
}

func KentikErrorFromHTTP(response *http.Response, err error) error {
	if err == nil {
		return nil
	}
	ke := KentikError{Msg: err.Error()}
	if response != nil {
		switch response.StatusCode {
		case http.StatusUnauthorized, http.StatusForbidden:
			ke.Codes = []Code{AuthError}
		case http.StatusNotFound:
			ke.Codes = []Code{NotFound}
		case http.StatusRequestTimeout:
			ke.Codes = []Code{Timeout, Temporary}
		case http.StatusTooManyRequests:
			ke.Codes = []Code{RateLimitExhausted, Temporary}
		case http.StatusServiceUnavailable, http.StatusGatewayTimeout:
			ke.Codes = []Code{Temporary}
		default:
			ke.Codes = []Code{InvalidRequest}
		}
	} else {
		ke.Codes = []Code{Timeout}
	}
	return ke
}

func KentikErrorFromGRPC(err error) error {
	if err == nil {
		return nil
	}
	ke := KentikError{Msg: err.Error()}
	if e, ok := status.FromError(err); ok {
		switch e.Code() {
		case codes.DeadlineExceeded:
			ke.Codes = []Code{Timeout, Temporary}
		case codes.NotFound:
			ke.Codes = []Code{NotFound}
		case codes.PermissionDenied, codes.Unauthenticated:
			ke.Codes = []Code{AuthError}
		case codes.ResourceExhausted:
			ke.Codes = []Code{RateLimitExhausted, Temporary}
		case codes.Unavailable:
			ke.Codes = []Code{Temporary}
		default:
			ke.Codes = []Code{InvalidRequest}
		}
	}
	return ke
}

func (ke KentikError) AuthError() bool {
	for _, c := range ke.Codes {
		if c == AuthError {
			return true
		}
	}
	return false
}

func (ke KentikError) InvalidRequest() bool {
	for _, c := range ke.Codes {
		if c == InvalidRequest {
			return true
		}
	}
	return false
}

func (ke KentikError) NotFound() bool {
	for _, c := range ke.Codes {
		if c == NotFound {
			return true
		}
	}
	return false
}

func (ke KentikError) RateLimitExhausted() bool {
	for _, c := range ke.Codes {
		if c == RateLimitExhausted {
			return true
		}
	}
	return false
}

func (ke KentikError) Temporary() bool {
	for _, c := range ke.Codes {
		if c == Temporary {
			return true
		}
	}
	return false
}

func (ke KentikError) Timeout() bool {
	for _, c := range ke.Codes {
		if c == Timeout {
			return true
		}
	}
	return false
}

func IsAuthError(err error) bool {
	var ktErr interface {
		AuthError() bool
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.AuthError() {
			return true
		}
	}
	return false
}

func IsInvalidRequestError(err error) bool {
	var ktErr interface {
		InvalidRequest() bool
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.InvalidRequest() {
			return true
		}
	}
	return false
}

func IsNotFoundError(err error) bool {
	var ktErr interface {
		NotFound() bool
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.NotFound() {
			return true
		}
	}
	return false
}

func IsRateLimitExhaustedError(err error) bool {
	var ktErr interface {
		RateLimitExhausted() bool
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.RateLimitExhausted() {
			return true
		}
	}
	return false
}

func IsTemporaryError(err error) bool {
	var ktErr interface {
		Temporary() bool
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.Temporary() {
			return true
		}
	}
	return false
}

func IsTimeoutError(err error) bool {
	var ktErr interface {
		Timeout() bool
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.Timeout() {
			return true
		}
	}
	return false
}
