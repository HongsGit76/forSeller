package client

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var MacroData = binding.BindStringList(
	&MacroInput,
)

// Client is main function of frontend
func Client(){
	macroApp := app.New()
	mainWindow := macroApp.NewWindow("1000 셀러 매크로 프로그램")
	mainWindow.Resize(fyne.NewSize(WindowSize.width, WindowSize.height))

	mainMenu(mainWindow)

	// main label view
	mainLabel := widget.NewLabel("매크로 프로그램")
	
	// button list view
	buttonBox := container.NewVBox()
	addButton(buttonBox)

	// macro list view
	mBox := widget.NewListWithData(
		MacroData,
		func () fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func (di binding.DataItem, co fyne.CanvasObject)  {
			co.(*widget.Label).Bind(di.(binding.String))
		},
	)

	mBox.OnSelected = func(id widget.ListItemID) {
		selectedId = id
	}

	// make card with list views
	macroCard := widget.NewCard("", "macro", mBox)
	buttonCard := widget.NewCard("", "macro type", buttonBox)

	mainContainer := container.NewBorder(mainLabel, nil, nil, nil,
		container.NewGridWithColumns(2,macroCard,buttonCard),
	)
	mainWindow.SetContent(mainContainer)

	// save mouse position or keyboard macro 
	if deskCanvas, ok := mainWindow.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ev *fyne.KeyEvent) {
			if KeyboardMode {
				MacroInput = append(MacroInput, fmt.Sprintf("   키보드 누름:%s",string(ev.Name)))
				MacroData.Reload()
			} 
		})
		deskCanvas.SetOnKeyUp(func(ev *fyne.KeyEvent) {
			if KeyboardMode {
				MacroInput = append(MacroInput, fmt.Sprintf("   키보드 뗌 :%s",string(ev.Name)))
				MacroData.Reload()
			}
		})
	}

	mainWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent)  {
		if ev.Name == "F7" && !KeyboardMode && MouseMode{
			x,y := GetCurMousePosition()
			MacroInput = append(MacroInput, fmt.Sprintf("   마우스 이동 :%d,%d",x,y))
			MacroData.Reload()
		}
	})
	
	// run
	mainWindow.ShowAndRun()
}