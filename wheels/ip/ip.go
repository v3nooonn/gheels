package ip

import (
	"fmt"
	"net"
	"net/http"
	"strings"
)

const (
	unknown string = "unknown"
	local   string = "127.0.0.1"
)

func Parse(r *http.Request) (string, error) {
	forwards := r.Header.Get("X-Forwarded-For")
	if len(forwards) > 0 {
		ips := strings.Split(forwards, ",")
		ip := net.ParseIP(ips[len(ips)-1])
		if ip != nil {
			return ip.String(), nil
		}
	}

	ipStr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", fmt.Errorf("parsing remote address occurred an error: %s", err.Error())
	}

	ip := net.ParseIP(ipStr)
	if ip != nil {
		ip := ip.String()
		if ip == "::1" {
			return local, nil
		}
		return ip, nil
	}

	return unknown, nil
}
