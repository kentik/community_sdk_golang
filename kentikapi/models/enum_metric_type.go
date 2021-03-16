// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _MetricTypeName = "bytesin_bytesout_bytespacketsin_packetsout_packetstcp_retransmitperc_retransmitretransmits_inperc_retransmits_inout_of_order_inperc_out_of_order_infragmentsperc_fragmentsclient_latencyserver_latencyappl_latencyfpsunique_src_ipunique_dst_ip"

var _MetricTypeIndex = [...]uint8{0, 5, 13, 22, 29, 39, 50, 64, 79, 93, 112, 127, 147, 156, 170, 184, 198, 210, 213, 226, 239}

func (i MetricType) String() string {
	if i < 0 || i >= MetricType(len(_MetricTypeIndex)-1) {
		return fmt.Sprintf("MetricType(%d)", i)
	}
	return _MetricTypeName[_MetricTypeIndex[i]:_MetricTypeIndex[i+1]]
}

var _MetricTypeValues = []MetricType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19}

var _MetricTypeNameToValueMap = map[string]MetricType{
	_MetricTypeName[0:5]:     0,
	_MetricTypeName[5:13]:    1,
	_MetricTypeName[13:22]:   2,
	_MetricTypeName[22:29]:   3,
	_MetricTypeName[29:39]:   4,
	_MetricTypeName[39:50]:   5,
	_MetricTypeName[50:64]:   6,
	_MetricTypeName[64:79]:   7,
	_MetricTypeName[79:93]:   8,
	_MetricTypeName[93:112]:  9,
	_MetricTypeName[112:127]: 10,
	_MetricTypeName[127:147]: 11,
	_MetricTypeName[147:156]: 12,
	_MetricTypeName[156:170]: 13,
	_MetricTypeName[170:184]: 14,
	_MetricTypeName[184:198]: 15,
	_MetricTypeName[198:210]: 16,
	_MetricTypeName[210:213]: 17,
	_MetricTypeName[213:226]: 18,
	_MetricTypeName[226:239]: 19,
}

// MetricTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func MetricTypeString(s string) (MetricType, error) {
	if val, ok := _MetricTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to MetricType values", s)
}

// MetricTypeValues returns all values of the enum
func MetricTypeValues() []MetricType {
	return _MetricTypeValues
}

// IsAMetricType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i MetricType) IsAMetricType() bool {
	for _, v := range _MetricTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
