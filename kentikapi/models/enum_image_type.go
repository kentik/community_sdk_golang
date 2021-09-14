// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _ImageTypeName = "pngjpegsvgpdfunknown"

var _ImageTypeIndex = [...]uint8{0, 3, 7, 10, 13, 20}

func (i ImageType) String() string {
	if i < 0 || i >= ImageType(len(_ImageTypeIndex)-1) {
		return fmt.Sprintf("ImageType(%d)", i)
	}
	return _ImageTypeName[_ImageTypeIndex[i]:_ImageTypeIndex[i+1]]
}

var _ImageTypeValues = []ImageType{0, 1, 2, 3, 4}

var _ImageTypeNameToValueMap = map[string]ImageType{
	_ImageTypeName[0:3]:   0,
	_ImageTypeName[3:7]:   1,
	_ImageTypeName[7:10]:  2,
	_ImageTypeName[10:13]: 3,
	_ImageTypeName[13:20]: 4,
}

// ImageTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ImageTypeString(s string) (ImageType, error) {
	if val, ok := _ImageTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ImageType values", s)
}

// ImageTypeValues returns all values of the enum.
func ImageTypeValues() []ImageType {
	return _ImageTypeValues
}

// IsAImageType returns "true" if the value is listed in the enum definition. "false" otherwise.
func (i ImageType) IsAImageType() bool {
	for _, v := range _ImageTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
