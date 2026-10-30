package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	err := run(os.Args, os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	/*	commandLine := os.Args
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
	*/
}

func run(args []string, out io.Writer) error {
	if len(args) < 2 {
		return errors.New("usage: run <command> <args>")
	}
	fileName := args[1]
	_, err := os.Stat(fileName)
	if err != nil {
		return fmt.Errorf("cannot open file '%s': %w", fileName, err)
	}
	return nil

}
