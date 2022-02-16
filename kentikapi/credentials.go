package kentikapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

const DefaultProfile = "default"

// GetCredentials reads and returns (email, token) pair from environment variables or specified profile config file.
// Profile name specifies file name containing Kentik profile data. DefaultProfile constant can be used for that.
// Credentials lookup order (assuming profile name is "default"):
// 1. env vars: KTAPI_AUTH_EMAIL, KTAPI_AUTH_TOKEN
// 2. file at $KTAPI_CFG_FILE
// 3. file at $KTAPI_HOME/.kentik/default
// 4. file at $HOME/.kentik/default
// 5. file at ./.kentik/default
// Profile file format: {"email":"dummy@acme.com","api-key":"dummy"}.
func GetCredentials(profileName string) (authEmail, authToken string, err error) {
	authEmail, authToken = readCredentialsFromEnv()
	if authEmail == "" || authToken == "" {
		authEmail, authToken, err = readCredentialsFromProfile(profileName)
	} else {
		log.Printf("Credentials loaded from environment variables\n")
	}
	return authEmail, authToken, err
}

// ReadCredentialsFromEnv reads and returns (email, token) pair from environment variables.
func readCredentialsFromEnv() (authEmail, authToken string) {
	authEmail, _ = os.LookupEnv("KTAPI_AUTH_EMAIL")
	authToken, _ = os.LookupEnv("KTAPI_AUTH_TOKEN")
	return authEmail, authToken
}

type profile struct {
	AuthEmail string `json:"email"`
	AuthToken string `json:"api-key"`
}

// ReadCredentialsFromProfile reads and returns (email, token) pair from profile config file.
func readCredentialsFromProfile(profileName string) (authEmail, authToken string, err error) {
	var result profile
	filename := getFilename(profileName)
	jsonFile, err := os.Open(filepath.Clean(filename))
	if err != nil {
		return "", "", err
	}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return "", "", err
	}
	if cErr := jsonFile.Close(); cErr != nil {
		return "", "", cErr
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return "", "", err
	}
	if result.AuthEmail == "" {
		return "", "", fmt.Errorf("no AuthEmail in config file: %s", filename)
	}
	if result.AuthToken == "" {
		return "", "", fmt.Errorf("no AuthToken in config file: %s", filename)
	}
	log.Printf("Credentials loaded from %v\n", filename)
	return result.AuthEmail, result.AuthToken, nil
}

func getFilename(profileName string) string {
	home := resolveHome()
	filename, ok := os.LookupEnv("KTAPI_CFG_FILE")
	if !ok || filename == "" {
		return filepath.Join(home, ".kentik", profileName)
	}
	return filename
}

func resolveHome() string {
	home, ok := os.LookupEnv("KTAPI_HOME")
	if !ok || home == "" {
		home, ok = os.LookupEnv("HOME")
		if !ok || home == "" {
			home = "."
		}
	}
	return home
}
