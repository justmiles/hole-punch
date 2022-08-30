package hp

import (
	"fmt"
	"io"
	"os/exec"

	_ "embed"

	"github.com/creack/pty"
	"github.com/gliderlabs/ssh"
	gossh "golang.org/x/crypto/ssh"
)

// TODO: handle PTY for windows with custom _windows.sh

// go func() {
// 	for win := range winCh {
// 		setWinsize(f, win.Width, win.Height)
// 	}
// }()

func Run(opts ...HolePunchOption) error {

	const (
		defaultShell = "sh"
		defaultHost  = "127.0.0.1"
		defaultPort  = "22"
		defaultUser  = "hp"
	)

	hp := &HolePunch{
		Shell: defaultShell,
		LocalEndpoint: SSHEndpoint{
			User: defaultUser,
			Host: defaultHost,
			Port: defaultPort,
		},
		RemoteEndpoint: SSHEndpoint{
			User: defaultUser,
			Host: defaultHost,
			Port: defaultPort,
		},
		TunnelEndpoint: SSHEndpoint{
			User: defaultUser,
			Host: defaultHost,
			Port: defaultPort,
		},
	}

	// Apply hp options
	for _, opt := range opts {
		opt(hp)
	}

	ssh.Handle(func(s ssh.Session) {
		cmd := exec.Command(hp.Shell)

		ptyReq, _, isPty := s.Pty()
		if isPty {
			cmd.Env = append(cmd.Env, fmt.Sprintf("TERM=%s", ptyReq.Term))
			f, err := pty.Start(cmd)
			if err != nil {
				panic(err)
			}
			go func() {
				io.Copy(f, s) // stdin
			}()
			io.Copy(s, f) // stdout
			cmd.Wait()
		} else {
			io.WriteString(s, "No PTY requested.\n")
			s.Exit(1)
		}
	})

	private, err := gossh.ParsePrivateKey(hp.privateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	// limit authentication to local ssh with _only_ this key
	publicKeyOption := ssh.PublicKeyAuth(func(ctx ssh.Context, key ssh.PublicKey) bool {
		return ssh.KeysEqual(private.PublicKey(), key)
	})

	go reversessh(hp)

	return ssh.ListenAndServe(hp.LocalEndpoint.Address(), nil, publicKeyOption)

}
