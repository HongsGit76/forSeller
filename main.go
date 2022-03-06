package main

import (
	"fmt"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"

	// "github.com/go-vgo/robotgo"

	macroActs "macroProj/macro/macroActs"
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

func main() {
	fmt.Println("Macro Develop Start!!")
	os.Setenv("FYNE_FONT","NanumGothicBold.ttf")


	// window init
	macroApp := app.New()
	mainWindow := macroApp.NewWindow("1000 셀러 매크로 프로그램 v1.0")
	mainWindow.Resize(fyne.NewSize(450, 450))
	var macroInput []string
	var macroActList []macroActs.Act

	mainMenu(mainWindow)

	mainLabel := widget.NewLabel("매크로 프로그램")
	// macroBox := container.NewVBox()
	
	data := binding.BindStringList(
		&[]string{},
	)
	mBox := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(di binding.DataItem, co fyne.CanvasObject) {
			co.(*widget.Label).Bind(di.(binding.String))
		},
	)

	buttonBox := container.NewVBox()

	keyboardMode := false 
	mouseMode := false

	keyboardModeState := widget.NewLabel("아래 버튼 중 수행할 매크로를 선택하세요.")
	
	buttonBox.Add(keyboardModeState)
	
	keboardMacro := widget.NewButton("키보드 매크로 입력", func ()  {
		if keyboardMode {
			mouseMode = false
			keyboardMode = false
			keyboardModeState.SetText("아래 버튼 중 수행할 매크로를 선택하세요.")
		}
		keyboardMode = true
		mouseMode = false
		keyboardModeState.SetText("키보드를 입력하세요.")
	})

	mouseMacro := widget.NewButton("마우스 위치 입력 (F7)", func ()  {
		if mouseMode {
			mouseMode = false
			keyboardMode = false
			keyboardModeState.SetText("아래 버튼 중 수행할 매크로를 선택하세요.")
	}
		mouseMode = true
		keyboardMode = false
		keyboardModeState.SetText("(F7)을 눌러 현재 마우스 위치를 복사합니다.")
	})

	mouseDown := widget.NewButton("마우스 누르기", func ()  {
		fmt.Println("Mouse Down")
		act := macroActs.Save_act("Mouse:Down",)
		macroActList = append(macroActList, act)
		data.Append("Mouse Down")
	})

	mouseUp := widget.NewButton("마우스 떼기", func ()  {
		fmt.Println("Mouse Up")
		act := macroActs.Save_act("Mouse:Up")
		macroActList = append(macroActList, act)
		data.Append("Mouse Up")
	})

	mouseClick := widget.NewButton("마우스 클릭", func ()  {
		fmt.Println("Mouse Click")
		act := macroActs.Save_act("Mouse:Click")
		macroActList = append(macroActList, act)
		data.Append("Mouse Click")
	})

	mBox.OnSelected = func(id widget.ListItemID) {
		fmt.Println("id: ", id)
		delMacro := widget.NewButton("매크로 삭제", func ()  {
			fmt.Println("Delete Macro")
		})
		buttonBox.Add(delMacro)
	}

	// TODO: 매크로 수행 간 시간 추가하는 매크로

	buttonBox.Add(keboardMacro)
	buttonBox.Add(mouseMacro)
	buttonBox.Add(mouseDown)
	buttonBox.Add(mouseUp)
	buttonBox.Add(mouseClick)

	excuteMacroButton := widget.NewButton("매크로 수행(F5)", func ()  {
		fmt.Println(macroInput)
		macroActs.ExcuteMacro(macroActList)
	})

	buttonBox.Add(excuteMacroButton)

	// macroCard := widget.NewCard("", "macro", container.NewVScroll(macroBox))
	macroCard := widget.NewCard("", "macro", mBox)
	buttonCard := widget.NewCard("", "macro type", buttonBox)

	// Keyboard Macro input
	if deskCanvas, ok := mainWindow.Canvas().(desktop.Canvas); ok {
		deskCanvas.SetOnKeyDown(func(ev *fyne.KeyEvent) {
			if keyboardMode {
				// prependTo(macroBox, "KeyDown: "+string(ev.Name))
				actString := "KeyDown:"+string(ev.Name)
				data.Append(actString)
				macroInput = append(macroInput, actString)
				act := macroActs.Save_keyboard("KeyDown",string(ev.Name))
				macroActList = append(macroActList, act)
			} 
		})
		deskCanvas.SetOnKeyUp(func(ev *fyne.KeyEvent) {
			if keyboardMode {
				// prependTo(macroBox, "KeyUp: "+string(ev.Name))
				actString := "KeyUp:"+string(ev.Name)
				data.Append(actString)
				macroInput = append(macroInput, actString)
				act := macroActs.Save_keyboard("KeyDown",string(ev.Name))
				macroActList = append(macroActList, act)
			}
		})
	}
	// Mouse Macro input
	mainWindow.Canvas().SetOnTypedKey(func(ev *fyne.KeyEvent)  {
		
		if ev.Name == "F7" && !keyboardMode && mouseMode{
			act := macroActs.Save_mouse("Move")
			data.Append(act.GetString())
			// prependTo(macroBox, act.GetString())
			macroInput = append(macroInput, act.GetString())
			macroActList = append(macroActList, act)
			fmt.Println(act)
		} else if ev.Name == "F5" {
			macroActs.ExcuteMacro(macroActList)
		} 
	})

	mainContainer := container.NewBorder(mainLabel, nil, nil, nil,
		container.NewGridWithColumns(2,macroCard,buttonCard),
	)
		
	mainWindow.SetContent(mainContainer)

	// run
	mainWindow.ShowAndRun()
}