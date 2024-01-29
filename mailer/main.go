package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/smtp"
)

const PORT = 6060

var appConfig *Config

func main() {
	appConfig = &Config{}

	conn, err := rpc.DialHTTP("tcp", "cache:7070") //connect to the cache rpc service
	if err != nil {
		log.Fatalln(err)
	}

	auth := smtp.PlainAuth("", "user@example.com", "password", "mailhog")

	appConfig.SMTPAuth = auth
	appConfig.SMTPAddr = "mailhog:1025" //mailhog
	appConfig.CacheConn = conn

	err = rpc.Register(&MailRPC{})
	if err != nil {
		log.Fatalln(err)
	}
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", PORT))
	if err != nil {
		log.Fatalln(err)
	}
	if err = http.Serve(l, nil); err != nil {
		log.Println(err)
	}
}
