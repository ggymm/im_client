package client

import (
	"encoding/json"
	"github.com/davyxu/cellnet"
	"im_client/constant"
	"time"
)

/**
 * 登陆
 */
func Login() {
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

/**
 * 发送消息
 */
func SendMsg(msgContent string, friendId int64) {
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
		Content:    msgContent,
		Extras:     constant.Empty,
	})
}

func ReadOfflineMsg(msgContent string, friendId int64) {
	contentMap := make(map[string]interface{}, 0)
	contentMap["chartType"] = ChatTypePublic.String()
	contentMap["friendId"] = friendId
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
