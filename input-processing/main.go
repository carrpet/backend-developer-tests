package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

const twelveMB = 1024 * 1024 * 12

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// These functions were used for testing.
	// They could be moved into an integration test suite for this application.
	//createLongLineWithError("longlineerror.txt")
	//createLongLineNoError("longlinenoerror.txt")
	//createLongFileNoError("longfilenoerror.txt")
	//createLongFileErrors("longfileerrors.txt")

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

	// This scanner has been configured to accept tokens (in this case, lines)
	// of size up to 12 MB. If a token is larger than that the scanner will encounter
	// an error (which is written to the logs).  Additionally, the initial size of the buffer has also
	// been made large (bufio.MaxScanTokenSize) so that new bigger buffers won't continually
	// be allocated if the tokens are up to bufio.MaxScanTokenSize. These parameters can be fine tuned
	// for differing expected inputs to maximize performance.
	scanner.Buffer(make([]byte, bufio.MaxScanTokenSize), twelveMB)

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
	if scanner.Err() != nil {
		log.Printf("Scanner encountered error: %s", scanner.Err().Error())
	}
	writer.Flush()
}

// helper functions for testing the performance of the input processor

func fileExists(fn string) bool {
	_, err := os.Stat(fn)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// createLongLineWithError outputs a file containing one huge
// line of around 10 MB and the string "error"
func createLongLineWithError(fn string) (err error) {
	if fileExists(fn) {
		return
	}
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fs := 1024 * 1024 * 10 // 10 MB
	for i := 0; i < fs; i++ {
		w.WriteRune('x')
	}
	w.WriteString(" error me\n")
	w.Flush()
	return nil
}

// createLongLineNoError outputs a file containing one huge
// line of around 10 MB containing no "error" string
func createLongLineNoError(fn string) (err error) {
	if fileExists(fn) {
		return
	}
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fs := 1024 * 1024 * 10 // 10 MB
	for i := 0; i < fs; i++ {
		w.WriteRune('x')
	}
	w.WriteRune('\n')
	w.Flush()
	return nil
}

// createLongFileNoError outputs a file containing 1024 * 1024 * 2 lines
// with no "error"
func createLongFileNoError(fn string) (err error) {
	if fileExists(fn) {
		return
	}
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fs := 1024 * 1024 * 2 // number of lines
	for i := 0; i < fs; i++ {
		w.WriteString("a t\n")
	}
	w.Flush()
	return nil
}

// createLongFileErrors outputs a file containing about 1024 * 1024 * 3 lines
// containing a few lines with "error"
func createLongFileErrors(fn string) (err error) {
	if fileExists(fn) {
		return
	}
	f, err := os.Create(fn)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	fs := 1024 * 1024 * 1
	for i := 0; i < fs; i++ {
		w.WriteString("a t\n")
	}
	w.WriteString("I have an error in me\n")
	for i := 0; i < fs; i++ { // Another 1 MB of lines
		w.WriteString("a t\n")
	}
	w.WriteString("This is another error\n")
	for i := 0; i < fs; i++ { // Yet another 1 MB of lines
		w.WriteString("a t\n")
	}
	w.WriteString("Final error\n")
	w.Flush()
	return nil
}
