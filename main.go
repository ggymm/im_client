package main

import (
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"im_client/client"
	"im_client/ui"
)

func main() {
	ui.StartView()
	client.StartClient()
}
