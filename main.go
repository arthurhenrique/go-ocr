package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetConfigFile("./config/test")
	client.SetLanguage("por+eng")
	client.SetImage("./files/2.jpg")

	text, err := client.Text()
	if err != nil {
		fmt.Printf("error %s\n", err)
		return
	}

	fmt.Println(text)
}
