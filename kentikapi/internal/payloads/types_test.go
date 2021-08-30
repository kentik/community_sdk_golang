package payloads_test

import (
	"encoding/json"
	"testing"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/payloads"
	"github.com/stretchr/testify/assert"
)

func TestBoolAsString_UnmarshalJSON(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input          string
		expectedResult payloads.BoolAsStringOrInt
		expectedError  bool
	}{
		{
			input:          `true`,
			expectedResult: payloads.BoolAsStringOrInt(true),
		}, {
			input:          `false`,
			expectedResult: payloads.BoolAsStringOrInt(false),
		}, {
			input:          `"true"`,
			expectedResult: payloads.BoolAsStringOrInt(true),
		}, {
			input:          `"True"`,
			expectedResult: payloads.BoolAsStringOrInt(true),
		}, {
			input:          `"false"`,
			expectedResult: payloads.BoolAsStringOrInt(false),
		}, {
			input:          `"False"`,
			expectedResult: payloads.BoolAsStringOrInt(false),
		}, {
			input:         `"invalid-string"`,
			expectedError: true,
		}, {
			input:          `1`,
			expectedResult: payloads.BoolAsStringOrInt(true),
		}, {
			input:          `0`,
			expectedResult: payloads.BoolAsStringOrInt(false),
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

			var result payloads.BoolAsStringOrInt
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
