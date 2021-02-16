// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _PrivacyProtocolName = "NoPrivDESAES"

var _PrivacyProtocolIndex = [...]uint8{0, 6, 9, 12}

func (i PrivacyProtocol) String() string {
	if i < 0 || i >= PrivacyProtocol(len(_PrivacyProtocolIndex)-1) {
		return fmt.Sprintf("PrivacyProtocol(%d)", i)
	}
	return _PrivacyProtocolName[_PrivacyProtocolIndex[i]:_PrivacyProtocolIndex[i+1]]
}

var _PrivacyProtocolValues = []PrivacyProtocol{0, 1, 2}

var _PrivacyProtocolNameToValueMap = map[string]PrivacyProtocol{
	_PrivacyProtocolName[0:6]:  0,
	_PrivacyProtocolName[6:9]:  1,
	_PrivacyProtocolName[9:12]: 2,
}

// PrivacyProtocolString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrivacyProtocolString(s string) (PrivacyProtocol, error) {
	if val, ok := _PrivacyProtocolNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to PrivacyProtocol values", s)
}

// PrivacyProtocolValues returns all values of the enum
func PrivacyProtocolValues() []PrivacyProtocol {
	return _PrivacyProtocolValues
}

// IsAPrivacyProtocol returns "true" if the value is listed in the enum definition. "false" otherwise
func (i PrivacyProtocol) IsAPrivacyProtocol() bool {
	for _, v := range _PrivacyProtocolValues {
		if i == v {
			return true
		}
	}
	return false
}
