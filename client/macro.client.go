package client

import (
	"fmt"
	"strings"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var canRun bool

func GetCurMousePosition() (int, int){
	return robotgo.GetMousePos()
}

func RunMacro() {
	fmt.Println("Run")
	canRun = true
	
	go robotgotest()

	breakAll := false
	for {
		for index, macros := range MacroInput {
			if !canRun {
				fmt.Println(index)
				breakAll = true
				break
			} else {
				if strings.Contains(macros, ":"){
					splitMacro := strings.Split(macros, ":")
					if splitMacro[0] == "   키보드 누름"{

					} 
				} else {
					switch macros{
					case LMouseDown:
						fmt.Println(LMouseDown)
					case RMouseDown:
						fmt.Println(RMouseDown)
					case LMouseUp:
						fmt.Println(LMouseUp)
					case RMouseUp:
						fmt.Println(RMouseUp)
					case LMouseClick:
						fmt.Println(LMouseClick)
					case RMouseClick:
						fmt.Println(RMouseClick)
					}
				}
			}
		}
		if breakAll {
			break
		}
	}
}

func robotgotest() {
	evc := hook.Start()
	defer hook.End()
	for ev := range evc {
		if ev.Keycode == 31 {
			canRun = false
		}
	}
}