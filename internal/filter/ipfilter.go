package filter

import (
	"fmt"
	"net"
)

// IPFilter matches lines that contain an IP address within a given CIDR range.
type IPFilter struct {
	cidr    string
	network *net.IPNet
}

// NewIPFilter creates a filter that matches lines containing an IP address
// belonging to the given CIDR block (e.g. "192.168.1.0/24").
func NewIPFilter(cidr string) (*IPFilter, error) {
	if cidr == "" {
		return nil, fmt.Errorf("%w: CIDR must not be empty", ErrInvalidConfig)
	}
	_, network, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, fmt.Errorf("%w: invalid CIDR %q: %v", ErrInvalidConfig, cidr, err)
	}
	return &IPFilter{cidr: cidr, network: network}, nil
}

// Match returns true if the line contains at least one IP address that falls
// within the configured CIDR range.
func (f *IPFilter) Match(line string) bool {
	// Walk through the line looking for candidate IP tokens.
	tokens := splitTokens(line)
	for _, tok := range tokens {
		ip := net.ParseIP(tok)
		if ip != nil && f.network.Contains(ip) {
			return true
		}
	}
	return false
}

// CIDR returns the configured CIDR string.
func (f *IPFilter) CIDR() string { return f.cidr }

// splitTokens splits a line on common non-IP characters to extract candidate
// tokens that could be IPv4 or IPv6 addresses.
func splitTokens(line string) []string {
	var tokens []string
	start := -1
	for i, ch := range line {
		if isIPChar(ch) {
			if start == -1 {
				start = i
			}
		} else {
			if start != -1 {
				tokens = append(tokens, line[start:i])
				start = -1
			}
		}
	}
	if start != -1 {
		tokens = append(tokens, line[start:])
	}
	return tokens
}

func isIPChar(ch rune) bool {
	return (ch >= '0' && ch <= '9') || (ch >= 'a' && ch <= 'f') ||
		(ch >= 'A' && ch <= 'F') || ch == '.' || ch == ':'
}
