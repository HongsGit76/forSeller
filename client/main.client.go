package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

var MacroData = binding.BindStringList(
	&MacroInput,
)

// Client is main function of frontend
func Client(){
	macroApp := app.New()
	mainWindow := macroApp.NewWindow("1000 셀러 매크로 프로그램")
	mainWindow.Resize(fyne.NewSize(WindowSize.widthWindow, WindowSize.heightWindow))

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

	// make card with list views
	macroCard := widget.NewCard("", "macro", mBox)
	buttonCard := widget.NewCard("", "macro type", buttonBox)

	mainContainer := container.NewBorder(mainLabel, nil, nil, nil,
		container.NewGridWithColumns(2,macroCard,buttonCard),
	)
	mainWindow.SetContent(mainContainer)
	
	// run
	mainWindow.ShowAndRun()
}