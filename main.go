// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"duardochating/api"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
	//	"github.com/gorilla/mux"
)

var addr = flag.String("addr", ":8090", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	//	if r.URL.Path != "/" {
	//		http.Error(w, "Not found", http.StatusNotFound)
	//		return
	//	}
	switch url := r.URL.Path; url {
	case "/":
		http.ServeFile(w, r, "html/index.html")
	case "/css/main.css":
		http.ServeFile(w, r, "html/css/main.css")
	case "/js/main.js":
		http.ServeFile(w, r, "html/js/main.js")
	case "/api/myip":
		api.GetIp(w, r)
	default:
		notFound(w, r)
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

/*
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "API running")
    fmt.Println("Endpoint Hit: homePage")
}
*/

func notFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	lookup, err := net.LookupAddr(string([]rune(r.RemoteAddr)[:len(r.RemoteAddr)-6]))
	if err != nil {
		lookup = strings.Split(err.Error(), "")
	}
	fmt.Fprintf(w, "404 pagina não encontrada 😭\n\nIp request -> "+r.RemoteAddr+"\nReverse -> "+strings.Join(lookup, "")+"\nURL -> http://eduadro137.dev"+r.RequestURI+"\nUser Agent -> "+r.UserAgent()+"\nTime UTC -> "+time.Now().UTC().Local().Format(time.UnixDate))
	log.Println("Endpoint Hit: notFound " + r.RemoteAddr + " request:" + r.RequestURI + " User Agent: " + r.UserAgent())
}

func main() {

	//	rota := mux.NewRouter().StrictSlash(true)

	//	rota.NotFoundHandler = http.HandlerFunc(notFound)
	//	rota.HandleFunc("/",homePage).Methods("GET")

	//	log.Fatal(http.ListenAndServe(":1337", rota))

	// main socket server
	flag.Parse()
	hub := newHub()
	go hub.run()

	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
