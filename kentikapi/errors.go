package kentikapi

import (
	"errors"

	kentikError "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
)

// IsAuthError checks if passed error is Authorization error, if yes return true.
func IsAuthError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.AuthError
	}
	return false
}

// IsInvalidRequestError checks if passed error is invalid request error, if yes return true.
func IsInvalidRequestError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.InvalidRequest
	}
	return false
}

// IsNotFoundError checks if passed error is not found error, if yes return true.
func IsNotFoundError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.NotFound
	}
	return false
}

// IsRateLimitExhaustedError checks if passed error is rate limit exhausted error, if yes return true.
func IsRateLimitExhaustedError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.RateLimitExhausted
	}
	return false
}

// IsTemporaryError checks if passed error is temporary error, if yes return true.
func IsTemporaryError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		if ktErr.Code() == kentikError.Temporary ||
			ktErr.Code() == kentikError.Timeout ||
			ktErr.Code() == kentikError.RateLimitExhausted {
			return true
		}
	}
	return false
}

// IsTimeoutError checks if passed error is timeout error, if yes return true.
func IsTimeoutError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.Timeout
	}
	return false
}

// IsUnavailableError checks if passed error is unavailable error, if yes return true.
func IsUnavailableError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.Unavailable
	}
	return false
}

// IsUnknownError checks if passed error is unknown error, if yes return true.
func IsUnknownError(err error) bool {
	var ktErr interface {
		Code() kentikError.Code
	}
	if ok := errors.As(err, &ktErr); ok {
		return ktErr.Code() == kentikError.Unknown
	}
	return true
}
