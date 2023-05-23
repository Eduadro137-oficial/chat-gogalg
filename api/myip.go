package api

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
)

type Ip struct {
	Ipv4    string `json:"ipv4,omitempty"`
	Reverse string `json:"reverse,omitempty"`
}

func GetIp(w http.ResponseWriter, r *http.Request) {
	r.Header.Get("X-REAL-IP")

	lookup, err := net.LookupAddr(string([]rune(r.RemoteAddr)[:len(r.RemoteAddr)-6]))
	if err != nil {
		lookup = strings.Split(err.Error(), "")
	}

	RemoteAddr := string(r.RemoteAddr)
	ip := Ip{
		Ipv4:    string([]rune(RemoteAddr)[:len(RemoteAddr)-6]),
		Reverse: strings.Join(lookup, ""),
	}

	w.Header().Set("Content-Type", "application/json")
	jsonapi, _ := json.Marshal(ip)
	fmt.Fprint(w, string(jsonapi))

}
