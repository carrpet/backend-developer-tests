package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Make sure input is being piped to this program. STDIN should be a
	// named pipe.
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeNamedPipe != 0 {
		fmt.Println("Please pipe input to this program.")
		return
	}

	// returnErrLinesFunc is a split function to customize the bufio.Scanner to return
	// only lines that contain the string "error"
	returnErrLinesFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		if err == nil && token != nil {
			if bytes.Contains(token, []byte("error")) {
				return advance, token, nil
			}
			return advance, []byte{}, nil
		}
		return advance, token, err
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(returnErrLinesFunc)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		text := scanner.Text()
		if len(text) > 0 {
			_, err := writer.WriteString(fmt.Sprintln(text))
			if err != nil {
				log.Printf("Error writing output: %s", err.Error())
			}
			writer.Flush()
		}
	}
	writer.Flush()

}
