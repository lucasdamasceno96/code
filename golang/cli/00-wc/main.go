package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// Calling the count function
	// Received from the stantar input and printing it out to the standard output
	fmt.Println(count(os.Stdin))
}

func count(r io.Reader) int {
	// Creating a new scanner to read from the reader
	scanner := bufio.NewScanner(r)

	// Setting the scanner split function to words
	scanner.Split(bufio.ScanWords)

	// Initializing the count variable
	wc := 0

	// Looping over the scanner
	for scanner.Scan() {
		// Incrementing the count variable
		wc++
	}

	// Returning the total
	return wc
}
