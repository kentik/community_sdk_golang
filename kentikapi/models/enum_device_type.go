// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _DeviceTypeName = "routerhost-nprobe-dns-www"

var _DeviceTypeIndex = [...]uint8{0, 6, 25}

func (i DeviceType) String() string {
	if i < 0 || i >= DeviceType(len(_DeviceTypeIndex)-1) {
		return fmt.Sprintf("DeviceType(%d)", i)
	}
	return _DeviceTypeName[_DeviceTypeIndex[i]:_DeviceTypeIndex[i+1]]
}

var _DeviceTypeValues = []DeviceType{0, 1}

// modified by hand
var _DeviceTypeNameToValueMap = map[string]DeviceType{
	_DeviceTypeName[0:6]:  0,
	_DeviceTypeName[6:25]: 1,
}

// DeviceTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DeviceTypeString(s string) (DeviceType, error) {
	if val, ok := _DeviceTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DeviceType values", s)
}

// DeviceTypeValues returns all values of the enum
func DeviceTypeValues() []DeviceType {
	return _DeviceTypeValues
}

// IsADeviceType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DeviceType) IsADeviceType() bool {
	for _, v := range _DeviceTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
