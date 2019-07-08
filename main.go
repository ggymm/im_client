package main

import (
	"flag"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"im_client/client"
)

func main() {
	clientId := flag.Int64("client_id", 1, "当前客户端ID")
	flag.Parse()
	client.UserId = *clientId
	client.StartClient()
	_ = client.StartView()
}
