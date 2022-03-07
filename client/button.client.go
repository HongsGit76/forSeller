package client

import (
	"fmt"
	"macroProj/macro/macro"

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
		macro.MakeAndAppendMacro(macro.DOWN_MOUSE_L, 0,0,"")
	})
	mouseDown_R := widget.NewButton(RMouseDown, func ()  {
		MacroData.Append(RMouseDown)
		macro.MakeAndAppendMacro(macro.DOWN_MOUSE_R, 0,0,"")
	})
	mouseDownBox.Add(mouseDown_L)
	mouseDownBox.Add(mouseDown_R)

	mouseUp_L := widget.NewButton(LMouseUp, func ()  {
		MacroData.Append(LMouseUp)
		macro.MakeAndAppendMacro(macro.UP_MOUSE_L, 0,0,"")
	})
	mouseUp_R := widget.NewButton(RMouseUp, func ()  {
		MacroData.Append(RMouseUp)
		macro.MakeAndAppendMacro(macro.UP_MOUSE_R, 0,0,"")
	})
	mouseUpBox.Add(mouseUp_L)
	mouseUpBox.Add(mouseUp_R)

	mouseClick_L := widget.NewButton(LMouseClick, func ()  {
		MacroData.Append(LMouseClick)
		macro.MakeAndAppendMacro(macro.CLICK_MOUSE_L, 0,0,"")
	})
	mouseClick_R := widget.NewButton(RMouseClick, func ()  {
		MacroData.Append(RMouseClick)
		macro.MakeAndAppendMacro(macro.CLICK_MOUSE_R, 0,0,"")
	})
	mouseClickBox.Add(mouseClick_L)
	mouseClickBox.Add(mouseClick_R)

	excuteMacroButton := widget.NewButton("매크로 수행(F5)", func ()  {
		fmt.Println(MacroInput)  // string
		fmt.Println(MacroData)  // list view
		fmt.Println(macro.MacroActs) // macros
	})

	keyboardModeState.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))
	keboardMacro.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))
	mouseMacro.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))
	mouseDownBox.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))
	mouseUpBox.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))
	mouseClickBox.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))
	excuteMacroButton.Resize(fyne.NewSize(ButtonSize.width, ButtonSize.height))

	btBox.Add(keyboardModeState)
	btBox.Add(keboardMacro)
	btBox.Add(mouseMacro)
	btBox.Add(mouseDownBox)
	btBox.Add(mouseUpBox)
	btBox.Add(mouseClickBox)
	btBox.Add(excuteMacroButton)
}