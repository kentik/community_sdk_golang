// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _FastDataTypeName = "AutoFastFull"

var _FastDataTypeIndex = [...]uint8{0, 4, 8, 12}

func (i FastDataType) String() string {
	if i < 0 || i >= FastDataType(len(_FastDataTypeIndex)-1) {
		return fmt.Sprintf("FastDataType(%d)", i)
	}
	return _FastDataTypeName[_FastDataTypeIndex[i]:_FastDataTypeIndex[i+1]]
}

var _FastDataTypeValues = []FastDataType{0, 1, 2}

var _FastDataTypeNameToValueMap = map[string]FastDataType{
	_FastDataTypeName[0:4]:  0,
	_FastDataTypeName[4:8]:  1,
	_FastDataTypeName[8:12]: 2,
}

// FastDataTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FastDataTypeString(s string) (FastDataType, error) {
	if val, ok := _FastDataTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to FastDataType values", s)
}

// FastDataTypeValues returns all values of the enum.
func FastDataTypeValues() []FastDataType {
	return _FastDataTypeValues
}

// IsAFastDataType returns "true" if the value is listed in the enum definition. "false" otherwise.
func (i FastDataType) IsAFastDataType() bool {
	for _, v := range _FastDataTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
