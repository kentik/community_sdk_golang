// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _AggregateFunctionTypeName = "sumaveragepercentilemaxcompositeexponentmodulusgreaterThangreaterThanEqualslessThanlessThanEqualsequalsnotEquals"

var _AggregateFunctionTypeIndex = [...]uint8{0, 3, 10, 20, 23, 32, 40, 47, 58, 75, 83, 97, 103, 112}

func (i AggregateFunctionType) String() string {
	if i < 0 || i >= AggregateFunctionType(len(_AggregateFunctionTypeIndex)-1) {
		return fmt.Sprintf("AggregateFunctionType(%d)", i)
	}
	return _AggregateFunctionTypeName[_AggregateFunctionTypeIndex[i]:_AggregateFunctionTypeIndex[i+1]]
}

var _AggregateFunctionTypeValues = []AggregateFunctionType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

var _AggregateFunctionTypeNameToValueMap = map[string]AggregateFunctionType{
	_AggregateFunctionTypeName[0:3]:     0,
	_AggregateFunctionTypeName[3:10]:    1,
	_AggregateFunctionTypeName[10:20]:   2,
	_AggregateFunctionTypeName[20:23]:   3,
	_AggregateFunctionTypeName[23:32]:   4,
	_AggregateFunctionTypeName[32:40]:   5,
	_AggregateFunctionTypeName[40:47]:   6,
	_AggregateFunctionTypeName[47:58]:   7,
	_AggregateFunctionTypeName[58:75]:   8,
	_AggregateFunctionTypeName[75:83]:   9,
	_AggregateFunctionTypeName[83:97]:   10,
	_AggregateFunctionTypeName[97:103]:  11,
	_AggregateFunctionTypeName[103:112]: 12,
}

// AggregateFunctionTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func AggregateFunctionTypeString(s string) (AggregateFunctionType, error) {
	if val, ok := _AggregateFunctionTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to AggregateFunctionType values", s)
}

// AggregateFunctionTypeValues returns all values of the enum
func AggregateFunctionTypeValues() []AggregateFunctionType {
	return _AggregateFunctionTypeValues
}

// IsAAggregateFunctionType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i AggregateFunctionType) IsAAggregateFunctionType() bool {
	for _, v := range _AggregateFunctionTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
