package main

import (
	"fmt"

	"github.com/pablodz/go-dtmf2/dtmf"
)

func main() {
	// Testing using file mono, 8000hz files

	fileName := "test/123456654321.raw"
	valueString, err := dtmf.DecodeDTMFFromFile(fileName, 8000.0, 12)
	if err != nil {
		fmt.Println("There is an error", err)
		return
	}
	fmt.Println("Decoded character is", valueString)

	fileName = "test/147258369.raw"
	valueString, err = dtmf.DecodeDTMFFromFile(fileName, 8000.0, 12)
	if err != nil {
		fmt.Println("There is an error", err)
		return
	}
	fmt.Println("Decoded character is", valueString)
}
