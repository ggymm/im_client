package ui

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"strings"
)

const (
	allWidth         = 800
	allHeight        = 600
	contentHeight    = 605
	friendWidth      = 200
	chartWidth       = 600
	chatHeight       = 450
	msgHeight        = 116
	msgContentWidth  = 450
	msgContentHeight = 150
)

type ChatMainWindow struct {
	*walk.MainWindow
	friendList      *walk.ListBox
	friendListModel *FriendListModel
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

func StartView() {
	title := "聊天客户端"
	mainWindow := &ChatMainWindow{friendListModel: GetFriendList()}
	// 不允许缩放窗口，不允许最大化窗口
	mainWindow.SetMaximizeBox(false)
	mainWindow.SetFixedSize(true)
	var inTE, outTE *walk.TextEdit
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
						AssignTo: &outTE,
						ReadOnly: true,
						MaxSize:  Size{Height: chatHeight},
					},
					Composite{
						Border:  false,
						Layout:  Grid{Columns: 5, MarginsZero: true, SpacingZero: true},
						MaxSize: Size{Height: msgHeight},
						Children: []Widget{
							TextEdit{
								ColumnSpan: 4,
								AssignTo:   &inTE,
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
									outTE.SetText(strings.ToUpper(inTE.Text()))
								},
							},
						},
					},
				},
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}
	mainWindow.Run()
}

func GetFriendList() *FriendListModel {
	friendListModel := &FriendListModel{items: make([]FriendItem, 10)}
	for i := 0; i < 10; i++ {
		friendListModel.items[i] = FriendItem{
			Id:    int64(i),
			Name:  "name",
			Value: string(i),
		}
	}

	return friendListModel
}

func (mw *ChatMainWindow) FriendCurrentIndexChanged() {

}

func (mw *ChatMainWindow) FriendItemActivated() {

}

func TestView() {
	var slv, slh *walk.Slider
	var maxEdit, minEdit, valueEdit *walk.NumberEdit

	data := struct{ Min, Max, Value int }{0, 100, 30}

	MainWindow{
		Title:   "Walk Slider Example",
		MinSize: Size{320, 240},
		Layout:  HBox{},
		Children: []Widget{
			Slider{
				AssignTo:    &slv,
				MinValue:    data.Min,
				MaxValue:    data.Max,
				Value:       data.Value,
				Orientation: Vertical,
				OnValueChanged: func() {
					data.Value = slv.Value()
					valueEdit.SetValue(float64(data.Value))

				},
			},
			Composite{
				Layout:        Grid{Columns: 3},
				StretchFactor: 4,
				Children: []Widget{
					Label{Text: "Min value"},
					Label{Text: "Value"},
					Label{Text: "Max value"},
					NumberEdit{
						AssignTo: &minEdit,
						Value:    float64(data.Min),
						OnValueChanged: func() {
							data.Min = int(minEdit.Value())
							slh.SetRange(data.Min, data.Max)
							slv.SetRange(data.Min, data.Max)
						},
					},
					NumberEdit{
						AssignTo: &valueEdit,
						Value:    float64(data.Value),
						OnValueChanged: func() {
							data.Value = int(valueEdit.Value())
							slh.SetValue(data.Value)
							slv.SetValue(data.Value)
						},
					},
					NumberEdit{
						AssignTo: &maxEdit,
						Value:    float64(data.Max),
						OnValueChanged: func() {
							data.Max = int(maxEdit.Value())
							slh.SetRange(data.Min, data.Max)
							slv.SetRange(data.Min, data.Max)
						},
					},
					Slider{
						ColumnSpan: 3,
						AssignTo:   &slh,
						MinValue:   data.Min,
						MaxValue:   data.Max,
						Value:      data.Value,
						OnValueChanged: func() {
							data.Value = slh.Value()
							valueEdit.SetValue(float64(data.Value))
						},
					},
					VSpacer{},
					PushButton{
						ColumnSpan: 3,
						Text:       "Print state",
						OnClicked: func() {
							log.Printf("H: < %d | %d | %d >\n", slh.MinValue(), slh.Value(), slh.MaxValue())
							log.Printf("V: < %d | %d | %d >\n", slv.MinValue(), slv.Value(), slv.MaxValue())
						},
					},
				},
			},
		},
	}.Run()
}
