package kentikapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// ReadCredentialsFromEnv reads and returns (email, token) pair from environment variables, or error if not set.
func ReadCredentialsFromEnv() (authEmail, authToken string, _ error) {
	authEmail, ok := os.LookupEnv("KTAPI_AUTH_EMAIL")
	if !ok || authEmail == "" {
		return "", "", errors.New("KTAPI_AUTH_EMAIL environment variable needs to be set")
	}

	authToken, ok = os.LookupEnv("KTAPI_AUTH_TOKEN")
	if !ok || authToken == "" {
		return "", "", errors.New("KTAPI_AUTH_TOKEN environment variable needs to be set")
	}

	return authEmail, authToken, nil
}

//nolint:gosec //Linter still warns about G307 but it should be dealt with.
// https://www.joeshaw.org/dont-defer-close-on-writable-files/
func ReadCredentialsFromProfile(profile string) (result map[string]interface{}, err error) {
	if profile == "" {
		profile = "default"
	}

	filename := getFilename(profile)

	jsonFile, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}
	for _, s := range []string{"email", "api-key"} {
		if _, ok := result[s]; !ok {
			return nil, fmt.Errorf("no '%s' in config file: %s", s, filename)
		}
	}
	return result, nil
}

func getFilename(profile string) string {
	home, ok := os.LookupEnv("KTAPI_HOME")
	if !ok || home == "" {
		home, ok = os.LookupEnv("HOME")
		if !ok || home == "" {
			home = "."
		}
	}
	filename, ok := os.LookupEnv("KTAPI_CFG_FILE")
	if !ok || filename == "" {
		return filepath.Join(home, ".kentik", profile)
	}
	return filename
}
