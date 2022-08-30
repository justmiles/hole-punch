package hp

import (
	"net"
	"os/exec"
	"strings"
)

type HolePunchOption func(*HolePunch)

func WithShell(shell string) HolePunchOption {
	return func(hp *HolePunch) {
		hp.Shell = shell
	}
}

func WithPrivateKey(key []byte) HolePunchOption {
	return func(hp *HolePunch) {
		hp.privateKey = key
	}
}

func WithLocalEndpoint(endpoint string) HolePunchOption {
	return func(hp *HolePunch) {
		host, port, _ := net.SplitHostPort(endpoint)
		if host != "" {
			hp.LocalEndpoint.Host = host
		}

		if port != "" {
			hp.LocalEndpoint.Port = port
		}
	}
}

func WithRemoteEndpoint(endpoint string) HolePunchOption {
	return func(hp *HolePunch) {

		// check for user in endpoint string
		s := strings.Split(endpoint, "@")
		if len(s) > 1 {
			hp.RemoteEndpoint.User = s[0]
			endpoint = s[1]
		}

		host, port, _ := net.SplitHostPort(endpoint)
		if host != "" {
			hp.RemoteEndpoint.Host = host
		}

		if port != "" {
			hp.RemoteEndpoint.Port = port
		}
	}
}
func WithTunnelEndpoint(endpoint string) HolePunchOption {
	return func(hp *HolePunch) {
		host, port, _ := net.SplitHostPort(endpoint)
		if host != "" {
			hp.TunnelEndpoint.Host = host
		}

		if port != "" {
			hp.TunnelEndpoint.Port = port
		}
	}
}

func DefaultShell() string {
	path, err := exec.LookPath("bash")
	if err == nil {
		return path
	}

	path, err = exec.LookPath("sh")
	if err == nil {
		return path
	}

	path, err = exec.LookPath("cmd.exe")
	if err == nil {
		return path
	}

	return ""
}
