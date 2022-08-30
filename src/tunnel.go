package hp

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	_ "embed"

	gossh "golang.org/x/crypto/ssh"
)

// publicKeys returns auth with given private key
func publicKey(key []byte) (gossh.AuthMethod, error) {
	keys, err := gossh.ParsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("Cannot extract SSH public key from private")
	}
	return gossh.PublicKeys(keys), nil
}

func reversessh(hp *HolePunch) {

	publicKeyAuth, err := publicKey(hp.privateKey)
	if err != nil {
		log.Fatalln(fmt.Printf("invalid private key: %s", err))
		os.Exit(1)
	}

	sshConfig := &gossh.ClientConfig{
		User:            hp.RemoteEndpoint.User,
		Auth:            []gossh.AuthMethod{publicKeyAuth},
		HostKeyCallback: gossh.InsecureIgnoreHostKey(),
	}

	// Connect to SSH remote server using serverEndpoint
	serverConn, err := gossh.Dial("tcp", hp.RemoteEndpoint.Address(), sshConfig)
	if err != nil {
		log.Printf("[remote endpoint] %s", err)
		os.Exit(1)
	}
	log.Printf("[remote endpoint] connected %s", hp.RemoteEndpoint.Address())

	// Listen on remote server port
	listener, err := serverConn.Listen("tcp", hp.TunnelEndpoint.Address())
	if err != nil {
		log.Printf("[tunnel endpoint] %s", err)
		os.Exit(1)
	}
	log.Printf("[tunnel endpoint] created %s <--- %s", hp.TunnelEndpoint.Address(), hp.RemoteEndpoint.Address())
	defer listener.Close()

	log.Printf("[local endpoint] creating %s", hp.LocalEndpoint.Address())

	// handle incoming connections on reverse forwarded tunnel
	for {
		// Open a (local) connection to localEndpoint whose content will be forwarded so serverEndpoint
		local, err := net.Dial("tcp", hp.LocalEndpoint.Address())
		if err != nil {
			log.Printf("[local endpoint] %s", err)
			os.Exit(1)
		}

		client, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		tunnel(client, local)
	}

}

func tunnel(client net.Conn, remote net.Conn) {
	defer client.Close()
	chDone := make(chan bool)

	// Start remote -> local data transfer
	go func() {
		_, err := io.Copy(client, remote)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy remote->local: %s", err))
		}
		chDone <- true
	}()

	// Start local -> remote data transfer
	go func() {
		_, err := io.Copy(remote, client)
		if err != nil {
			log.Println(fmt.Sprintf("error while copy local->remote: %s", err))
		}
		chDone <- true
	}()

	<-chDone
}
