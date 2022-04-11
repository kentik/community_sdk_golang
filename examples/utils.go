//nolint:forbidigo
package examples

import (
	"fmt"
	"log"
	"reflect"

	"github.com/kentik/community_sdk_golang/kentikapi"
)

// NewClient creates kentikapi client with credentials read from env variables.
func NewClient() (*kentikapi.Client, error) {
	email, token, err := kentikapi.GetCredentials(kentikapi.DefaultProfile)
	if err != nil {
		return nil, err
	}

	client, err := kentikapi.NewClient(
		kentikapi.WithCredentials(email, token),
		kentikapi.WithLogPayloads(),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// PrettyPrint prints an object recursively in an indented way.
func PrettyPrint(resource interface{}) {
	prettyPrintRecursively(reflect.TypeOf(resource), reflect.ValueOf(resource), 0)
}

//nolint:gocyclo
func prettyPrintRecursively(t reflect.Type, v reflect.Value, level int) {
	const maxDepth = 20
	if level > maxDepth {
		log.Printf("Pretty print depth level exceeded %v - omitting\n", maxDepth)
		return
	}

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

//nolint:gomnd
func prettyPrintIndented(format string, level int, args ...interface{}) {
	fmt.Printf("%*s", level*2, "")
	fmt.Printf(format, args...)
}

func stringPointer(s string) *string {
	return &s
}
