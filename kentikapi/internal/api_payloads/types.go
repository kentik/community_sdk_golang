package api_payloads

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IntAsString evens out deserialization of numbers represented in JSON document sometimes as int and sometimes as string.
type IntAsString int

func (p *IntAsString) UnmarshalJSON(data []byte) (err error) {
	// unmarshall int or string
	var obj interface{}
	if err = json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("IntAsString.UnmarshalJSON error: %v", err)
	}

	// decide the type and convert to number
	switch val := obj.(type) {
	case float64: // json.Unmarshall recognizes numbers as float64
		*p = IntAsString(val)
	case string:
		intval, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("IntAsString.UnmarshalJSON Atoi conversion error: %v", err)
		}
		*p = IntAsString(intval)
	default:
		return fmt.Errorf("IntAsString.UnmarshalJSON input should be string or int, got %v (%T)", val, val)
	}
	return nil
}

// boolAsString evens out deserialization of numbers represented in JSON document sometimes as bool and sometimes as string.
type boolAsString bool

func (p *boolAsString) UnmarshalJSON(data []byte) (err error) {
	var valueIf interface{}
	if err = json.Unmarshal(data, &valueIf); err != nil {
		return fmt.Errorf("boolAsString.UnmarshalJSON: %v", err)
	}

	switch value := valueIf.(type) {
	case bool:
		*p = boolAsString(value)
	case string:
		v, pErr := strconv.ParseBool(value)
		if pErr != nil {
			return fmt.Errorf("boolAsString.UnmarshalJSON: parse bool (%v): %v", value, pErr)
		}

		*p = boolAsString(v)
	default:
		return fmt.Errorf("boolAsString.UnmarshalJSON input should be bool or string, got %v (%T)", value, value)
	}

	return nil
}
