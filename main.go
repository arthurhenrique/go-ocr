package main

import (
	"fmt"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	client := gosseract.NewClient()
	defer client.Close()

	// configurations
	client.SetConfigFile("./config/test")
	client.SetLanguage("por+eng")
	client.SetImage("./files/2.jpg")

	// tesseract parameters
	//http://www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version
	client.SetVariable("classify_bln_numeric_mode", "1")

	text, err := client.Text()
	if err != nil {
		fmt.Printf("error %s\n", err)
		return
	}

	// tesseract version
	fmt.Println(client.Version())

	// all bounding boxes
	fmt.Println(client.GetBoundingBoxesVerbose())

	// output
	fmt.Println(text)
}
