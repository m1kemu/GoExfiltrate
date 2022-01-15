package servers

import (
	"fmt"
	"net"
	"bufio"
	"strings"
)

type TCPServer struct {
	BindIP	 string
	BindPort string
	FilePath string
}

type UDPServer struct {
	BindIP   string
	BindPort string
	FilePath string
}

func (server *TCPServer) Listen() {
	listener, err := net.Listen("tcp", server.BindIP + ":" + server.BindPort)

        fmt.Printf("[+] Started TCP server on port %v\n", server.BindPort)

        if err != nil {
                fmt.Println(err)
                return
        }
        defer listener.Close()

        conn, err := listener.Accept()
        if err != nil {
                fmt.Println(err)
                return
        }

        fmt.Printf("[+] Connection from %v received\n", conn.RemoteAddr())

        var final_data []byte

        data, _ := bufio.NewReader(conn).ReadBytes('\n')

        fmt.Printf("[*] Received %v bytes of data\n", len(data))

        final_data = append(final_data, data...)

        for len(data) == 1024 {
                fmt.Printf("[*] Received %v bytes of data\n", len(data))
                final_data = append(final_data, data...)
        }
}

func (server *UDPServer) Listen() {
	sock, err := net.ResolveUDPAddr("udp4", server.BindIP + ":" + server.BindPort)

	if err != nil {
                fmt.Printf("[!] Error binding UDP port: %v\n", err)
                return
        }

        conn, err := net.ListenUDP("udp4", sock)

        defer conn.Close()

        buffer := make([]byte, 1024)

        fmt.Printf("[+] Started UDP server on port %v\n", server.BindPort)

        var final_data []byte

        for {
                num_bytes, _, _ := conn.ReadFromUDP(buffer)

                if strings.TrimSpace(string(buffer[0:num_bytes])) == "STOPPOTS" {
                        break
                }

                final_data = append(final_data, buffer[0:num_bytes-1]...)
                fmt.Printf("[+] Recieved data: %v\n", buffer[0:num_bytes-1])
        }
}
