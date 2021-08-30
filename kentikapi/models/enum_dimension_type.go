// NOTE: modified by hand to reflect actual API enumerators
package models

import (
	"fmt"
)

const _DimensionTypeName = "AS_srcGeography_srcInterfaceID_srcPort_srcsrc_eth_macVLAN_srcIP_srcAS_dstGeography_dstInterfaceID_dstPort_dstdst_eth_macVLAN_dstIP_dstTopFlowProtoTrafficASTopTalkersInterfaceTopTalkersPortPortTalkersTopFlowsIPsrc_geo_regionsrc_geo_citydst_geo_regiondst_geo_cityRegionTopTalkersi_device_idi_device_site_namesrc_route_prefix_lensrc_route_lengthsrc_bgp_communitysrc_bgp_aspathsrc_nexthop_ipsrc_nexthop_asnsrc_second_asnsrc_third_asnsrc_proto_portdst_route_prefix_lendst_route_lengthdst_bgp_communitydst_bgp_aspathdst_nexthop_ipdst_nexthop_asndst_second_asndst_third_asndst_proto_portinet_familyTOStcp_flags"

var _DimensionTypeIndex = [...]uint16{0, 6, 19, 34, 42, 53, 61, 67, 73, 86, 101, 109, 120, 128, 134, 141, 146, 153, 165, 184, 199, 209, 223, 235, 249, 261, 277, 288, 306, 326, 342, 359, 373, 387, 402, 416, 429, 443, 463, 479, 496, 510, 524, 539, 553, 566, 580, 591, 594, 603}

func (i DimensionType) String() string {
	if i < 0 || i >= DimensionType(len(_DimensionTypeIndex)-1) {
		return fmt.Sprintf("DimensionType(%d)", i)
	}
	return _DimensionTypeName[_DimensionTypeIndex[i]:_DimensionTypeIndex[i+1]]
}

var _DimensionTypeValues = []DimensionType{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48}

var _DimensionTypeNameToValueMap = map[string]DimensionType{
	_DimensionTypeName[0:6]:     0,
	_DimensionTypeName[6:19]:    1,
	_DimensionTypeName[19:34]:   2,
	_DimensionTypeName[34:42]:   3,
	_DimensionTypeName[42:53]:   4,
	_DimensionTypeName[53:61]:   5,
	_DimensionTypeName[61:67]:   6,
	_DimensionTypeName[67:73]:   7,
	_DimensionTypeName[73:86]:   8,
	_DimensionTypeName[86:101]:  9,
	_DimensionTypeName[101:109]: 10,
	_DimensionTypeName[109:120]: 11,
	_DimensionTypeName[120:128]: 12,
	_DimensionTypeName[128:134]: 13,
	_DimensionTypeName[134:141]: 14,
	_DimensionTypeName[141:146]: 15,
	_DimensionTypeName[146:153]: 16,
	_DimensionTypeName[153:165]: 17,
	_DimensionTypeName[165:184]: 18,
	_DimensionTypeName[184:199]: 19,
	_DimensionTypeName[199:209]: 20,
	_DimensionTypeName[209:223]: 21,
	_DimensionTypeName[223:235]: 22,
	_DimensionTypeName[235:249]: 23,
	_DimensionTypeName[249:261]: 24,
	_DimensionTypeName[261:277]: 25,
	_DimensionTypeName[277:288]: 26,
	_DimensionTypeName[288:306]: 27,
	_DimensionTypeName[306:326]: 28,
	_DimensionTypeName[326:342]: 29,
	_DimensionTypeName[342:359]: 30,
	_DimensionTypeName[359:373]: 31,
	_DimensionTypeName[373:387]: 32,
	_DimensionTypeName[387:402]: 33,
	_DimensionTypeName[402:416]: 34,
	_DimensionTypeName[416:429]: 35,
	_DimensionTypeName[429:443]: 36,
	_DimensionTypeName[443:463]: 37,
	_DimensionTypeName[463:479]: 38,
	_DimensionTypeName[479:496]: 39,
	_DimensionTypeName[496:510]: 40,
	_DimensionTypeName[510:524]: 41,
	_DimensionTypeName[524:539]: 42,
	_DimensionTypeName[539:553]: 43,
	_DimensionTypeName[553:566]: 44,
	_DimensionTypeName[566:580]: 45,
	_DimensionTypeName[580:591]: 46,
	_DimensionTypeName[591:594]: 47,
	_DimensionTypeName[594:603]: 48,
}

// DimensionTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DimensionTypeString(s string) (DimensionType, error) {
	if val, ok := _DimensionTypeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to DimensionType values", s)
}

// DimensionTypeValues returns all values of the enum.
func DimensionTypeValues() []DimensionType {
	return _DimensionTypeValues
}

// IsADimensionType returns "true" if the value is listed in the enum definition. "false" otherwise.
func (i DimensionType) IsADimensionType() bool {
	for _, v := range _DimensionTypeValues {
		if i == v {
			return true
		}
	}
	return false
}
