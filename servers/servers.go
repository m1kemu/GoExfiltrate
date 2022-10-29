package servers

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type ControlServer struct {
	BindIP   string
	BindPort string
	Key      []byte
}

type TCPServer struct {
	BindIP   string
	BindPort string
	FilePath string
	Key      []byte
}

type UDPServer struct {
	BindIP   string
	BindPort string
	FilePath string
	Key      []byte
}

func ControlServerStatus(w http.ResponseWriter, req *http.Request) {
	/*
	   ControlClient will send the following data:
	   - 'User-Agent' header: Client user-agent with GoExfiltrate data baked in as a basic signaturing mechanism
	   - 'GET' param 'src_id_b64'
	   - 'GET' param 'file_name_b64'
	*/

	// Confirm that src_id and file_name exist

	// Respond with 404 if they don't, 200 and correct status data if they do
}

func WriteControlData(file_name string, chunk_size int, src_id string, chunk_num int, encryption_mode string, data_chunk []byte) {
	// If initial message (chunk num == 0), create DB entry
	//      - file_name, src_id, chunk_size, src_ip, encryption_method, current_chunk_num, total_chunk_num, finished

	// Print status to console
	// Insert data into file
	// Update db
}

func ControlServerHandler(w http.ResponseWriter, req *http.Request) {
	/*
	   ControlClient will send the following data:
	   - 'User-Agent' header: Client user-agent with GoExfiltrate data baked in as a basic signaturing mechanism
	   - POST data: [file_name_b64]--[chunk_size_b64]--[src_id_b64]--[chunk_num_b64]--[total_chunk_num_b64]--[encryption_mode_b64]--[data_chunk_b64]
	*/

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}

	// Validate data is coming from a client

	// Send properly formatted data to Data Writer
}

func (server *ControlServer) Listen() {
	http.HandleFunc("/exfiltrate", ControlServerHandler)
	http.HandleFunc("/status", ControlServerStatus)

	http.ListenAndServe(server.BindIP+":"+server.BindPort, nil)
}

func (server *TCPServer) Listen() {
	listener, err := net.Listen("tcp", server.BindIP+":"+server.BindPort)

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
	sock, err := net.ResolveUDPAddr("udp4", server.BindIP+":"+server.BindPort)

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
