package models

import (
	"io/ioutil"
	"time"
)

// QueryObject is the root object describing QueryAPI Data/Chart/URL query
type QueryObject struct {
	Queries   []QueryArrayItem
	ImageType *ImageType // used in QueryChart
}

type QueryArrayItem struct {
	Query       Query
	Bucket      string
	BucketIndex *int
	IsOverlay   *bool // used in QueryChart, QueryURL
}

type ImageType int

const (
	ImageTypePNG ImageType = iota // "png"
	ImageTypeJPG                  // "jpeg"
	ImageTypeSVG                  // "svg"
	ImageTypePDF                  // "pdf"
	ImageTypeUnknown
)

type Query struct {
	Metric          MetricType
	Dimension       []DimensionType
	FiltersObj      *Filters
	SavedFilters    []interface{}  // []SavedFilter; To be implemented in SavedFiltersAPI
	MatrixBy        []string       // DimensionType or custom dimension, required non-nil even if no elements
	CIDR            *int           // valid: number 0..32
	CIDR6           *int           // valid: number 0..128
	PPSThreshold    *int           // valid: number > 0
	TopX            int            // valid: number 1..40, default 8
	Depth           int            // valid: number 25..250, default 100
	FastData        FastDataType   // default FastDataTypeAuto
	TimeFormat      TimeFormat     // default TimeFormatUTC
	HostnameLookup  bool           // default True
	LookbackSeconds int            // default 3600, value != 0 overrides "StartingTime" and "EndingTime"
	StartingTime    *time.Time     // alternative with "LookbackSeconds"
	EndingTime      *time.Time     // alternative with "LookbackSeconds"
	AllSelected     *bool          // overrides "DeviceName" if true (makes it ignored)
	DeviceName      []string       // alternative with "AllSelected"; but required non-nil even if no elements
	Descriptor      string         // default "", only used when Dimension is "Traffic"
	Aggregates      []Aggregate    // if empty, will be auto-filled based on "Metric" field
	Outsort         *string        // name of aggregate object, required when more than 1 objects on "Aggregates" list
	QueryTitle      string         // default "", only used in QueryChart
	VizType         *ChartViewType // only used in QueryChart, QueryURL
	ShowOverlay     *bool          // only used in QueryChart, QueryURL
	OverlayDay      *int           // only used in QueryChart, QueryURL
	SyncAxes        *bool          // only used in QueryChart, QueryURL
}

// NewQuery creates a Query with all required fields set
func NewQuery(metric MetricType, dimension []DimensionType) *Query {
	return &Query{
		Metric:          metric,
		Dimension:       dimension,
		TopX:            8,
		Depth:           100,
		FastData:        FastDataTypeAuto,
		TimeFormat:      TimeFormatUTC,
		HostnameLookup:  true,
		LookbackSeconds: 3600,
		DeviceName:      []string{},
		MatrixBy:        []string{},
	}
}

type Aggregate struct {
	Name       string
	Column     string
	Fn         AggregateFunctionType
	SampleRate int   // default 1
	Rank       *int  // valid: number 5..99; only used when Fn == Percentile
	Raw        *bool // required for chart queries
}

// NewAggregate creates an Aggregate with all required fields set
func NewAggregate(name string, column string, fn AggregateFunctionType) Aggregate {
	return Aggregate{
		Name:       name,
		Column:     column,
		Fn:         fn,
		SampleRate: 1,
	}
}

type MetricType int

const (
	MetricTypeBytes             MetricType = iota // "bytes"
	MetricTypeInBytes                             // "in_bytes"
	MetricTypeOutBytes                            // "out_bytes"
	MetricTypePackets                             // "packets"
	MetricTypeInPackets                           // "in_packets"
	MetricTypeOutPackets                          // "out_packets"
	MetricTypeTCPRetransmit                       // "tcp_retransmit"
	MetricTypePercRetransmit                      // "perc_retransmit"
	MetricTypeRetransmitsIn                       // "retransmits_in"
	MetricTypePercRetransmitsIn                   // "perc_retransmits_in"
	MetricTypeOutOfOrderIn                        // "out_of_order_in"
	MetricTypePercOutOfOrderIn                    // "perc_out_of_order_in"
	MetricTypeFragments                           // "fragments"
	MetricTypePercFragmets                        // "perc_fragments"
	MetricTypeClientLatency                       // "client_latency"
	MetricTypeServerLatency                       // "server_latency"
	MetricTypeApplLatency                         // "appl_latency"
	MetricTypeFPS                                 // "fps"
	MetricTypeUniqueSrcIP                         // "unique_src_ip"
	MetricTypeUniqueDstIP                         // "unique_dst_ip"
)

type DimensionType int

const (
	DimensionTypeASSrc               DimensionType = iota // "AS_src"
	DimensionTypeGeographySrc                             // "Geography_src"
	DimensionTypeInterfaceIDSrc                           // "InterfaceID_src"
	DimensionTypePortSrc                                  // "Port_src"
	DimensionTypeSrcEthMac                                // "src_eth_mac"
	DimensionTypeVLANSRC                                  // "VLAN_src"
	DimensionTypeIPSrc                                    // "IP_src"
	DimensionTypeASDst                                    // "AS_dst"
	DimensionTypeGeographyDst                             // "Geography_dst"
	DimensionTypeInterfaceIDDst                           // "InterfaceID_dst"
	DimensionTypePortDst                                  // "Port_dst"
	DimensionTypeDstEthMac                                // "dst_eth_mac"
	DimensionTypeVLANDst                                  // "VLAN_dst"
	DimensionTypeIPDst                                    // "IP_dst"
	DimensionTypeTopFlow                                  // "TopFlow"
	DimensionTypeProto                                    // "Proto"
	DimensionTypeTraffic                                  // "Traffic"
	DimensionTypeASTopTalkers                             // "ASTopTalkers"
	DimensionTypeInterfaceTopTalkers                      // "InterfaceTopTalkers"
	DimensionTypePortPortTalkers                          // "PortPortTalkers"
	DimensionTypeTopFlowsIP                               // "TopFlowsIP"
	DimensionTypeSrcGeoRegion                             // "src_geo_region"
	DimensionTypeSrcGeoCity                               // "src_geo_city"
	DimensionTypeDstGeoRegion                             // "dst_geo_region"
	DimensionTypeDstGeoCity                               // "dst_geo_city"
	DimensionTypeRegionTopTalkers                         // "RegionTopTalkers"
	DimensionTypeIDeviceID                                // "i_device_id"
	DimensionTypeIDeviceSiteName                          // "i_device_site_name"
	DimensionTypeSrcRoutePrefixLen                        // "src_route_prefix_len"
	DimensionTypeSrcRouteLength                           // "src_route_length"
	DimensionTypeSrcBGPCommunity                          // "src_bgp_community"
	DimensionTypeSrcBGPAspath                             // "src_bgp_aspath"
	DimensionTypeSrcNexthopIP                             // "src_nexthop_ip"
	DimensionTypeSrcNexthopASN                            // "src_nexthop_asn"
	DimensionTypeSrcSecondASN                             // "src_second_asn"
	DimensionTypeSrcThirdASN                              // "src_third_asn"
	DimensionTypeSrcProtoPort                             // "src_proto_port"
	DimensionTypeDstRoutePrefixLen                        // "dst_route_prefix_len"
	DimensionTypeDstRouteLength                           // "dst_route_length"
	DimensionTypeDstBGPCommunity                          // "dst_bgp_community"
	DimensionTypeDstBGPAspath                             // "dst_bgp_aspath"
	DimensionTypeDstNexthopIP                             // "dst_nexthop_ip"
	DimensionTypeDstNexthopASN                            // "dst_nexthop_asn"
	DimensionTypeDstSecondASN                             // "dst_second_asn"
	DimensionTypeDstThirdASN                              // "dst_third_asn"
	DimensionTypeDstProtoPort                             // "dst_proto_port"
	DimensionTypeINetFamily                               // "inet_family"
	DimensionTypeTOS                                      // "TOS"
	DimensionTypeTCPFlags                                 // "tcp_flags"
)

type FastDataType int

const (
	FastDataTypeAuto FastDataType = iota // "Auto"
	FastDataTypeFast                     // "Fast"
	FastDataTypeFull                     // "Full"
)

type TimeFormat int

const (
	TimeFormatUTC   TimeFormat = iota // "UTC"
	TimeFormatLocal                   // "Local"
)

type ChartViewType int

const (
	ChartViewTypeStackedArea ChartViewType = iota // "stackedArea"
	ChartViewTypeLine                             // "line"
	ChartViewTypeStackedBar                       // "stackedBar"
	ChartViewTypeBar                              // "bar"
	ChartViewTypePie                              // "pie"
	ChartViewTypeSankey                           // "sankey"
	ChartViewTypeTable                            // "table"
	ChartViewTypeMatrix                           // "matrix"
)

type AggregateFunctionType int

const (
	AggregateFunctionTypeSum               AggregateFunctionType = iota // "sum"
	AggregateFunctionTypeAverage                                        // "average"
	AggregateFunctionTypePercentile                                     // "percentile"
	AggregateFunctionTypeMax                                            // "max"
	AggregateFunctionTypeComposite                                      // "composite"
	AggregateFunctionTypeExponent                                       // "exponent"
	AggregateFunctionTypeModulus                                        // "modulus"
	AggregateFunctionTypeGreaterThan                                    // "greaterThan"
	AggregateFunctionTypeGreaterThanEquals                              // "greaterThanEquals"
	AggregateFunctionTypeLessThan                                       // "lessThan"
	AggregateFunctionTypeLessThanEquals                                 // "lessThanEquals"
	AggregateFunctionTypeEquals                                         // "equals"
	AggregateFunctionTypeNotEquals                                      // "notEquals"
)

type QueryDataResult struct {
	Results []interface{} // contents depend on used query object
}

type QueryURLResult struct {
	URL string // URL to Kentik Explorer webpage that was generated for given query
}

type QueryChartResult struct {
	ImageType ImageType
	ImageData []byte // raw chart image binary data that can be directly dumped into a file
}

func (r QueryChartResult) SaveImageAs(filename string) error {
	return ioutil.WriteFile(filename, r.ImageData, 0644)
}
