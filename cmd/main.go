package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	commandLine := os.Args
	numberOfLinesToPrint := 10
	if len(commandLine) != 2 {
		// TODO: need to revisit when allowing flags like -f, -v, -F, for now just filename
		fmt.Println("Usage: go run main.go <file>")
		os.Exit(1)
	}

	fileName := commandLine[1]
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lineContent []string
	for scanner.Scan() {
		txt := scanner.Text()
		lineContent = append(lineContent, txt)

	}

	var contentToPrint []string
	contentLength := len(lineContent)
	if contentLength < numberOfLinesToPrint {
		contentToPrint = lineContent
	} else {

		contentLength := len(lineContent) - numberOfLinesToPrint
		contentToPrint = lineContent[contentLength:]
	}
	for _, v := range contentToPrint {
		fmt.Println(v)
	}
}
