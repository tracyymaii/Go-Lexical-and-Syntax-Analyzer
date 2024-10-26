package main

import (
	"os"
	"strings"
	"unicode"
)

// "bufio"
// "os"
// "strings"

func main() {

	// opens and reads file
	file, err := os.ReadFile("test1.cpl")
	if err != nil {
		// print out why error occurs
		panic(err)
	}

	// turns file into a string
	fileString := string(file)

	// removes whitespace, new line, and tabs from the file then
	// sends the clean string file the tokenizer
	tokenizer(removeWhiteSpace(fileString))
}

func removeWhiteSpace(entireFile string) string {
	noWhiteSpace := strings.ReplaceAll(entireFile, " ", "")
	noWhiteSpace = strings.ReplaceAll(noWhiteSpace, "\n", "")
	noWhiteSpace = strings.ReplaceAll(noWhiteSpace, "\t", "")

	return noWhiteSpace
}

func tokenizer(entireFile string) {

	//  holds the temp string to tokenize
	var temp strings.Builder
	var curr string
	var currRune rune
	var tempType string
	var currType string

	/*
		match types: LOWER letters/apha, digit, [; , . () = ], alphanumeric

		2 step process: 1. we jsut put it in to the word if its the same
						2. THEN we judge the word and kick it out if its not in the club
	*/

	// parses through the file
	// if string is empty, j add letter, and move onto the next
	// scan letter, if letter does not match type of temp string or its a symbol, process temp string
	// always append char
	for _, character := range entireFile {
		currRune = character
		currType = stringType(currRune)

		curr = string(character)

		if temp.String() == "" {
			tempType = currType
		} else if currType != tempType || tempType == "symbol" {
			//process token
			temp.Reset()
			tempType = currType
		}
		temp.WriteString(curr)
	}
}

func stringType(char rune) string {

	var charType string
	if unicode.IsLetter(char) {
		charType = "letter"
	} else if unicode.IsDigit(char) {
		charType = "digit"
	} else {
		charType = "symbol"
	}

	return charType
}
