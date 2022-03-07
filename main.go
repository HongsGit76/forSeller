package main

import (
	"os"
	"macroProj/macro/client"
)

func main(){
	os.Setenv("FYNE_FONT","NanumGothicBold.ttf")
	client.Client()
}