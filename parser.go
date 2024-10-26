package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

// "bufio"
// "os"
// "strings"

type grammar struct {
	token  string
	lexeme string
}

var tokenPatterns = map[string]string{
	"ID":        `^[a-z]+`, // PROBLEM, prints, point, triangle, and squares as
	"NUM":       `^\d+`,
	"SEMICOLON": `^;`,
	"COMMA":     `^,`,
	"PERIOD":    `(^\.$)`,
	"LPAREN":    `^\(`,
	"RPAREN":    `^\)`,
	"ASSIGN":    `^=`,
}

var tokenPatterns2 = map[string]string{
	"POINT":    `^point\b`,
	"TRIANGLE": `^triangle\b`,
	"SQUARE":   `^square\b`,
	"TEST":     `^test\b`,
}

func newToken(token string) *grammar {

	g := grammar{token: token}
	g.lexeme = ""
	return &g
}

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

	print(noWhiteSpace)
	return noWhiteSpace
}

func tokenizer(entireFile string) {

	//  holds the temp string to tokenize
	var temp strings.Builder
	var curr string
	var currRune rune
	var tempType string
	var currType string

	compiledPatterns := compiledRegex()
	compiledPatterns2 := compiledRegex2()

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
		} else if tempType == "symbol" || currType != tempType {
			//process token
			processToken(temp.String(), compiledPatterns, compiledPatterns2)

			temp.Reset()
			tempType = currType
		}

		// fence post to get last char, and converts last rune to char
		// if len(entireFile) > 0 {
		// 	lastRune, _ := utf8.DecodeLastRune([]byte(entireFile))
		// 	processToken(string(lastRune), compiledPatterns)
		// }
		temp.WriteString(curr)
	}

	processToken(temp.String(), compiledPatterns, compiledPatterns2)

	//process period/last token
}

func processToken(passedToken string, compiledPatterns map[string]*regexp.Regexp, compiledPatterns2 map[string]*regexp.Regexp) {

	for key, regex := range compiledPatterns2 {
		if regex.MatchString(passedToken) {
			fmt.Println(key + " " + passedToken)
			return
		}
		// else panic(err)
	}

	for key, regex := range compiledPatterns {
		if regex.MatchString(passedToken) {
			fmt.Println(key + " " + passedToken)

		}
		// else panic(err)
	}

}

func compiledRegex() map[string]*regexp.Regexp {
	compiledPatterns := make(map[string]*regexp.Regexp)
	for key, pattern := range tokenPatterns {
		compiledPatterns[key] = regexp.MustCompile(pattern)
	}

	return compiledPatterns
}

func compiledRegex2() map[string]*regexp.Regexp {
	compiledPatterns2 := make(map[string]*regexp.Regexp)
	for key, pattern := range tokenPatterns2 {
		compiledPatterns2[key] = regexp.MustCompile(pattern)
	}

	return compiledPatterns2
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
