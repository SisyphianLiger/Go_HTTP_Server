package main

import (
	"log"
	"os"
)

const filePath = "messages.txt"

func main() {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("No File cannot process information from %v with %v\n", filePath, err)
		return
	}
	


	printChannels(getLinesChannel(file))

	defer file.Close()


}
