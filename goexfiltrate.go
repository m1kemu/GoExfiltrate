package goexfiltrate

import (
	"fmt"
)

type Listener interface {
	Listen()
	bind_ip string
	bind_port string
	dest_file_path string
	exfil_method string
	exfil_domain string
}

