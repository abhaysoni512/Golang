package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func errorCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("book.txt")
	errorCheck(err)
	defer file.Close()

	/*
		var wordCount map[string]int
		wordCount["hello"]++  // PANIC! Crash!
		Because var creates a nil map (empty, doesn't exist yet).
		You cannot add anything to a nil map → program crashes.

		Correct ways:
			wordCount := make(map[string]int)           // most common
			// or
			wordCount := map[string]int{}               // also works (empty map)
			// or long way:
			var wordCount = make(map[string]int)
	*/
	wordCount := make(map[string]int)

	// This creates a smart reader that can read text easily.
	/*
		Use bufio.Scanner to read word by word
		Line 1: bufio.NewScanner(file)
		This creates a smart reader that can read text easily.
		Think of scanner like a robot that:

		Holds the file
		Has a 4096-byte buffer (remember the water tank?)
		Can give you data in nice pieces (lines, words, etc.)

		bufio.NewScanner() takes any io.Reader (our file is one) and returns a scanner.
		Line 2: scanner.Split(bufio.ScanWords)
		This is the most important line for reading words!
		By default, NewScanner splits on lines (using bufio.ScanLines)
		→ useful for scanner.Text() giving one full line
		But we want words, not lines!
		So we change the splitting rule:
		Now the robot says: scanner.Split(bufio.ScanWords)
		“From now on, every time someone calls scanner.Scan(),
		give them one word (separated by space, tab, or newline)”



	*/
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// scanner.Scan() → returns true if there is a new token (word), false if file ended
	for scanner.Scan() {
		word := scanner.Text()
		word = strings.ToLower(word)
		word = strings.Trim(word, ".,!?:;\"'()[]{}—-")

		if word != "" {
			wordCount[word]++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	for word, freq := range wordCount {
		fmt.Println(word, "appears ", freq, "times")
	}

}
