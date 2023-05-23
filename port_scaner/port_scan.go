package api

import (
	"net/http"
	"fmt"
    "net"
    "time"
)

type Host struct {
	State       string
	Address     string
	AddressType string
	Hostnames   []string
	Ports       []int32
	// contains filtered or unexported fields
}

func GetPorts (w http.ResponseWriter, r *http.Request){
	net.TCPConn
}