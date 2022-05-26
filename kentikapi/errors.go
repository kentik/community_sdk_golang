package kentikapi

import (
	"errors"

	kentikerrors "github.com/kentik/community_sdk_golang/kentikapi/internal/errors"
)

// IsAuthError returns true if passed error is an authentication/authorization error.
func IsAuthError(err error) bool {
	var ke kentikerrors.StatusError
	if ok := errors.As(err, &ke); ok {
		return ke.Code() == kentikerrors.AuthError
	}
	return false
}

// IsInvalidRequestError returns true if passed error is due to invalid request.
// It might be returned if user passed invalid data or there is a resource conflict.
func IsInvalidRequestError(err error) bool {
	var ke kentikerrors.StatusError
	if ok := errors.As(err, &ke); ok {
		return ke.Code() == kentikerrors.InvalidRequest
	}
	return false
}

// IsNotFoundError returns true if passed error indicates that resource was not found.
func IsNotFoundError(err error) bool {
	var ke kentikerrors.StatusError
	if ok := errors.As(err, &ke); ok {
		return ke.Code() == kentikerrors.NotFound
	}
	return false
}

// IsRateLimitExhaustedError returns true if passed error indicates that the rate limit for a resource
// has been exhausted.
func IsRateLimitExhaustedError(err error) bool {
	var ke kentikerrors.StatusError
	if ok := errors.As(err, &ke); ok {
		return ke.Code() == kentikerrors.RateLimitExhausted
	}
	return false
}

// IsTemporaryError returns true if passed error is temporary, e.g. it is due to a timeout,
// rate limit exhaustion or service unavailability.
func IsTemporaryError(err error) bool {
	var ke kentikerrors.StatusError
	if ok := errors.As(err, &ke); ok {
		if ke.Code() == kentikerrors.Timeout ||
			ke.Code() == kentikerrors.Unavailable ||
			ke.Code() == kentikerrors.RateLimitExhausted {
			return true
		}
	}
	return false
}

// IsTimeoutError returns true if passed error is a timeout error.
func IsTimeoutError(err error) bool {
	var ke kentikerrors.StatusError
	if ok := errors.As(err, &ke); ok {
		return ke.Code() == kentikerrors.Timeout
	}
	return false
}
