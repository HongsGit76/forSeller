package main

import (
	// "fmt"
	"macroProj/macro/client"
	"os"
)

func main(){
	os.Setenv("FYNE_FONT","NanumGothicBold.ttf")
	// if os.Getenv("PERMITION") == "true" {
	// 	fmt.Println("good!")
	// } else {
	// 	fmt.Println("bad")
	// }
	client.Client()
}