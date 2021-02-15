package examples

import (
	"errors"
	"fmt"
	"os"
	"reflect"
)

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

// prettyPrint prints an object recursively in an indented way
func prettyPrint(resource interface{}) {
	prettyPrintRecursively(reflect.TypeOf(resource), reflect.ValueOf(resource), 0)
}

func prettyPrintRecursively(t reflect.Type, v reflect.Value, level int) {
	switch v.Kind() {

	case reflect.Struct:
		if _, hasStringer := t.MethodByName("String"); hasStringer {
			printfIndented("%v\n", level, v)
			return
		}
		for i := 0; i < v.NumField(); i++ {
			printfIndented("%s:\n", level, t.Field(i).Name)
			prettyPrintRecursively(v.Field(i).Type(), v.Field(i), level+1)
		}

	case reflect.Slice:
		count := v.Len()
		if count == 0 {
			printfIndented("[no items]\n", level)

		} else {
			for i := 0; i < count; i++ {
				printfIndented("[%d]\n", level, i)
				s := v.Index(i)
				prettyPrintRecursively(s.Type(), s, level)
			}
		}

	case reflect.Ptr, reflect.Interface:
		if v.IsNil() {
			printfIndented("[empty]\n", level)
		} else {
			prettyPrintRecursively(v.Elem().Type(), v.Elem(), level)
		}

	default:
		printfIndented("%v\n", level, v)
	}

}

// Indented Printf
func printfIndented(format string, level int, args ...interface{}) {
	fmt.Printf("%*s", level*2, "")
	fmt.Printf(format, args...)
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
