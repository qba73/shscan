package shscan

import (
	"encoding/binary"
	"fmt"
	"net"
)

func VerifySocket(addr string) (bool, error) {
	ipAddr, err := net.ResolveIPAddr("ip4", addr)
	if err != nil {
		return false, fmt.Errorf("netscan: resolving IP Address %w", err)
	}
	conn, err := net.ListenIP("ip4:tcp", ipAddr)
	if err != nil {
		return false, fmt.Errorf("netscan: creating connection %w", err)
	}
	conn.Close()
	return true, nil
}

func GenerateHostRange(network string) ([]string, error) {
	_, ipNet, err := net.ParseCIDR(network)
	if err != nil {
		return nil, fmt.Errorf("netscan: generating host range %w", err)
	}
	mask := binary.BigEndian.Uint32(ipNet.Mask)
	start := binary.BigEndian.Uint32(ipNet.IP)
	end := (start & mask) | (mask ^ 0xffffffff)

	var hosts []string
	for i := start + 1; i <= end-1; i++ {
		ip := make(net.IP, 4)
		binary.BigEndian.PutUint32(ip, i)
		hosts = append(hosts, ip.String())
	}
	return hosts, nil
}
