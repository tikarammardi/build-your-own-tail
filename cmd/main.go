package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	commandLine := os.Args
	numberOfLinesToPrint := 10
	if len(commandLine) != 4 {
		// TODO: need to revisit when allowing flags like -f, -v, -F, for now just filename
		fmt.Println("Usage: go run main.go <file>")
		os.Exit(1)
	}
	flag := commandLine[1]
	if flag == "-n" {
		num, err := strconv.Atoi(commandLine[2])
		if err != nil {
			fmt.Print("Need number of lines to print")
			os.Exit(1)
		}
		numberOfLinesToPrint = num
	} else {
		fmt.Print("Unsupported flag")
	}

	fileName := commandLine[3]
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
