package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

var userStore *Store

func main() {
	userStore = &Store{
		data: make(map[string]string),
	}
	err := rpc.Register(&CacheRPC{})
	if err != nil {
		log.Println(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "0.0.0.0:7070")
	if err != nil {
		log.Println(err)
	}
	if err = http.Serve(l, nil); err != nil {
		log.Println(err)
	}
}
