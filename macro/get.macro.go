package macro

import "github.com/go-vgo/robotgo"

func GetCurMousePosition() (int, int){
	return robotgo.GetMousePos()
}