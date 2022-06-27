// Code generated by "stringer -type=Code"; DO NOT EDIT.

package errors

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[AuthError-1]
	_ = x[InvalidRequest-2]
	_ = x[InvalidResponse-3]
	_ = x[NotFound-4]
	_ = x[RateLimitExhausted-5]
	_ = x[Timeout-6]
	_ = x[Unavailable-7]
}

const _Code_name = "UnknownAuthErrorInvalidRequestInvalidResponseNotFoundRateLimitExhaustedTimeoutUnavailable"

var _Code_index = [...]uint8{0, 7, 16, 30, 45, 53, 71, 78, 89}

func (i Code) String() string {
	if i < 0 || i >= Code(len(_Code_index)-1) {
		return "Code(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Code_name[_Code_index[i]:_Code_index[i+1]]
}