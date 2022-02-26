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

func Save_act(act_name string) Act {
	action := Act{
		Act_name: act_name,
		Mouse_xPos: 0,
		Mouse_yPos: 0,
		Keyboard_action: "",
	}
	return action
}

func (a Act) GetString() string{
	return a.Act_name + ":" + fmt.Sprint(a.Mouse_xPos) + ":" + fmt.Sprint(a.Mouse_yPos)
}

func ExcuteMacro (macroInput []Act) {
	robotgo.MouseSleep = 5
	for _, macroID := range macroInput {
		switch macroID.Act_name {
		case "Move":
			fmt.Println("Move")
			robotgo.MoveMouseSmooth(macroID.Mouse_xPos, macroID.Mouse_yPos)
		case "KeyUp":
			fmt.Println("KeyUp")
			robotgo.KeyUp(macroID.Keyboard_action)
		case "KeyDown":
			fmt.Println("KeyDown")
			robotgo.KeyDown(macroID.Keyboard_action)
		case "Mouse:Down":
			fmt.Println("Mouse:Down")
			robotgo.MouseToggle("down", "left")
		case "Mouse:Up":
			fmt.Println("Mouse:Up")
			robotgo.MouseToggle("up", "left")
		default :
			fmt.Println("Should Not Pass")
		}
	}
}