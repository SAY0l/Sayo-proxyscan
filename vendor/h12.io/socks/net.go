package socks

import (
	"bytes"
	"fmt"
	"net"
	"strconv"
	"time"
)

type requestBuilder struct {
	bytes.Buffer
}

func (b *requestBuilder) add(data ...byte) {
	_, _ = b.Write(data)
}

func (c *config) sendReceive(conn net.Conn, req []byte) (resp []byte, err error) {
	if c.Timeout > 0 {
		if err := conn.SetWriteDeadline(time.Now().Add(c.Timeout)); err != nil {
			return nil, err
		}
	}
	_, err = conn.Write(req)
	if err != nil {
		return
	}
	resp, err = c.readAll(conn)
	return
}

func (c *config) readAll(conn net.Conn) (resp []byte, err error) {
	resp = make([]byte, 1024)
	if c.Timeout > 0 {
		if err := conn.SetReadDeadline(time.Now().Add(c.Timeout)); err != nil {
			return nil, err
		}
	}
	n, err := conn.Read(resp)
	resp = resp[:n]
	return
}

func lookupIPv4(host string) (net.IP, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	for _, ip := range ips {
		ipv4 := ip.To4()
		if ipv4 == nil {
			continue
		}
		return ipv4, nil
	}
	return nil, fmt.Errorf("no IPv4 address found for host: %s", host)
}

func splitHostPort(addr string) (host string, port uint16, err error) {
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return "", 0, err
	}
	portInt, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return "", 0, err
	}
	port = uint16(portInt)
	return
}
