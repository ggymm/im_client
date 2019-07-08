package client

import (
	"fmt"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"strconv"
)

const (
	title       = "聊天客户端"
	allWidth    = 800
	allHeight   = 600
	friendWidth = 200
	chartWidth  = 600
	chatHeight  = 450
	msgHeight   = 117
)

type ChatMainWindow struct {
	*walk.MainWindow
	friendList      *walk.ListBox
	friendListModel *FriendListModel
	chatContentList *walk.TextEdit
	msgContent      *walk.TextEdit
}

type FriendListModel struct {
	walk.ListModelBase
	items []FriendItem
}

func (m *FriendListModel) ItemCount() int {
	return len(m.items)
}

func (m *FriendListModel) Value(index int) interface{} {
	return m.items[index].Name
}

type FriendItem struct {
	Id    int64
	Name  string
	Value string
}

func StartView() error {
	mainWindow := &ChatMainWindow{friendListModel: GetFriendList()}
	mainWindow.SetMaximizeBox(false)
	mainWindow.SetFixedSize(true)
	if err := (MainWindow{
		AssignTo: &mainWindow.MainWindow,
		Title:    title,
		Size:     Size{Width: allWidth, Height: allHeight},
		MinSize:  Size{Width: allWidth, Height: allHeight},
		Layout:   HBox{MarginsZero: true, SpacingZero: true},
		Children: []Widget{
			ListBox{
				AssignTo:              &mainWindow.friendList,
				Model:                 mainWindow.friendListModel,
				OnCurrentIndexChanged: mainWindow.FriendCurrentIndexChanged,
				OnItemActivated:       mainWindow.FriendItemActivated,
				MaxSize:               Size{Width: friendWidth},
			},
			Composite{
				MinSize: Size{Width: chartWidth},
				MaxSize: Size{Width: chartWidth},
				Layout:  VBox{MarginsZero: true, SpacingZero: true},
				Children: []Widget{
					TextEdit{
						AssignTo: &mainWindow.chatContentList,
						ReadOnly: true,
						MaxSize:  Size{Height: chatHeight},
					},
					Composite{
						Border:  false,
						Layout:  Grid{Columns: 10, MarginsZero: true, SpacingZero: true},
						MaxSize: Size{Height: msgHeight},
						Children: []Widget{
							TextEdit{
								ColumnSpan: 9,
								AssignTo:   &mainWindow.msgContent,
								VScroll:    true,
								MinSize:    Size{Height: msgHeight},
								MaxSize:    Size{Height: msgHeight},
							},
							PushButton{
								Text:       "提交",
								ColumnSpan: 1,
								MinSize:    Size{Height: msgHeight},
								MaxSize:    Size{Height: msgHeight},
								OnClicked: func() {
									// 发送消息到服务端
									msgContent := mainWindow.msgContent.Text()
									fmt.Println(mainWindow.friendList.CurrentIndex())
									friendId := mainWindow.friendListModel.items[mainWindow.friendList.CurrentIndex()].Id
									SendMsg(msgContent, friendId)
								},
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		return err
	}
	mainWindow.Run()
	return nil
}

func GetFriendList() *FriendListModel {
	friendListModel := &FriendListModel{items: make([]FriendItem, 10)}
	for i := 0; i < 10; i++ {
		friendListModel.items[i] = FriendItem{
			Id:    int64(i),
			Name:  "name" + strconv.Itoa(i),
			Value: strconv.Itoa(i),
		}
	}

	return friendListModel
}

func (chatMainWindow *ChatMainWindow) FriendCurrentIndexChanged() {

}

func (chatMainWindow *ChatMainWindow) FriendItemActivated() {
}
