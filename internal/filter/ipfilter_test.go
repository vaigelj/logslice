package filter

import (
	"testing"
)

func TestNewIPFilter_EmptyCIDR(t *testing.T) {
	_, err := NewIPFilter("")
	if err == nil {
		t.Fatal("expected error for empty CIDR")
	}
}

func TestNewIPFilter_InvalidCIDR(t *testing.T) {
	_, err := NewIPFilter("not-a-cidr")
	if err == nil {
		t.Fatal("expected error for invalid CIDR")
	}
}

func TestNewIPFilter_Valid(t *testing.T) {
	f, err := NewIPFilter("10.0.0.0/8")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if f.CIDR() != "10.0.0.0/8" {
		t.Errorf("expected CIDR 10.0.0.0/8, got %s", f.CIDR())
	}
}

func TestIPFilter_Match_InRange(t *testing.T) {
	f, _ := NewIPFilter("192.168.1.0/24")
	if !f.Match("client connected from 192.168.1.42 port 5000") {
		t.Error("expected match for IP in range")
	}
}

func TestIPFilter_Match_OutOfRange(t *testing.T) {
	f, _ := NewIPFilter("192.168.1.0/24")
	if f.Match("client connected from 10.0.0.1 port 5000") {
		t.Error("expected no match for IP out of range")
	}
}

func TestIPFilter_Match_NoIP(t *testing.T) {
	f, _ := NewIPFilter("192.168.0.0/16")
	if f.Match("no ip address here at all") {
		t.Error("expected no match when line has no IP")
	}
}

func TestIPFilter_Match_MultipleIPs_OneInRange(t *testing.T) {
	f, _ := NewIPFilter("172.16.0.0/12")
	line := "route from 10.0.0.1 via 172.16.5.10 to 8.8.8.8"
	if !f.Match(line) {
		t.Error("expected match when one of multiple IPs is in range")
	}
}

func TestIPFilter_Match_IPv6(t *testing.T) {
	f, _ := NewIPFilter("::1/128")
	if !f.Match("loopback request from ::1") {
		t.Error("expected match for IPv6 loopback")
	}
}

func TestIPFilter_InChain(t *testing.T) {
	ipf, _ := NewIPFilter("10.0.0.0/8")
	chain, _ := NewChain(ipf)
	if !chain.Match("error from 10.1.2.3: timeout") {
		t.Error("expected chain match")
	}
	if chain.Match("error from 8.8.8.8: timeout") {
		t.Error("expected chain no-match")
	}
}
