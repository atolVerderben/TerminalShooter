package termvel

import tl "github.com/JoelOtter/termloop"

//EventInfo used primarly for debugging right now
type EventInfo struct {
	text             *tl.Text
	right            bool
	center           bool
	OffsetX, OffsetY int
}

//NewEventInfo returns a new EventInfo...
func NewEventInfo(x, y int) *EventInfo {
	info := &EventInfo{}
	info.text = tl.NewText(x, y, "Click somewhere", tl.ColorWhite, tl.ColorBlack)
	return info
}

//Draw the text on the screen
func (info *EventInfo) Draw(screen *tl.Screen) {
	//TODO: fix this nonsense
	if info.right {
		w, _ := info.text.Size()
		screenWidth, _ := screen.Size()
		info.text.SetPosition(screenWidth-w, 0)
	}
	if info.center { //TODO: Fix this to check for level offset!
		w, h := info.text.Size()
		screenWidth, screenHeight := screen.Size()
		info.text.SetPosition(screenWidth/2-w/2-info.OffsetX, screenHeight/2-h/2-info.OffsetY)
	}
	info.text.Draw(screen)
}

//Tick to handle user input
func (info *EventInfo) Tick(ev tl.Event) {
	if ev.Type != tl.EventMouse {
		return
	}
	/*var name string
	switch ev.Key {
	case tl.MouseLeft:
		name = "Mouse Left"
	case tl.MouseMiddle:
		name = "Mouse Middle"
	case tl.MouseRight:
		name = "Mouse Right"
	case tl.MouseWheelUp:
		name = "Mouse Wheel Up"
	case tl.MouseWheelDown:
		name = "Mouse Wheel Down"
	case tl.MouseRelease:
		name = "Mouse Release"
	default:
		name = fmt.Sprintf("Unknown Key (%#x)", ev.Key)
	}
	info.text.SetText(fmt.Sprintf("%s @ [%d, %d]", name, ev.MouseX, ev.MouseY))*/
}
