package main

import (
	"bufio"
	"bytes"
	"fmt"
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

	// Read STDIN into a new buffered reader

	splitFunc := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanLines(data, atEOF)
		if err == nil && token != nil {
			if bytes.Contains(token, []byte("error")) {
				return advance, token, nil
			}
			return advance, []byte{}, nil
		}
		return advance, token, err

	}
	// TODO: Look for lines in the STDIN reader that contain "error" and output them.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(splitFunc)
	writer := bufio.NewWriter(os.Stdout)
	for scanner.Scan() {
		w, _ := writer.WriteString(fmt.Sprintln(scanner.Text()))
		if w == 0 {
			fmt.Println("Wrote 0 bytes")
		}
	}
	writer.Flush()

}
