package convert

import (
	"fmt"
	"net"
	"time"
)

func MillisecondsF32ToDuration(ms float32) time.Duration {
	// scale to nanoseconds before conversion to duration to minimise conversion loss
	const nanosPerMicro = 1e6
	return time.Duration(nanosPerMicro*ms) * time.Nanosecond
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
