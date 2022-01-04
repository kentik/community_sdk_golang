package kentikapi_test

import (
	"os"
	"testing"

	"github.com/AlekSi/pointer"
	"github.com/kentik/community_sdk_golang/kentikapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadCredentialsFromEnv(t *testing.T) {
	assert.NoError(t, os.Setenv("KTAPI_AUTH_TOKEN", "12345"))
	defer func() {
		err := os.Unsetenv("KTAPI_AUTH_TOKEN")
		assert.NoError(t, err)
	}()
	assert.NoError(t, os.Setenv("KTAPI_AUTH_EMAIL", "john.doe@domain.com"))
	defer func() {
		err := os.Unsetenv("KTAPI_AUTH_EMAIL")
		assert.NoError(t, err)
	}()
	email, token, err := kentikapi.ReadCredentialsFromEnv()
	assert.NoError(t, err)
	assert.Equal(t, "12345", token)
	assert.Equal(t, "john.doe@domain.com", email)
}

func TestReadCredentialsFromProfile(t *testing.T) {
	jsonContent := `{"email": "john.doe@domain.com","api-key": "12345"}`
	tests := []struct {
		name    string
		content *string
		profile string
		config  *string
	}{
		{
			name: "default config with preexisting file ./.kentik/default",
		}, {
			name:    "create config and use it to get credentials",
			content: pointer.ToString(jsonContent),
			config:  pointer.ToString("kentik_config_for_test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.config != nil {
				require.NoError(t, os.WriteFile(*tt.config, []byte(*tt.content), os.ModeTemporary))
				defer func() {
					err := os.Remove(*tt.config)
					assert.NoError(t, err)
				}()
			}
			res, err := kentikapi.ReadCredentialsFromProfile(tt.profile)
			assert.NoError(t, err)
			assert.Equal(t, "12345", res["api-key"])
			assert.Equal(t, "john.doe@domain.com", res["email"])
		})
	}
}
