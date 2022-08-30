package hp

import "fmt"

type HolePunch struct {
	Shell string

	RemoteEndpoint SSHEndpoint
	LocalEndpoint  SSHEndpoint
	TunnelEndpoint SSHEndpoint

	TunnelAddress string
	TunnelPort    int

	privateKey []byte
}

type SSHEndpoint struct {
	User string
	Host string
	Port string
}

func (endpoint *SSHEndpoint) Address() string {
	return fmt.Sprintf("%s:%s", endpoint.Host, endpoint.Port)
}
