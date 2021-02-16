// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _DeviceSubtypeName = "routercisco_asapaloaltosilverpeakmplsviptelapfe_syslogsyslogmerakiistioios_xrcisco_zone_based_firewallcisco_nbarcisco_asa_syslogadvanced_sflowa10_cgnkprobenprobeaws_subnetazure_subnetgcp_subnetkappaibm_subnet"

var _DeviceSubtypeIndex = [...]uint8{0, 6, 15, 23, 33, 37, 44, 54, 60, 66, 71, 77, 102, 112, 128, 142, 149, 155, 161, 171, 183, 193, 198, 208}

func (i DeviceSubtype) String() string {
	if i < 0 || i >= DeviceSubtype(len(_DeviceSubtypeIndex)-1) {
		return fmt.Sprintf("DeviceSubtype(%d)", i)
	}
	return _DeviceSubtypeName[_DeviceSubtypeIndex[i]:_DeviceSubtypeIndex[i+1]]
}

var _DeviceSubtypeValues = []DeviceSubtype{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22}

var _DeviceSubtypeNameToValueMap = map[string]DeviceSubtype{
	_DeviceSubtypeName[0:6]:     0,
	_DeviceSubtypeName[6:15]:    1,
	_DeviceSubtypeName[15:23]:   2,
	_DeviceSubtypeName[23:33]:   3,
	_DeviceSubtypeName[33:37]:   4,
	_DeviceSubtypeName[37:44]:   5,
	_DeviceSubtypeName[44:54]:   6,
	_DeviceSubtypeName[54:60]:   7,
	_DeviceSubtypeName[60:66]:   8,
	_DeviceSubtypeName[66:71]:   9,
	_DeviceSubtypeName[71:77]:   10,
	_DeviceSubtypeName[77:102]:  11,
	_DeviceSubtypeName[102:112]: 12,
	_DeviceSubtypeName[112:128]: 13,
	_DeviceSubtypeName[128:142]: 14,
	_DeviceSubtypeName[142:149]: 15,
	_DeviceSubtypeName[149:155]: 16,
	_DeviceSubtypeName[155:161]: 17,
	_DeviceSubtypeName[161:171]: 18,
	_DeviceSubtypeName[171:183]: 19,
	_DeviceSubtypeName[183:193]: 20,
	_DeviceSubtypeName[193:198]: 21,
	_DeviceSubtypeName[198:208]: 22,
}

// DeviceSubtypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DeviceSubtypeString(s string) (DeviceSubtype, error) {
	if val, ok := _DeviceSubtypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DeviceSubtype values", s)
}

// DeviceSubtypeValues returns all values of the enum
func DeviceSubtypeValues() []DeviceSubtype {
	return _DeviceSubtypeValues
}

// IsADeviceSubtype returns "true" if the value is listed in the enum definition. "false" otherwise
func (i DeviceSubtype) IsADeviceSubtype() bool {
	for _, v := range _DeviceSubtypeValues {
		if i == v {
			return true
		}
	}
	return false
}
