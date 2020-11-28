package main

import (
	"bufio"
	"fmt"
	"os"
)

// convertToBytes return bytes
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
