package savemacro

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type Act struct {
	Act_name        string
	Mouse_xPos      int
	Mouse_yPos      int
	Keyboard_action string
}

func Save_mouse(act_name string) Act{
	var x, y = robotgo.GetMousePos()
	action := Act{
		Act_name: act_name,
		Mouse_xPos: x,
		Mouse_yPos: y,
		Keyboard_action: "",
	}
	return action
}

func Save_keyboard(act_name, keyboard_action string) Act{
	action := Act{
		Act_name: act_name,
		Mouse_xPos: 0,
		Mouse_yPos: 0,
		Keyboard_action: keyboard_action,
	}
	return action
}

func (a Act) GetString() string{
	return a.Act_name + " " + fmt.Sprint(a.Mouse_xPos) + " " + fmt.Sprint(a.Mouse_yPos)
}
