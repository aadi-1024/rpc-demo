package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/smtp"
)

type Config struct {
	SMTPAuth  smtp.Auth
	SMTPAddr  string //address for the smtp server, localhost:1025 in case of mailhog
	CacheConn *rpc.Client
}

func (c *Config) SendMail(mail string, content string) error {
	if mail == "" {
		err := smtp.SendMail(c.SMTPAddr, c.SMTPAuth, "admin@example.com", []string{mail}, []byte(content))
		if err != nil {
			fmt.Println(err)
		}
		return err
	}
	var emails []string
	if err := c.CacheConn.Call("CacheRPC.GetAll", 0, &emails); err != nil {
		log.Println(err)
	}

	err := smtp.SendMail(c.SMTPAddr, c.SMTPAuth, "admin@example.com", emails, []byte(content))
	if err != nil {
		log.Println(err)
	}
	return err
}
