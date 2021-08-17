// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _TimeFormatName = "UTCLocal"

var _TimeFormatIndex = [...]uint8{0, 3, 8}

func (i TimeFormat) String() string {
	if i < 0 || i >= TimeFormat(len(_TimeFormatIndex)-1) {
		return fmt.Sprintf("TimeFormat(%d)", i)
	}
	return _TimeFormatName[_TimeFormatIndex[i]:_TimeFormatIndex[i+1]]
}

var _TimeFormatValues = []TimeFormat{0, 1}

var _TimeFormatNameToValueMap = map[string]TimeFormat{
	_TimeFormatName[0:3]: 0,
	_TimeFormatName[3:8]: 1,
}

// TimeFormatString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func TimeFormatString(s string) (TimeFormat, error) {
	if val, ok := _TimeFormatNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to TimeFormat values", s)
}

// TimeFormatValues returns all values of the enum
func TimeFormatValues() []TimeFormat {
	return _TimeFormatValues
}

// IsATimeFormat returns "true" if the value is listed in the enum definition. "false" otherwise
func (i TimeFormat) IsATimeFormat() bool {
	for _, v := range _TimeFormatValues {
		if i == v {
			return true
		}
	}
	return false
}
