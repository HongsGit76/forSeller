package client

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

var selectedId int

func remove_macro(slice []string, n int) []string {
	return append(slice[:n], slice[n+1:]...)
}

// addButton is add buttons on right side of main windows
func addButton(btBox *fyne.Container) {
	// current state message
	noticeString := widget.NewLabel(DefaultModeString)

	// {button} keyboard macro
	keboardMacro := widget.NewButton("키보드 매크로", func ()  {
		if KeyboardMode {
			MouseMode = false
			KeyboardMode = false
			noticeString.SetText(DefaultModeString)
		}
		KeyboardMode = true
		MouseMode = false
		noticeString.SetText(KeyboardModeString)
	})
	
	// {button} mouse position macro
	mouseMacro := widget.NewButton("마우스 매크로 (F7)", func ()  {
		if MouseMode {
			MouseMode = false
			KeyboardMode = false
			noticeString.SetText(DefaultModeString)
	}
		MouseMode = true
		KeyboardMode = false
		noticeString.SetText(MouseModeString)
	})

	mouseDownBox := container.NewHBox()
	mouseUpBox := container.NewHBox()
	mouseClickBox := container.NewHBox()

	// {button} mouse act macro
	mouseDown_L := widget.NewButton(LMouseDown, func ()  {
		MacroInput = append(MacroInput, LMouseDown)
		MacroData.Reload()
	})
	mouseDown_R := widget.NewButton(RMouseDown, func ()  {
		MacroInput = append(MacroInput, RMouseDown)
		MacroData.Reload()
	})
	mouseDownBox.Add(mouseDown_L)
	mouseDownBox.Add(mouseDown_R)

	mouseUp_L := widget.NewButton(LMouseUp, func ()  {
		MacroInput = append(MacroInput, LMouseUp)
		MacroData.Reload()
	})
	mouseUp_R := widget.NewButton(RMouseUp, func ()  {
		MacroInput = append(MacroInput, RMouseUp)
		MacroData.Reload()
	})
	mouseUpBox.Add(mouseUp_L)
	mouseUpBox.Add(mouseUp_R)

	mouseClick_L := widget.NewButton(LMouseClick, func ()  {
		MacroInput = append(MacroInput, LMouseClick)
		MacroData.Reload()
	})
	mouseClick_R := widget.NewButton(RMouseClick, func ()  {
		MacroInput = append(MacroInput, RMouseClick)
		MacroData.Reload()
	})
	mouseClickBox.Add(mouseClick_L)
	mouseClickBox.Add(mouseClick_R)

	entry := widget.NewEntry()
	entry.Validator = validation.NewRegexp(`[^a-z]`, "숫자만 입력해 주세요.")

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "시간 추가", Widget: entry}},
		OnSubmit: func() { // optional, handle form submission
			MacroInput = append(MacroInput, fmt.Sprintf("   시간 추가:%s", entry.Text))
			MacroData.Reload()
		},
	}

	excuteMacroButton := widget.NewButton("매크로 수행(F5)", func ()  {
		if len(MacroInput) == 0 {
			noticeString.SetText(NoMacroListString)
		} else {
			noticeString.SetText(StopString)
			RunMacro()
		}
	})

	delMacro := widget.NewButton("매크로 삭제", func ()  {
		if len(MacroInput) - 1 < selectedId {
			return
		}
		MacroInput = remove_macro(MacroInput, selectedId)
		MacroData.Reload()
	})

	btBox.Add(noticeString)
	btBox.Add(keboardMacro)
	btBox.Add(mouseMacro)
	btBox.Add(mouseDownBox)
	btBox.Add(mouseUpBox)
	btBox.Add(mouseClickBox)
	btBox.Add(excuteMacroButton)
	btBox.Add(form)
	btBox.Add(delMacro)
}