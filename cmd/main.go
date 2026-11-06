package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(args []string, out io.Writer) error {
	if len(args) < 2 {
		return errors.New("usage: run <command> <args>")
	}
	flag := args[1]
	var fileName string
	var line int
	if flag == "-n" {

		n, err1 := strconv.Atoi(args[2])
		line = n
		if err1 != nil {
			return fmt.Errorf("illegal offset -- %s: Invalid argument", args[2])
		}

		fileName = args[3]
	} else {
		fileName = args[1]
	}
	_, err := os.Stat(fileName)
	if err != nil {
		return fmt.Errorf("cannot open file '%s': %w", fileName, err)
	}

	file, err := os.Open(fileName)

	if err != nil {
		os.Exit(1)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lineContent []string
	for scanner.Scan() {
		txt := scanner.Text()
		lineContent = append(lineContent, txt)
	}
	var contentToPrint []string
	numberOfLinesToPrint := 10
	if flag == "-n" {
		numberOfLinesToPrint = line
	}
	contentLength := len(lineContent)
	if contentLength < numberOfLinesToPrint {

		contentToPrint = lineContent
	} else {
		contentLength := contentLength - numberOfLinesToPrint
		contentToPrint = lineContent[contentLength:]
	}
	for _, v := range contentToPrint {
		fmt.Fprintln(out, v)
	}
	return nil
}
