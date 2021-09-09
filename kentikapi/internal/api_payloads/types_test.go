package api_payloads_test

import (
	"encoding/json"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/stretchr/testify/assert"
)

func TestBoolAsString_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input          string
		expectedResult api_payloads.BoolAsStringOrInt
		expectedError  bool
	}{
		{
			input:          `true`,
			expectedResult: api_payloads.BoolAsStringOrInt(true),
		}, {
			input:          `false`,
			expectedResult: api_payloads.BoolAsStringOrInt(false),
		}, {
			input:          `"true"`,
			expectedResult: api_payloads.BoolAsStringOrInt(true),
		}, {
			input:          `"True"`,
			expectedResult: api_payloads.BoolAsStringOrInt(true),
		}, {
			input:          `"false"`,
			expectedResult: api_payloads.BoolAsStringOrInt(false),
		}, {
			input:          `"False"`,
			expectedResult: api_payloads.BoolAsStringOrInt(false),
		}, {
			input:         `"invalid-string"`,
			expectedError: true,
		}, {
			input:          `1`,
			expectedResult: api_payloads.BoolAsStringOrInt(true),
		}, {
			input:          `0`,
			expectedResult: api_payloads.BoolAsStringOrInt(false),
		}, {
			input:         `1.0`,
			expectedError: true,
		}, {
			input:         `0.0`,
			expectedError: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()

			var result api_payloads.BoolAsStringOrInt
			err := json.Unmarshal([]byte(tt.input), &result)

			if tt.expectedError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, tt.expectedResult, result)
		})
	}
}
