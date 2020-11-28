package main

import (
	"fmt"
	"log"
	"time"

	"github.com/otiai10/gosseract/v2"
)

func main() {
	iterations := 50
	total := time.Duration(0)

	for i := 0; i < iterations; i++ {
		start := time.Now()

		fileName := "files/bode.jpg"

		// fast model
		// https://github.com/tesseract-ocr/tessdata_fast/raw/master/por.tra
		instance = gosseract.NewClient()

		tm := TextMethod{
			Name:     "tesseract",
			Language: "por",
			Variables: map[string]string{
				"tessedit_pageseg_mode": "3", // auto page segmentation mode
				"load_system_dawg":      "0", // removing dict to increase recognition
				"load_freq_dawg":        "0",
			},
			Client: instance,
		}
		defer instance.Close()

		// Handler tesseract's settings
		tm.tesseractSettings()

		bytesImage := convertToBytes(fileName)
		object, err := tm.extract(bytesImage)
		if err != nil {
			return
		}

		if object != nil {
			fmt.Println(*object)
		}
		elapsed := time.Since(start)
		log.Printf("took %s", elapsed)
		total += elapsed
	}
	log.Printf("All Duration: %s", total)
	log.Printf("loops: %d", iterations)
	log.Printf("Average: %s", total/time.Duration(iterations))
}
