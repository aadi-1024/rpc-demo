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
		data: make(map[int]string),
	}
	err := rpc.Register(&CacheRPC{})
	if err != nil {
		log.Println(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "127.0.0.1:7070")
	if err != nil {
		log.Println(err)
	}
	if err = http.Serve(l, nil); err != nil {
		log.Println(err)
	}
}
