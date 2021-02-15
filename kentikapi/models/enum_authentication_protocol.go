// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _AuthenticationProtocolName = "NoAuthMD5SHA"

var _AuthenticationProtocolIndex = [...]uint8{0, 6, 9, 12}

func (i AuthenticationProtocol) String() string {
	if i < 0 || i >= AuthenticationProtocol(len(_AuthenticationProtocolIndex)-1) {
		return fmt.Sprintf("AuthenticationProtocol(%d)", i)
	}
	return _AuthenticationProtocolName[_AuthenticationProtocolIndex[i]:_AuthenticationProtocolIndex[i+1]]
}

var _AuthenticationProtocolValues = []AuthenticationProtocol{0, 1, 2}

var _AuthenticationProtocolNameToValueMap = map[string]AuthenticationProtocol{
	_AuthenticationProtocolName[0:6]:  0,
	_AuthenticationProtocolName[6:9]: 1,
	_AuthenticationProtocolName[9:12]: 2,
}

// AuthenticationProtocolString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AuthenticationProtocolString(s string) (AuthenticationProtocol, error) {
	if val, ok := _AuthenticationProtocolNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AuthenticationProtocol values", s)
}

// AuthenticationProtocolValues returns all values of the enum
func AuthenticationProtocolValues() []AuthenticationProtocol {
	return _AuthenticationProtocolValues
}

// IsAAuthenticationProtocol returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AuthenticationProtocol) IsAAuthenticationProtocol() bool {
	for _, v := range _AuthenticationProtocolValues {
		if i == v {
			return true
		}
	}
	return false
}
