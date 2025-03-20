package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)



func splitStringByNewline(buffer []byte) (string, string) {
	partOne := ""
	partTwo := ""
	input := string(buffer)


	idx := strings.IndexRune(input, '\n') 

	if idx != -1 {
		partOne = string(buffer[:idx])
		partTwo = string(buffer[idx+1:])

	} else {

		partOne = input 
	}

	return partOne, partTwo
}




func printChannels(channel <-chan string) {
	for str := range channel {
		fmt.Printf("read: %v \n", str)
	}
}

func getLinesChannel(file io.ReadCloser) <-chan string {
	
	stringStream := make(chan string) 
	
	go func () {
		defer close(stringStream)
		
		buf := make([]byte, 8)
		stringDisplay := ""
		for {

			buf_len, err := file.Read(buf)

			if err != nil {
				if errors.Is(err, io.EOF) {
					stringStream <- stringDisplay
					break
				}
				fmt.Printf("Error: %s\n", err.Error())
			}


			first,second := splitStringByNewline(buf[:buf_len])


			if second == "" {
				stringDisplay += first
			} else {
				stringDisplay += first
				stringStream <- stringDisplay
				stringDisplay = second
			}
		}
	}()

	return stringStream
}
