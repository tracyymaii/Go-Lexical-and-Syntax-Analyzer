package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

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

var collectedTokens []grammar

// func newToken(token string) *grammar {

// 	g := grammar{token: token}
// 	g.lexeme = ""
// 	return &g
// }

func main() {

	/*

				prompt>go run . input.txt -s
				; Processing Input File input.txt
				; Lexical and Syntax analysis passed
				; Generating Scheme Code
				(process-triangle (make-point 2 3) (make-point 1 4) (make-point 3 4))
				prompt>



		func compiledRegex2() map[string]*regexp.Regexp {
			compiledPatterns2 := make(map[string]*regexp.Regexp)
			for key, pattern := range tokenPatterns2 {
				compiledPatterns2[key] = regexp.MustCompile(pattern)
			}

			return compiledPatterns2
		}

		if regex.MatchString(passedToken) {
					fmt.Println(key + " " + passedToken)
					collectedTokens

	*/

	var i, j string

	fmt.Print("prompt>go run .")
	fmt.Scan(&i, &j)

	filepattern := regexp.MustCompile(`/^[a-z0-9]+\.[a-z]+$\b/gm`)

	if i == "" {
		errorMessage := `Missing parameter, file:
		go run . filename - flag
		flag can be p for prolog generation
		flag can be s for prolog generation`

		fmt.Printf("%s\n", errorMessage)
		return
	} else if j == "" {
		errorMessage := `Missing parameter, usage:
		go run . filename - flag
		flag can be p for prolog generation
		flag can be s for prolog generation`

		fmt.Printf("%s\n", errorMessage)
		return
	} else if !filepattern.MatchString(i) || (strings.ToLower(j) != "-s" || strings.ToLower(j) != "-p") {
		errorMessage := `Incorrect syntax for filename and flag:
		go run . filename - flag
		flag can be p for prolog generation
		flag can be s for prolog generation`

		fmt.Printf("%s\n", errorMessage)
		return
	}

	if strings.ToLower(j) == "-s" {
		//checking for fails in code

		// assume should be good by now

		scheme()
	} else {
		// its accepted by p
		print("hello")
	}

	// i should be a file name, j needs to be -s or -p

	// need to write if input wrong, to give t

	// check, if not fail

	// need to move this to after an actual things, aka where we check for lexical and synteax error
	inputSuccess := `; Processing Input File input.txt
		; Lexical and Syntax analysis passed
		; Generating Scheme Code
		(process-triangle (make-point 2 3) (make-point 1 4) (make-point 3 4))
		prompt>`

	fmt.Printf("%s\n", inputSuccess)

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

	fmt.Println(collectedTokens)
}

func removeWhiteSpace(entireFile string) string {
	noWhiteSpace := strings.ReplaceAll(entireFile, " ", "")
	noWhiteSpace = strings.ReplaceAll(noWhiteSpace, "\n", "")
	noWhiteSpace = strings.ReplaceAll(noWhiteSpace, "\t", "")

	print(noWhiteSpace)
	return noWhiteSpace
}

func tokenizer(entireFile string) { // does not yet reject cap letters

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

/*
func newToken(token string) *grammar {

	g := grammar{token: token}
	g.lexeme = ""
	return &g
}

type grammar struct {
	token  string
	lexeme string
}
*/

func processToken(passedToken string, compiledPatterns map[string]*regexp.Regexp, compiledPatterns2 map[string]*regexp.Regexp) {

	for key, regex := range compiledPatterns2 {
		if regex.MatchString(passedToken) {
			fmt.Println(key + " " + passedToken)
			collectedTokens = append(collectedTokens, grammar{token: key, lexeme: passedToken})
			return
		}
		// else panic(err)
	}

	for key, regex := range compiledPatterns {
		if regex.MatchString(passedToken) {
			fmt.Println(key + " " + passedToken)

			if key == "ID" || key == "NUM" {
				collectedTokens = append(collectedTokens, grammar{token: key, lexeme: passedToken})
			} else {
				collectedTokens = append(collectedTokens, grammar{token: key})
			}

		}
		// else panic(err)
	}

}

func scheme() {
	output := `; Processing Input File input.txt
	; Lexical and Syntax analysis passed
	; Generating Scheme Code`

	fmt.Println("(process-")

	fmt.Printf("%s\n", output)
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
