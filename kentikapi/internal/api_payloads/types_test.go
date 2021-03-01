package api_payloads

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolAsString_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult boolAsString
		expectedError  bool
	}{
		{
			input:          `true`,
			expectedResult: boolAsString(true),
		}, {
			input:          `false`,
			expectedResult: boolAsString(false),
		}, {
			input:          `"true"`,
			expectedResult: boolAsString(true),
		}, {
			input:          `"True"`,
			expectedResult: boolAsString(true),
		}, {
			input:          `"false"`,
			expectedResult: boolAsString(false),
		}, {
			input:          `"False"`,
			expectedResult: boolAsString(false),
		}, {
			input:         `"invalid-string"`,
			expectedError: true,
		}, {
			input:         `1`,
			expectedError: true,
		}, {
			input:         `0`,
			expectedError: true,
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
			var result boolAsString
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
