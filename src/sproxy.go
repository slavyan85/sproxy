/*
author Vyacheslav Kryuchenko
*/
package main

import (
	socks5 "github.com/armon/go-socks5"
	"flag"
	"log"
)

var (
	bindAddress *string = flag.String("bind", ":8899", "binding host:port")
)

func main() {
	flag.Parse()
	proxyConfig := &socks5.Config{}
	server, err := socks5.New(proxyConfig)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Start s-proxy with bind %s", *bindAddress)
	log.Fatal(server.ListenAndServe("tcp", *bindAddress))
}
