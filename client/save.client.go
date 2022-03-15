package client

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
)

func clear_macro(){
	macroLen := len(MacroInput)
	MacroInput = MacroInput[macroLen:]
	MacroData.Reload()
}

func save_macro(win fyne.Window) {
	dialog.ShowFileSave(func(writer fyne.URIWriteCloser, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if writer == nil {
			log.Println("Cancelled")
			return
		}
		fileSaved(writer, win)
	}, win)
}

func fileSaved(f fyne.URIWriteCloser, w fyne.Window) {
	defer f.Close()
	_, err := f.Write([]byte("Written by Fyne demo\n"))
	if err != nil {
		dialog.ShowError(err, w)
	}
	err = f.Close()
	if err != nil {
		dialog.ShowError(err, w)
	}
	log.Println("Saved to...", f.URI())
}

func load_macro(win fyne.Window){
	dialog.ShowFolderOpen(func(list fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		if list == nil {
			log.Println("Cancelled")
			return
		}

		children, err := list.List()
		if err != nil {
			dialog.ShowError(err, win)
			return
		}
		out := fmt.Sprintf("Folder %s (%d children):\n%s", list.Name(), len(children), list.String())
		dialog.ShowInformation("Folder Open", out, win)
	}, win)
}