package main

import "fmt"

func main() {

	tm := TextMethod{
		Name:     "tesseract",
		Language: "por+eng",
		Variables: map[string]string{
			"tessedit_pageseg_mode": "3", // auto page segmentation mode
			"load_system_dawg":      "0", // removing dict to increase recognition
			"load_freq_dawg":        "0",
		},
	}

	bytesImage := convertToBytes("files/2.jpg")

	object, err := tm.extract(bytesImage)
	if err != nil {
		return
	}

	if object != nil {
		fmt.Print(*object)
	}
}
