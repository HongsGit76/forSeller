package client

import (
	"fmt"

	"fyne.io/fyne/v2"
)

// MainMenu is define menus in top of main window
func mainMenu(w fyne.Window) {
	// menu items
	new_:=fyne.NewMenuItem("New Macro", func() {
		clear_macro()
	})
	save_:=fyne.NewMenuItem("매크로 저장", func() {
		save_macro(w)
	})
	load_:=fyne.NewMenuItem("Load Macro", func() {
		fmt.Println("매크로 불러오기")
	})
	// menu drawer
	fileMenuDrawer := fyne.NewMenu("File", new_, save_, load_)

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
	main_menu := fyne.NewMainMenu(fileMenuDrawer, settingMenuDrawer, infoMenuDrawer)
	w.SetMainMenu(main_menu)
}