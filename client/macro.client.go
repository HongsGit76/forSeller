package client

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

var canRun bool
var stoppedIndex int

func GetCurMousePosition() (int, int){
	return robotgo.GetMousePos()
}

func RunMacro() {
	fmt.Println("Run")
	canRun = true
	robotgo.MouseSleep = 100  // 100 millisecond
	robotgo.KeySleep = 10
	go robotgotest()

	breakAll := false
	for {
		for index, macros := range MacroInput {
			if !canRun {
				stoppedIndex = index
				breakAll = true
				break
			} else {
				if strings.Contains(macros, ":"){
					splitMacro := strings.Split(macros, ":")
					if strings.Compare(splitMacro[0], "   키보드 누름") == 0 {
						if strings.Contains(splitMacro[1], "Control") {
							nextSplitMacro := strings.Split(MacroInput[index + 1], ":")
							robotgo.KeyToggle(splitMacro[1], nextSplitMacro[1])
						} else if beforeSplit := strings.Split(MacroInput[index - 1], ":"); strings.Contains(beforeSplit[1], "Control") {
							continue
						} 
						robotgo.KeyToggle(splitMacro[1])
					} else if strings.Compare(splitMacro[0], "   키보드 뗌") == 0  {
						if strings.Contains(splitMacro[1], "Control") {
							nextSplitMacro := strings.Split(MacroInput[index - 1], ":")
							robotgo.KeyToggle(splitMacro[1], nextSplitMacro[1], "up")
						} else if beforeSplit := strings.Split(MacroInput[index + 1], ":"); strings.Contains(beforeSplit[1], "Control") {
							continue
						} 
						robotgo.KeyToggle(splitMacro[1], "up")
					} else if strings.Compare(splitMacro[0], "   마우스 이동") == 0 {
						location := strings.Split(splitMacro[1], ",")
						x, errX := strconv.Atoi(location[0])
						if errX != nil {
						}
						y, errY := strconv.Atoi(location[1])
						if errY != nil {
						}
						robotgo.Move(x,y)
					} else if strings.Compare(splitMacro[0], "   시간 추가") == 0 {
						times, err := strconv.Atoi(splitMacro[1])
						if err != nil {
						}
						time.Sleep(time.Duration(times) * time.Second)
					}
				} else {
					switch macros{
					case LMouseDown:
						robotgo.Toggle("left", "down")
					case RMouseDown:
						robotgo.Toggle("right", "down")
					case LMouseUp:
						robotgo.Toggle("left", "up")
					case RMouseUp:
						robotgo.Toggle("right", "up")
					case LMouseClick:
						robotgo.Click()
					case RMouseClick:
						robotgo.Click("right")
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
			fmt.Println("Stop!")
			robotgo.Toggle("left", "up")
			robotgo.Toggle("left", "up")
			robotgo.KeyToggle("up")
			canRun = false
		}
	}
}