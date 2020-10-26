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

	// tesseract parameters:
	// http: //www.sk-spell.sk.cx/tesseract-ocr-parameters-in-302-version

	// auto page segmentation mode
	client.SetVariable("tessedit_pageseg_mode", "3")
	// allow only numerics
	client.SetVariable("tessedit_char_whitelist", "0123456789")
	// exclude fragments that do not look like whole characters from training and adaption
	client.SetVariable("classify_character_fragments_garbage_certainty_threshold", "1")

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
