package client


import (
	"github.com/davyxu/cellnet"
	"github.com/davyxu/cellnet/codec"
	_ "github.com/davyxu/cellnet/codec/protoplus"
	"github.com/davyxu/cellnet/util"
	"github.com/davyxu/protoplus/proto"
	"reflect"
	"unsafe"
)

var (
	_ *proto.Buffer
	_ unsafe.Pointer
)

type CommandType int32

//noinspection GoUnusedConst
const (
	CmdTypeUnknow               CommandType = 0
	CmdTypeHandshakeReq         CommandType = 1
	CmdTypeHandshakeResp        CommandType = 2
	CmdTypeAuthReq              CommandType = 3
	CmdTypeAuthResp             CommandType = 4
	CmdTypeLoginReq             CommandType = 5
	CmdTypeLoginResp            CommandType = 6
	CmdTypeJoinGroupReq         CommandType = 7
	CmdTypeJoinGroupResp        CommandType = 8
	CmdTypeJoinGroupNotifyResp  CommandType = 9
	CmdTypeExitGroupNotifyResp  CommandType = 10
	CmdTypeChatReq              CommandType = 11
	CmdTypeChatResp             CommandType = 12
	CmdTypeHeartbeatReq         CommandType = 13
	CmdTypeCloseReq             CommandType = 14
	CmdTypeCancelMsgReq         CommandType = 15
	CmdTypeCancelMsgResp        CommandType = 16
	CmdTypeGetUserReq           CommandType = 17
	CmdTypeGetUserResp          CommandType = 18
	CmdTypeGetMessageReq        CommandType = 19
	CmdTypeGetMessageResp       CommandType = 20
	CmdTypeGetUnReadMessageReq  CommandType = 21
	CmdTypeGetUnReadMessageResp CommandType = 22
)

var (
	_ = map[string]int32{
		"UNKNOW":                   0,
		"HANDSHAKE_REQ":            1,
		"HANDSHAKE_RESP":           2,
		"AUTH_REQ":                 3,
		"AUTH_RESP":                4,
		"LOGIN_REQ":                5,
		"LOGIN_RESP":               6,
		"JOIN_GROUP_REQ":           7,
		"JOIN_GROUP_RESP":          8,
		"JOIN_GROUP_NOTIFY_RESP":   9,
		"EXIT_GROUP_NOTIFY_RESP":   10,
		"CHAT_REQ":                 11,
		"CHAT_RESP":                12,
		"HEARTBEAT_REQ":            13,
		"CLOSE_REQ":                14,
		"CANCEL_MSG_REQ":           15,
		"CANCEL_MSG_RESP":          16,
		"GET_USER_REQ":             17,
		"GET_USER_RESP":            18,
		"GET_MESSAGE_REQ":          19,
		"GET_MESSAGE_RESP":         20,
		"GET_UN_READ_MESSAGE_REQ":  21,
		"GET_UN_READ_MESSAGE_RESP": 22,
	}

	CommandTypeMapperNameByValue = map[int32]string{
		0:  "UNKNOW",
		1:  "HANDSHAKE_REQ",
		2:  "HANDSHAKE_RESP",
		3:  "AUTH_REQ",
		4:  "AUTH_RESP",
		5:  "LOGIN_REQ",
		6:  "LOGIN_RESP",
		7:  "JOIN_GROUP_REQ",
		8:  "JOIN_GROUP_RESP",
		9:  "JOIN_GROUP_NOTIFY_RESP",
		10: "EXIT_GROUP_NOTIFY_RESP",
		11: "CHAT_REQ",
		12: "CHAT_RESP",
		13: "HEARTBEAT_REQ",
		14: "CLOSE_REQ",
		15: "CANCEL_MSG_REQ",
		16: "CANCEL_MSG_RESP",
		17: "GET_USER_REQ",
		18: "GET_USER_RESP",
		19: "GET_MESSAGE_REQ",
		20: "GET_MESSAGE_RESP",
		21: "GET_UN_READ_MESSAGE_REQ",
		22: "GET_UN_READ_MESSAGE_RESP",
	}
)

func (cmdType CommandType) String() string {
	return CommandTypeMapperNameByValue[int32(cmdType)]
}

type MsgType int32

//noinspection GoUnusedConst
const (
	MsgTypeText  MsgType = 0
	MsgTypeImg   MsgType = 2
	MsgTypeVoice MsgType = 3
	MsgTypeVideo MsgType = 4
	MsgTypeMusic MsgType = 5
	MsgTypeNews  MsgType = 6
)

var (
	_ = map[string]int32{
		"TEXT":  0,
		"IMG":   2,
		"VOICE": 3,
		"VIDEO": 4,
		"MUSIC": 5,
		"NEWS":  6,
	}

	MsgTypeMapperNameByValue = map[int32]string{
		0: "TEXT",
		2: "IMG",
		3: "VOICE",
		4: "VIDEO",
		5: "MUSIC",
		6: "NEWS",
	}
)

func (msgType MsgType) String() string {
	return MsgTypeMapperNameByValue[int32(msgType)]
}

func (msgType MsgType) Int() int32 {
	return int32(msgType)
}

type ChatType int32

//noinspection ALL
const (
	ChatTypeUnknow  ChatType = 0
	ChatTypePublic  ChatType = 1
	ChatTypePrivate ChatType = 2
)

var (
	_ = map[string]int32{
		"UNKNOW":  0,
		"PUBLIC":  1,
		"PRIVATE": 2,
	}

	ChatTypeMapperNameByValue = map[int32]string{
		0: "UNKNOW",
		1: "PUBLIC",
		2: "PRIVATE",
	}
)

func (chatType ChatType) String() string {
	return ChatTypeMapperNameByValue[int32(chatType)]
}

func (chatType ChatType) Int() int32 {
	return int32(chatType)
}

type Message struct {
	From       int64
	To         int64
	Cmd        CommandType
	CreateTime int64
	MsgType    MsgType
	ChatType   ChatType
	GroupId    string
	Content    string
	Extras     string
}

func (msg *Message) String() string { return proto.CompactTextString(msg) }

func (msg *Message) Size() (ret int) {

	ret += proto.SizeInt64(0, msg.From)

	ret += proto.SizeInt64(1, msg.To)

	ret += proto.SizeInt32(2, int32(msg.Cmd))

	ret += proto.SizeInt64(3, msg.CreateTime)

	ret += proto.SizeInt32(4, int32(msg.MsgType))

	ret += proto.SizeInt32(5, int32(msg.ChatType))

	ret += proto.SizeString(6, msg.GroupId)

	ret += proto.SizeString(7, msg.Content)

	ret += proto.SizeString(8, msg.Extras)

	return
}

func (msg *Message) Marshal(buffer *proto.Buffer) error {

	_ = proto.MarshalInt64(buffer, 0, msg.From)

	_ = proto.MarshalInt64(buffer, 1, msg.To)

	_ = proto.MarshalInt32(buffer, 2, int32(msg.Cmd))

	_ = proto.MarshalInt64(buffer, 3, msg.CreateTime)

	_ = proto.MarshalInt32(buffer, 4, int32(msg.MsgType))

	_ = proto.MarshalInt32(buffer, 5, int32(msg.ChatType))

	_ = proto.MarshalString(buffer, 6, msg.GroupId)

	_ = proto.MarshalString(buffer, 7, msg.Content)

	_ = proto.MarshalString(buffer, 8, msg.Extras)

	return nil
}

func (msg *Message) Unmarshal(buffer *proto.Buffer, fieldIndex uint64, wt proto.WireType) error {
	switch fieldIndex {
	case 0:
		return proto.UnmarshalInt64(buffer, wt, &msg.From)
	case 1:
		return proto.UnmarshalInt64(buffer, wt, &msg.To)
	case 2:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&msg.Cmd))
	case 3:
		return proto.UnmarshalInt64(buffer, wt, &msg.CreateTime)
	case 4:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&msg.MsgType))
	case 5:
		return proto.UnmarshalInt32(buffer, wt, (*int32)(&msg.ChatType))
	case 6:
		return proto.UnmarshalString(buffer, wt, &msg.GroupId)
	case 7:
		return proto.UnmarshalString(buffer, wt, &msg.Content)
	case 8:
		return proto.UnmarshalString(buffer, wt, &msg.Extras)

	}

	return proto.ErrUnknownField
}

func init() {

	cellnet.RegisterMessageMeta(&cellnet.MessageMeta{
		Codec: codec.MustGetCodec("protoplus"),
		Type:  reflect.TypeOf((*Message)(nil)).Elem(),
		ID:    int(util.StringHash("proto.Message")),
	})
}
