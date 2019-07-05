package client

import (
	"encoding/json"
	"github.com/davyxu/cellnet"
	"github.com/davyxu/golog"
	"im_client/constant"
	"im_client/utils"
	"strings"
	"time"
)

var handleLog = golog.New("client.handle")

func Login(str string) {
	// 注册消息：login 1
	loginInfos := strings.Split(str, " ")
	if err := utils.StrToInt(loginInfos[1], &UserId); err != nil {
		handleLog.Infof(err.Error())
	}
	// 发送注册方法
	GenericPeer.(interface {
		Session() cellnet.Session
	}).Session().Send(&Message{
		From:       UserId,
		To:         constant.ServerId,
		Cmd:        CmdTypeLoginReq,
		CreateTime: time.Now().Unix() / 1e6,
		MsgType:    MsgTypeText,
		ChatType:   ChatTypePrivate,
		GroupId:    constant.Empty,
		Content:    constant.Empty,
		Extras:     constant.Empty,
	})
}

func ReadOfflineMsg(str string) {
	// 读取消息列表，read 1
	// 是好友或者群聊
	getMessageInfos := strings.Split(str, " ")
	contentMap := make(map[string]string, 0)
	contentMap["chartType"] = ChatTypePublic.String()
	contentMap["friendId"] = getMessageInfos[1]
	contentString, _ := json.Marshal(contentMap)
	// 发送获取消息方法
	GenericPeer.(interface {
		Session() cellnet.Session
	}).Session().Send(&Message{
		From:       UserId,
		To:         constant.ServerId,
		Cmd:        CmdTypeGetUnReadMessageReq,
		CreateTime: time.Now().Unix() / 1e6,
		MsgType:    MsgTypeText,
		ChatType:   ChatTypePrivate,
		GroupId:    constant.Empty,
		Content:    string(contentString),
		Extras:     constant.Empty,
	})
}

func SendMsg(str string) {
	// 发送消息给其他客户端：内容 1
	msgInfo := strings.Split(str, " ")
	var friendId int64
	if err := utils.StrToInt(msgInfo[1], &friendId); err != nil {
		log.Infof(err.Error())
	}
	GenericPeer.(interface {
		Session() cellnet.Session
	}).Session().Send(&Message{
		From:       UserId,
		To:         friendId,
		Cmd:        CmdTypeChatReq,
		CreateTime: time.Now().Unix() / 1e6,
		MsgType:    MsgTypeText,
		ChatType:   ChatTypePrivate,
		GroupId:    constant.Empty,
		Content:    msgInfo[0],
		Extras:     constant.Empty,
	})
}
