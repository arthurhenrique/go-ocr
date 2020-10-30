package main

import (
	"bufio"
	"fmt"
	"os"
)

func convertToBytes(fileName string) []byte {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	fileInfo, _ := file.Stat()
	bytes := make([]byte, fileInfo.Size())

	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	return bytes
}

// func threshold(data []byte, factor uint8) []byte {
// 	if factor == 0 {
// 		return data
// 	}
// 	RedLuminosity := 0.299
// 	GreenLuminosity := 0.587
// 	BlueLuminosity := 0.114

// 	// fmt.Print(data)
// 	img, _, _ := image.Decode(bytes.NewReader(data))
// 	bounds := img.Bounds()

// 	newImg := image.NewRGBA(bounds)
// 	buf := new(bytes.Buffer)
// 	jpeg.Encode(buf, newImg, nil)
// 	return buf.Bytes()
// 	bounds := img.Bounds()

// 	newImg := image.NewRGBA(bounds)

// 	for y := 0; y < bounds.Max.Y; y++ {
// 		for x := 0; x < bounds.Max.X; x++ {
// 			actualPixel := img.At(x, y)
// 			r, g, b, _ := actualPixel.RGBA()

// 			luminosity := RedLuminosity*float64(r) +
// 				GreenLuminosity*float64(g) + BlueLuminosity*float64(b)

// 			pixel := color.Gray{uint8(luminosity / 256)}

// 			if pixel.Y > factor {
// 				newImg.Set(x, y, color.Gray{255})
// 			} else {
// 				newImg.Set(x, y, color.Gray{0})
// 			}
// 		}
// 	}

// 	buf := new(bytes.Buffer)
// 	jpeg.Encode(buf, newImg, nil)

// 	return buf.Bytes()
// }
