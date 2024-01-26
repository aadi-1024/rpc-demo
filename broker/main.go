package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
)

const (
	Port      = 8080
	MailPort  = 6060
	CachePort = 7070
)

var app *AppConfig

func main() {
	var err error
	app = &AppConfig{}

	app.CacheConn, err = rpc.DialHTTP("tcp", fmt.Sprintf("127.0.0.1:%v", CachePort))
	if err != nil {
		log.Println(err)
	}

	srv := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%v", Port),
		Handler: NewRouter(),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
