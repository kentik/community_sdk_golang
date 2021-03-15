// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _ChartViewTypeName = "stackedArealinestackedBarbarpiesankeytablematrix"

var _ChartViewTypeIndex = [...]uint8{0, 11, 15, 25, 28, 31, 37, 42, 48}

func (i ChartViewType) String() string {
	if i < 0 || i >= ChartViewType(len(_ChartViewTypeIndex)-1) {
		return fmt.Sprintf("ChartViewType(%d)", i)
	}
	return _ChartViewTypeName[_ChartViewTypeIndex[i]:_ChartViewTypeIndex[i+1]]
}

var _ChartViewTypeValues = []ChartViewType{0, 1, 2, 3, 4, 5, 6, 7}

var _ChartViewTypeNameToValueMap = map[string]ChartViewType{
	_ChartViewTypeName[0:11]:  0,
	_ChartViewTypeName[11:15]: 1,
	_ChartViewTypeName[15:25]: 2,
	_ChartViewTypeName[25:28]: 3,
	_ChartViewTypeName[28:31]: 4,
	_ChartViewTypeName[31:37]: 5,
	_ChartViewTypeName[37:42]: 6,
	_ChartViewTypeName[42:48]: 7,
}

// ChartViewTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ChartViewTypeString(s string) (ChartViewType, error) {
	if val, ok := _ChartViewTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ChartViewType values", s)
}

// ChartViewTypeValues returns all values of the enum
func ChartViewTypeValues() []ChartViewType {
	return _ChartViewTypeValues
}

// IsAChartViewType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ChartViewType) IsAChartViewType() bool {
	for _, v := range _ChartViewTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
