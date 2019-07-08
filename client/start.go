package client

import (
	"encoding/json"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/peer"
	_ "github.com/davyxu/cellnet/peer/tcp"
	"github.com/davyxu/cellnet/proc"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"github.com/davyxu/golog"
	"im_client/config"
)

var log = golog.New("client.start")

var UserId int64 = 1

var GenericPeer cellnet.GenericPeer

func StartClient() {
	peerType := config.Config.Get("cellnet.peerType").(string)
	name := config.Config.Get("cellnet.name").(string)
	addr := config.Config.Get("cellnet.addr").(string)
	procName := config.Config.Get("cellnet.procName").(string)
	queue := cellnet.NewEventQueue()
	GenericPeer = peer.NewGenericPeer(peerType, name, addr, queue)
	proc.BindProcessorHandler(GenericPeer, procName, func(ev cellnet.Event) {
		switch msg := ev.Message().(type) {
		case *cellnet.SessionConnected:
			log.Infoln("client connected")
		case *cellnet.SessionClosed:
			log.Infoln("client error")
		case *Message:
			if msg.Cmd == CmdTypeLoginResp {
				// 展示未读消息
				chatUnReadMsgListString := msg.Extras
				var chatUnReadMsgList []map[string]string
				if err := json.Unmarshal([]byte(chatUnReadMsgListString), &chatUnReadMsgList); err == nil {
					log.Errorln(err)
				}
				for chatUnReadMsg := range chatUnReadMsgList {
					log.Infoln("%s", chatUnReadMsg)
				}
			} else {
				log.Infoln("%s", msg)
			}
		}
	})
	GenericPeer.Start()
	queue.StartLoop()
}
