package main

import (
	"goexfiltrate/servers"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	control_server := servers.ControlServer{
		BindIP:   "127.0.0.1",
		BindPort: "4444",
		Key:      []byte("passwordpassword"),
	}

	tcp_server := servers.TCPServer{
		BindIP:   "127.0.0.1",
		BindPort: "8080",
		Key:      []byte("passwordpassword"),
	}

	go control_server.Listen()
	go tcp_server.Listen()
}
