package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	// "github.com/go-vgo/robotgo"

	savemacro "macroProj/macro/save_macro"
)

func mainMenu(w fyne.Window){
	// menu items
	new_:=fyne.NewMenuItem("New Macro", func() {
		fmt.Println("새매크로 입력")
	})
	save_:=fyne.NewMenuItem("Save Macro", func() {
		fmt.Println("매크로 저장")
	})
	load_:=fyne.NewMenuItem("Load Macro", func() {
		fmt.Println("매크로 불러오기")
	})
	// menu drawer
	fileMenuDrawer := fyne.NewMenu("File", new_, save_, load_)

	start_:=fyne.NewMenuItem("Start", func() {
		fmt.Println("매크로 시작")
	})
	stop_:=fyne.NewMenuItem("Stop", func() {
		fmt.Println("매크로 중지")
	})
	startMenuDrawer := fyne.NewMenu("Start", start_, stop_)

	mouseSettings_:=fyne.NewMenuItem("Mouse Settings", func() {
		fmt.Println("마우스 설정")
	})
	keyboardSettings_:=fyne.NewMenuItem("Keyboard Settings", func() {
		fmt.Println("키보드 설정")
	})
	otherSettings_:=fyne.NewMenuItem("Other Settings", func() {
		fmt.Println("기타 설정")
	})
	settingMenuDrawer := fyne.NewMenu("Settings", mouseSettings_, keyboardSettings_, otherSettings_)

	information_:=fyne.NewMenuItem("Information", func() {
		fmt.Println("정보")
	})
	infoMenuDrawer := fyne.NewMenu("Information", information_)

	// set main menu
	main_menu := fyne.NewMainMenu(fileMenuDrawer, startMenuDrawer, settingMenuDrawer, infoMenuDrawer)
	w.SetMainMenu(main_menu)
}

func prependTo(g *fyne.Container, s string) {
	g.Objects = append([]fyne.CanvasObject{widget.NewLabel(s)}, g.Objects...)
	g.Refresh()
}

func excuteMacro (macroInput []string) {
	for _, macroID := range macroInput {
		macroMode := strings.Split(macroID, ":")
		switch macroMode[0] {
		case "Move":
			fmt.Println(macroMode[1])
		case "KeyUP":
			fmt.Println(macroMode[1])
		case "KeyDown":
			fmt.Println(macroMode[1])
		case "Mouse Down":
			fmt.Println(macroMode[1])
		case "Mouse Up":
			fmt.Println(macroMode[1])
		default :
			fmt.Println(macroMode[1])
		}
	}
}

func macroTesting () {
	fmt.Println("hihi")
}

func main() {
	fmt.Println("Macro Develop Start!!")

	// window init
	macroApp := app.New()
	mainWindow := macroApp.NewWindow("Hongs App!")
	mainWindow.Resize(fyne.NewSize(400,400))

	mainMenu(mainWindow)

	mainLavel := widget.NewLabel("Hongs Macro!")
	macroBox := container.NewVBox()
	buttonBox := container.NewVBox()

	keyboardMode := false 
	mouseMode := false

	keyboardModeState := widget.NewLabel(fmt.Sprint(keyboardMode))
	mouseModeState := widget.NewLabel(fmt.Sprint(mouseMode))
	
	buttonBox.Add(keyboardModeState)
	buttonBox.Add(mouseModeState)
	
	keboardMacro := widget.NewButton("Keyboard Macro", func ()  {
		if keyboardMode { keyboardMode = false } else { keyboardMode = true }
		keyboardModeState.SetText(fmt.Sprint(keyboardMode))
	})

	mouseMacro := widget.NewButton("Mouse Macro", func ()  {
		if mouseMode { mouseMode = false} else { mouseMode = true }
		mouseModeState.SetText(fmt.Sprint(mouseMode))
	})

	mouseDown := widget.NewButton("Mouse Down", func ()  {
		fmt.Println("Mouse Down")
	})

	mouseUp := widget.NewButton("Mouse Up", func ()  {
		fmt.Println("Mouse Up")
	})

	mouseClick := widget.NewButton("Mouse Click", func ()  {
		fmt.Println("Mouse Click")
	})

	// 마우스모드 키보드모드 설정하기
	// if keyboardMode {
	// 	button1.Enable()
	// 	button2.Disable()
	// } else if mouseMode {
	// 	button1.Disable()
	// 	button2.Enable()
	// } else {
	// 	button1.Enable()
	// 	button2.Enable()
	// }

	buttonBox.Add(keboardMacro)
	buttonBox.Add(mouseMacro)
	buttonBox.Add(mouseDown)
	buttonBox.Add(mouseUp)
	buttonBox.Add(mouseClick)

	macroCard := widget.NewCard("", "macro", container.NewVScroll(macroBox))
	buttonCard := widget.NewCard("", "macro type", buttonBox)

	// mainWindow.Canvas().SetOnTypedRune(func(r rune) {
	// 	prependTo(macroBox, "Rune: "+string(r))
	// })
	// mainWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent) {
	// 	prependTo(buttonBox, "Key : "+string(ev.Name))
	// })

	var macroInput []string

	// Keyboard Macro input
	if deskCanvas, ok := mainWindow.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ev *fyne.KeyEvent) {
			if keyboardMode {
				prependTo(macroBox, "KeyDown: "+string(ev.Name))
				macroInput = append(macroInput, "KeyDown:"+string(ev.Name))
			} 
		})
		deskCanvas.SetOnKeyUp(func(ev *fyne.KeyEvent) {
			if keyboardMode {
				prependTo(macroBox, "KeyUp: "+string(ev.Name))
				macroInput = append(macroInput, "KeyUp:"+string(ev.Name))
			}
		})
	}
	// Mouse Macro input
	mainWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent)  {
		if ev.Name == "F7" && !keyboardMode && mouseMode{
			act := savemacro.Save_mouse("Move:")
			prependTo(macroBox, act.GetString())
			macroInput = append(macroInput, act.GetString())
			fmt.Println(act)
		} else if ev.Name == "F5" {
			fmt.Println("Press F5")
			excuteMacro(macroInput)
		}
	})

	mainContainer := container.NewBorder(mainLavel, nil, nil, nil,
		container.NewGridWithColumns(2,macroCard,buttonCard),
	)
		
	mainWindow.SetContent(mainContainer)

	// run
	mainWindow.ShowAndRun()
}