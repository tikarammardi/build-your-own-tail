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

	fileName := commandLine[1]
	fmt.Println("File Name:", fileName)
	file, err := os.Open("test.log") // For read access.
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
