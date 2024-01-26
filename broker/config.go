package main

import "net/rpc"

type AppConfig struct {
	CacheConn, MailConn *rpc.Client
}
