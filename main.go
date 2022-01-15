package main

import "goexfiltrate/servers"

func main() {
	/*
	server := servers.TCPServer {
		BindIP: "127.0.0.1",
		BindPort: "4444",
		FilePath: "/tmp/tcp_output.out",
	}
	*/

	server := servers.UDPServer {
		BindIP: "127.0.0.1",
		BindPort: "4444",
		FilePath: "/tmp/udp_output.out",
	}

	server.Listen()
}
