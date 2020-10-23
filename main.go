package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()
	// client.SetConfigFile("/usr/local/Cellar/tesseract/4.1.1/share/tessdata/tessconfigs/test")
	// client.SetVariable(gosseract.SettableVariable("save_blob_choices"), "0")

	client.SetImage("./files/1.jpg")
	client.SetLanguage("por+eng")

	// x, y := client.GetBoundingBoxes(gosseract.PageIteratorLevel(1))
	// fmt.Println(x, y)
	text, err := client.Text()
	if err != nil {
		fmt.Printf("error %s\n", err)
		return
	}

	fmt.Println(text)
}
