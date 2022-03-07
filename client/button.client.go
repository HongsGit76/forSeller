package client

import (
	"fmt"
	"macroProj/macro/macroActs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func addButton(btBox *fyne.Container) {
	// current state message
	keyboardModeState := widget.NewLabel(DefaultModeString)

	// {button} keyboard macro
	keboardMacro := widget.NewButton("키보드 매크로", func ()  {
		if KeyboardMode {
			MouseMode = false
			KeyboardMode = false
			keyboardModeState.SetText(DefaultModeString)
		}
		KeyboardMode = true
		MouseMode = false
		keyboardModeState.SetText(KeyboardModeString)
	})
	
	// {button} mouse position macro
	mouseMacro := widget.NewButton("마우스 매크로 (F7)", func ()  {
		if MouseMode {
			MouseMode = false
			KeyboardMode = false
			keyboardModeState.SetText(DefaultModeString)
	}
		MouseMode = true
		KeyboardMode = false
		keyboardModeState.SetText(MouseModeString)
	})

	mouseDownBox := container.NewHBox()
	mouseUpBox := container.NewHBox()
	mouseClickBox := container.NewHBox()

	// {button} mouse act macro
	mouseDown_L := widget.NewButton(LMouseDown, func ()  {
		MacroData.Append(LMouseDown)
		macroActs.MakeAndAppendMacro(macroActs.DOWN_MOUSE_L, 0,0,"")
	})
	mouseDown_R := widget.NewButton(RMouseDown, func ()  {
		MacroData.Append(RMouseDown)
		macroActs.MakeAndAppendMacro(macroActs.DOWN_MOUSE_R, 0,0,"")
	})
	mouseDownBox.Add(mouseDown_L)
	mouseDownBox.Add(mouseDown_R)

	mouseUp_L := widget.NewButton(LMouseUp, func ()  {
		MacroData.Append(LMouseUp)
		macroActs.MakeAndAppendMacro(macroActs.UP_MOUSE_L, 0,0,"")
	})
	mouseUp_R := widget.NewButton(RMouseUp, func ()  {
		MacroData.Append(RMouseUp)
		macroActs.MakeAndAppendMacro(macroActs.UP_MOUSE_R, 0,0,"")
	})
	mouseUpBox.Add(mouseUp_L)
	mouseUpBox.Add(mouseUp_R)

	mouseClick_L := widget.NewButton(LMouseClick, func ()  {
		MacroData.Append(LMouseClick)
		macroActs.MakeAndAppendMacro(macroActs.CLICK_MOUSE_L, 0,0,"")
	})
	mouseClick_R := widget.NewButton(RMouseClick, func ()  {
		MacroData.Append(RMouseClick)
		macroActs.MakeAndAppendMacro(macroActs.CLICK_MOUSE_R, 0,0,"")
	})
	mouseClickBox.Add(mouseClick_L)
	mouseClickBox.Add(mouseClick_R)

	excuteMacroButton := widget.NewButton("매크로 수행(F5)", func ()  {
		fmt.Println(MacroInput)
		fmt.Println(MacroData)
		fmt.Println(macroActs.MacroActs)
	})

	btBox.Add(keyboardModeState)
	btBox.Add(keboardMacro)
	btBox.Add(mouseMacro)
	btBox.Add(mouseDownBox)
	btBox.Add(mouseUpBox)
	btBox.Add(mouseClickBox)
	btBox.Add(excuteMacroButton)
}