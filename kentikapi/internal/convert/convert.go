package convert

import (
	"fmt"
	"net"
	"time"

	"github.com/AlekSi/pointer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func MillisecondsF32ToDuration(ms float32) time.Duration {
	// scale to nanoseconds before conversion to duration to minimise conversion loss
	const nanosPerMilli = 1e6
	return time.Duration(nanosPerMilli*ms) * time.Nanosecond
}

func StringsToIPs(ips []string) ([]net.IP, error) {
	result := make([]net.IP, 0, len(ips))
	for _, ipStr := range ips {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			return nil, fmt.Errorf("invalid IP: %v", ipStr)
		}

		result = append(result, ip)
	}
	return result, nil
}

func IPsToStrings(ips []net.IP) []string {
	result := make([]string, 0, len(ips))
	for _, ip := range ips {
		result = append(result, ip.String())
	}
	return result
}

func TimestampPtrToTime(ts *timestamppb.Timestamp) *time.Time {
	if ts == nil {
		return nil
	}

	return pointer.ToTime(ts.AsTime())
}

func TimePtrToTimestamp(t *time.Time) *timestamppb.Timestamp {
	if t == nil {
		return nil
	}

	return timestamppb.New(*t)
}
