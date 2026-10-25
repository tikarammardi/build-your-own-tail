package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/*
TODO:	1. Accept a single filename as an argument,
		2.Read the file, and
		3.Print only its final 10 lines.
		4. If the file has fewer than 10 lines, print all of them.
*/

func main() {
	commandLine := os.Args
	if len(commandLine) != 2 {
		//TODO: need to revisit when allowing flags like -f, -v, -F, for now just filename
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
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
