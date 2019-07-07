package main

import (
	"flag"
	"fmt"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"im_client/client"
	"im_client/ui"
)

func main() {

	clientId := flag.Int("client_id", 0, "当前客户端ID")
	flag.Parse()
	fmt.Println("当前客户端ID:", *clientId)
	ui.StartView()
	client.StartClient()
}
