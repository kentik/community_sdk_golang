// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _CDNAttributeName = "NoneYN"

var _CDNAttributeIndex = [...]uint8{0, 4, 5, 6}

func (i CDNAttribute) String() string {
	if i < 0 || i >= CDNAttribute(len(_CDNAttributeIndex)-1) {
		return fmt.Sprintf("CDNAttribute(%d)", i)
	}
	return _CDNAttributeName[_CDNAttributeIndex[i]:_CDNAttributeIndex[i+1]]
}

var _CDNAttributeValues = []CDNAttribute{0, 1, 2}

var _CDNAttributeNameToValueMap = map[string]CDNAttribute{
	_CDNAttributeName[0:4]: 0,
	_CDNAttributeName[4:5]: 1,
	_CDNAttributeName[5:6]: 2,
}

// CDNAttributeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CDNAttributeString(s string) (CDNAttribute, error) {
	if val, ok := _CDNAttributeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CDNAttribute values", s)
}

// CDNAttributeValues returns all values of the enum
func CDNAttributeValues() []CDNAttribute {
	return _CDNAttributeValues
}

// IsACDNAttribute returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CDNAttribute) IsACDNAttribute() bool {
	for _, v := range _CDNAttributeValues {
		if i == v {
			return true
		}
	}
	return false
}
