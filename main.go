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

	fileName := "files/bode.jpg"
	bytesImage := convertToBytes(fileName)

	tm := TextMethod{
		Name:     "tesseract",
		Language: "eng",
		Variables: map[string]string{
			"tessedit_pageseg_mode": "3", // auto page segmentation mode
			"load_system_dawg":      "0", // removing dict to increase recognition
			"load_freq_dawg":        "0",
		},
		Client: nil,
	}
	instance = gosseract.NewClient()
	defer instance.Close()
	tm.Client = instance

	tm.tesseractSettings()

	for i := 0; i < iterations; i++ {
		start := time.Now()

		// Handler tesseract's settings
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
