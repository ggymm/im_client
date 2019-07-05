package main

import (
	"bufio"
	"encoding/json"
	"github.com/davyxu/cellnet"
	_ "github.com/davyxu/cellnet/peer/tcp"
	_ "github.com/davyxu/cellnet/proc/tcp"
	"im_client/client"
	"im_client/constant"
	"im_client/ui"
	"im_client/utils"
	"log"
	"os"
	"strings"
	"time"
)

var userId int64 = 0

func ReadConsole(callback func(string)) {
	for {
		// 从标准输入读取字符串，以\n为分割
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}
		// 去掉读入内容的空白符
		text = strings.TrimSpace(text)
		callback(text)
	}
}

//noinspection GoUnusedExportedFunction
func StartClient() {
	// 阻塞的从命令行获取聊天输入
	ReadConsole(func(str string) {
		// 判断是不是注册消息
		if strings.Contains(str, "login") {
			// 注册消息：login 1
			loginInfos := strings.Split(str, " ")
			if err := utils.StrToInt(loginInfos[1], &userId); err != nil {
				log.Println(err.Error())
			}
			// 发送注册方法
			client.GenericPeer.(interface {
				Session() cellnet.Session
			}).Session().Send(&client.Message{
				From:       userId,
				To:         constant.ServerId,
				Cmd:        client.CmdTypeLoginReq,
				CreateTime: time.Now().Unix() / 1e6,
				MsgType:    client.MsgTypeText,
				ChatType:   client.ChatTypePrivate,
				GroupId:    constant.Empty,
				Content:    constant.Empty,
				Extras:     constant.Empty,
			})
		} else if strings.Contains(str, "read") {
			if userId != 0 {
				// 读取消息列表，read 1
				// 是好友或者群聊
				getMessageInfos := strings.Split(str, " ")
				contentMap := make(map[string]string, 0)
				contentMap["chartType"] = client.ChatTypePublic.String()
				contentMap["friendId"] = getMessageInfos[1]
				contentString, _ := json.Marshal(contentMap)
				// 发送获取消息方法
				client.GenericPeer.(interface {
					Session() cellnet.Session
				}).Session().Send(&client.Message{
					From:       userId,
					To:         constant.ServerId,
					Cmd:        client.CmdTypeGetUnReadMessageReq,
					CreateTime: time.Now().Unix() / 1e6,
					MsgType:    client.MsgTypeText,
					ChatType:   client.ChatTypePrivate,
					GroupId:    constant.Empty,
					Content:    string(contentString),
					Extras:     constant.Empty,
				})
			} else {
				log.Println("未注册！")
			}
		} else if strings.Contains(str, " ") {
			if userId != 0 {
				// 发送消息给其他客户端：内容 1
				msgInfo := strings.Split(str, " ")
				var friendId int64
				if err := utils.StrToInt(msgInfo[1], &friendId); err != nil {
					log.Println(err.Error())
				}
				client.GenericPeer.(interface {
					Session() cellnet.Session
				}).Session().Send(&client.Message{
					From:       userId,
					To:         friendId,
					Cmd:        client.CmdTypeChatReq,
					CreateTime: time.Now().Unix() / 1e6,
					MsgType:    client.MsgTypeText,
					ChatType:   client.ChatTypePrivate,
					GroupId:    constant.Empty,
					Content:    msgInfo[0],
					Extras:     constant.Empty,
				})
			} else {
				log.Println("未注册！")
			}
		} else {
			log.Println("消息格式不正确！")
		}
	})
}

func main() {
	ui.StartView()
	client.StartClient()
}
