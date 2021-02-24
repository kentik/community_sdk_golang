package api_payloads

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BrokenInt evens out deserialization of numbers represented in JSON document sometimes as int and sometimes as string. Ugliness
type BrokenInt int

func (p *BrokenInt) UnmarshalJSON(data []byte) (err error) {
	// unmarshall int or string
	var obj interface{}
	if err = json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("BrokenInt.UnmarshalJSON error: %v", err)
	}

	// decide the type and convert to number
	switch val := obj.(type) {
	case float64: // json.Unmarshall recognizes numbers as float64
		*p = BrokenInt(val)
	case string:
		intval, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("BrokenInt.UnmarshalJSON Atoi conversion error: %v", err)
		}
		*p = BrokenInt(intval)
	default:
		return fmt.Errorf("BrokenInt.UnmarshalJSON input should be string or int, got {%T}", obj)
	}
	return nil
}
