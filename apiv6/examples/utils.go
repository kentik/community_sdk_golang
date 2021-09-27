package examples

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/kentik/community_sdk_golang/kentikapi"
)

// PrettyPrint prints an object recursively in an indented way
func PrettyPrint(resource interface{}) {
	prettyPrintRecursively(reflect.TypeOf(resource), reflect.ValueOf(resource), 0)
}

func prettyPrintRecursively(t reflect.Type, v reflect.Value, level int) {
	switch v.Kind() {

	case reflect.Struct:
		if _, hasStringer := t.MethodByName("String"); hasStringer {
			prettyPrintIndented("%v\n", level, v)
			return
		}
		for i := 0; i < v.NumField(); i++ {
			prettyPrintIndented("%s:\n", level, t.Field(i).Name)
			prettyPrintRecursively(v.Field(i).Type(), v.Field(i), level+1)
		}

	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			prettyPrintIndented("[no items]\n", level)

		} else {
			for i := 0; i < count; i++ {
				prettyPrintIndented("[%d]\n", level, i)
				s := v.Index(i)
				prettyPrintRecursively(s.Type(), s, level+1)
			}
		}

	case reflect.Map:
		for _, key := range v.MapKeys() {
			prettyPrintIndented("[%s]\n", level, key)
			s := v.MapIndex(key)
			prettyPrintRecursively(s.Type(), s, level+1)
		}

	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			prettyPrintIndented("[empty]\n", level)
		} else {
			prettyPrintRecursively(v.Elem().Type(), v.Elem(), level)
		}

	default:
		prettyPrintIndented("%v\n", level, v)
	}
}

func prettyPrintIndented(format string, level int, args ...interface{}) {
	fmt.Printf("%*s", level*2, "")
	fmt.Printf(format, args...)
}

// PanicOnError converts err into panic; use it to reduce the number of: "if err != nil { return err }" statements
func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

// NewClient creates kentikapi client with credentials read from env variables
func NewClient() *kentikapi.Client {
	email, token, err := readCredentialsFromEnv()
	PanicOnError(err)
	cfg := kentikapi.Config{
		AuthEmail: email,
		AuthToken: token,

		// examples can also be run against localhost_apiserver by specifying the server url:
		// SyntheticsAPIURL:  "http://localhost:8080",
		// CloudExportAPIURL: "http://localhost:8080",
	}
	return kentikapi.NewClient(cfg)
}

func readCredentialsFromEnv() (authEmail, authToken string, _ error) {
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
