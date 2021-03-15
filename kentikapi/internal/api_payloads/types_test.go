package api_payloads

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolAsString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult BoolAsStringOrInt
		expectedError  bool
	}{
		{
			input:          `true`,
			expectedResult: BoolAsStringOrInt(true),
		}, {
			input:          `false`,
			expectedResult: BoolAsStringOrInt(false),
		}, {
			input:          `"true"`,
			expectedResult: BoolAsStringOrInt(true),
		}, {
			input:          `"True"`,
			expectedResult: BoolAsStringOrInt(true),
		}, {
			input:          `"false"`,
			expectedResult: BoolAsStringOrInt(false),
		}, {
			input:          `"False"`,
			expectedResult: BoolAsStringOrInt(false),
		}, {
			input:         `"invalid-string"`,
			expectedError: true,
		}, {
			input:          `1`,
			expectedResult: BoolAsStringOrInt(true),
		}, {
			input:          `0`,
			expectedResult: BoolAsStringOrInt(false),
		}, {
			input:         `1.0`,
			expectedError: true,
		}, {
			input:         `0.0`,
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			var result BoolAsStringOrInt
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
