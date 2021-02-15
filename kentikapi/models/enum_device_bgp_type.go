// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _DeviceBGPTypeName = "nonedeviceother_device"

var _DeviceBGPTypeIndex = [...]uint8{0, 4, 10, 22}

func (i DeviceBGPType) String() string {
	if i < 0 || i >= DeviceBGPType(len(_DeviceBGPTypeIndex)-1) {
		return fmt.Sprintf("DeviceBGPType(%d)", i)
	}
	return _DeviceBGPTypeName[_DeviceBGPTypeIndex[i]:_DeviceBGPTypeIndex[i+1]]
}

var _DeviceBGPTypeValues = []DeviceBGPType{0, 1, 2}

var _DeviceBGPTypeNameToValueMap = map[string]DeviceBGPType{
	_DeviceBGPTypeName[0:4]:   0,
	_DeviceBGPTypeName[4:10]:  1,
	_DeviceBGPTypeName[10:22]: 2,
}

// DeviceBGPTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DeviceBGPTypeString(s string) (DeviceBGPType, error) {
	if val, ok := _DeviceBGPTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DeviceBGPType values", s)
}

// DeviceBGPTypeValues returns all values of the enum
func DeviceBGPTypeValues() []DeviceBGPType {
	return _DeviceBGPTypeValues
}

// IsADeviceBGPType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DeviceBGPType) IsADeviceBGPType() bool {
	for _, v := range _DeviceBGPTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
