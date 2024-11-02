package main

/*
Title : Go Parser
Purpose: To search for lexicode and syntax errors based on given grammar
Author: Tracy Mai
Date: October 27, 2024
*/

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

type data struct {
	x string
	y string
}

var dataPoints = make(map[string]data)

var shape bool

var tokenPatterns = map[string]string{
	"ID":        `^[a-z]+`,
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

var stringGrammar = map[string]string{
	"ID":        "lowercase letters a-z",
	"NUM":       "numbers 0-9",
	"SEMICOLON": ";",
	"COMMA":     ",",
	"PERIOD":    ".",
	"LPAREN":    "(",
	"RPAREN":    ")",
	"ASSIGN":    "=",
	"POINT":     "point",
	"TRIANGLE":  "triangle",
	"SQUARE":    "square",
	"TEST":      "test",
}

var collectedTokens = make([]grammar, 0)

var idPatterns = []string{
	"ID", "ASSIGN", "POINT", "LPAREN", "NUM", "COMMA", "NUM", "RPAREN",
}

var testPattern = []string{
	"TEST", "LPAREN",
}

var trianglePattern = []string{
	"TRIANGLE", "COMMA", "ID", "COMMA", "ID", "COMMA", "ID", "RPAREN",
}

var squarePattern = []string{
	"SQUARE", "COMMA", "ID", "COMMA", "ID", "COMMA", "ID", "COMMA", "ID", "RPAREN",
}

var triangleNames = []string{
	"line", "vertical", "horizontal", "equilateral", "isosceles", "right", "scalene", "obtuse",
}

/*
Main
Combines all the methods of the program to correctly function
@param none
@return none
*/

func main() {

	filepattern := regexp.MustCompile(`^[a-z0-9]+\.[a-z]+$`)
	if len(os.Args) < 3 {
		errorMessage := `Missing parameter, usage:
go run . filename - flag
flag can be p for Prolog generation
flag can be s for Scheme generation`
		panic(errorMessage)

	} else if os.Args[1] == "" {
		errorMessage := `Missing parameter, file:
go run . filename - flag
flag can be p for Prolog generation
flag can be s for Scheme generation`
		panic(errorMessage)

	} else if os.Args[2] == "" {
		errorMessage := `Missing parameter, usage:
go run . filename - flag
flag can be p for Prolog generation
flag can be s for Scheme generation`
		panic(errorMessage)

	} else if !filepattern.MatchString(os.Args[1]) || (strings.ToLower(os.Args[2]) != "-s" && strings.ToLower(os.Args[2]) != "-p") {
		errorMessage := `Incorrect syntax for filename and flag:
go run . filename - flag
flag can be p for Prolog generation
flag can be s for Scheme generation`
		panic(errorMessage)

	}

	fileData, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic("Cannot read file")
	}

	fileString := string(fileData)

	tokenizer(removeWhiteSpace(fileString))

	checkSyntax()
	organizeVariables()

	if strings.ToLower(os.Args[2]) == "-s" {
		scheme()
	} else if strings.ToLower(os.Args[2]) == "-p" {
		prolog()
	}
}

/*
Removes White Space
Removes the whitespace, newlines, and tabs from the file
@param the file
@return the file without whitespace
*/
func removeWhiteSpace(entireFile string) string {
	noWhiteSpace := strings.ReplaceAll(entireFile, " ", "")
	noWhiteSpace = strings.ReplaceAll(noWhiteSpace, "\n", "")
	noWhiteSpace = strings.ReplaceAll(noWhiteSpace, "\t", "")

	return noWhiteSpace
}

/*
Tokenizer
Splits up the file into tokens and uses Process Tokens to check if it matches the grammar
@param the file
@return none
*/
func tokenizer(entireFile string) {

	var temp strings.Builder
	var curr string
	var currRune rune
	var tempType string
	var currType string

	compiledPatterns := compiledRegex()
	compiledPatterns2 := compiledRegex2()

	for _, character := range entireFile {
		currRune = character
		currType = stringType(currRune)
		curr = string(character)

		if temp.String() == "" {
			tempType = currType
		} else if tempType == "symbol" || currType != tempType {
			processToken(temp.String(), compiledPatterns, compiledPatterns2)

			temp.Reset()
			tempType = currType
		}
		temp.WriteString(curr)
	}

	processToken(temp.String(), compiledPatterns, compiledPatterns2)
}

/*
Process Token
Process tokens and checks to see if they exist in the grammar
Exits the program, if a token is not in the grammar
@param the token, compiled maps of the regex patterns
@return none
*/
func processToken(passedToken string, compiledPatterns map[string]*regexp.Regexp, compiledPatterns2 map[string]*regexp.Regexp) {

	for key, regex := range compiledPatterns2 {
		if regex.MatchString(passedToken) {
			collectedTokens = append(collectedTokens, grammar{token: key, lexeme: passedToken})
			return
		}
	}

	for key, regex := range compiledPatterns {
		if regex.MatchString(passedToken) {
			collectedTokens = append(collectedTokens, grammar{token: key, lexeme: passedToken})
			return
		}
	}

	panic("Lexicode error " + passedToken + " not recognized")
}

/*
Check Syntax
Uses Match Syntax to check the order of the tokens to our expected format
@param none
@return none
*/
func checkSyntax() {

	for i := 0; i < len(collectedTokens); i++ {

		switch collectedTokens[i].token {
		case "ID":
			if !shape {
				i = matchSyntax(idPatterns, i)
			}
		case "TEST":
			i = matchSyntax(testPattern, i)
		case "TRIANGLE":
			shape = true
			i = matchSyntax(trianglePattern, i)

		case "SQUARE":
			shape = true
			i = matchSyntax(squarePattern, i)
		}
	}
}

/*
Match Syntax
Uses check the order of the tokens to our expected format
Exits if there is a syntax error
@param the pattern we are matching, index of the current array
@return none
*/
func matchSyntax(pattern []string, index int) int {

	var i int

	for i := 0; i < len(pattern); i++ {
		if collectedTokens[index+i].token == pattern[i] {
			continue
		} else {
			panic("Syntax error " + collectedTokens[index+i].lexeme + " found " + pattern[i] + " expected")
		}
	}

	if collectedTokens[i+index+1].token == "SEMICOLON" && i+index+1 == len(collectedTokens)-1 {
		panic("Syntax error ; found . expected")
	}
	if collectedTokens[i+index+1].token == "PERIOD" && i+index+1 != len(collectedTokens)-1 {
		panic("Syntax error . found ; expected")
	}

	return i + index + 1
}

/*
Organize Variables
Groups the variables and their x and y values into structs
@param none
@return none
*/
func organizeVariables() {

	var id string
	var xVal string
	var yVal string

	for _, items := range collectedTokens {
		if items.token == "ID" {
			id = items.lexeme
		} else if items.token == "NUM" && xVal == "" {
			xVal = items.lexeme
		} else if items.token == "NUM" {
			yVal = items.lexeme
		}

		if id != "" && xVal != "" && yVal != "" {
			dataPoints[id] = data{x: xVal, y: yVal}
			id = ""
			xVal = ""
			yVal = ""
		}
	}
}

/*
Scheme
Runs the scheme output
@param none
@return none
*/
func scheme() {

	shapeSeen := false

	output := "\n; processing Input File input.txt \n; Lexical and Syntax analysis passed\n; Generating Scheme Code"

	for i := 0; i < len(collectedTokens); i++ {
		if collectedTokens[i].token == "TRIANGLE" || collectedTokens[i].token == "SQUARE" {
			shapeSeen = true
			output += "\n(process-" + collectedTokens[i].lexeme
		} else if collectedTokens[i].token == "SEMICOLON" {
			shapeSeen = false
		} else if collectedTokens[i].token == "ID" && shapeSeen == true {
			pointData, exists := dataPoints[collectedTokens[i].lexeme]
			if exists {
				output += " (make-point " + pointData.x + " " + pointData.y + ")"
			} else {
				panic("Attempted to use a data point that does not exist")
			}
		} else if collectedTokens[i].token == "TEST" {
			output += ")"
			shapeSeen = false
		}
	}
	output += ")"

	fmt.Printf("%s\n", output)
}

/*
Prolog
Runs the prolog output
@param none
@return none
*/
func prolog() {
	shapeSeen := false

	fmt.Println("\n/* processing input file input.txt \n   Lexical and Syntax analysis passed \n   Generating Prolog Code */")

	output := ""
	squareOutput := ""

	for i := 0; i < len(collectedTokens); i++ {
		if collectedTokens[i].token == "TRIANGLE" || collectedTokens[i].token == "SQUARE" {
			shapeSeen = true
			output += "\n/* Processing test(" + collectedTokens[i].lexeme
			squareOutput += "\nquery(" + collectedTokens[i].lexeme + "("

		} else if collectedTokens[i].token == "ID" && shapeSeen == true {
			pointData, exists := dataPoints[collectedTokens[i].lexeme]
			if exists {

				output += ", " + collectedTokens[i].lexeme

				squareOutput += "point2d(" + pointData.x + ", " + pointData.y + "),"

			} else {
				panic("Attempted to use a data point that does not exist")
			}
		} else if collectedTokens[i].token == "TEST" || collectedTokens[i].token == "PERIOD" {
			if shapeSeen {
				output += ") */"
			}
			if len(squareOutput) > 0 {
				squareOutput = squareOutput[:len(squareOutput)-1]
			}
			if shapeSeen {
				squareOutput += "))."
			}
			shapeSeen = false

			fmt.Printf("%s\n", output)
			fmt.Printf("%s", squareOutput)

			lines := strings.Split(squareOutput, "\n")

			var filteredLines []string

			for _, line := range lines {
				if strings.Contains(line, "triangle") {
					filteredLines = append(filteredLines, line)
				}
			}

			for j := 0; j < len(filteredLines); j++ {
				temp := filteredLines[j]
				for i := 0; i < len(triangleNames); i++ {
					squareOutput = strings.Replace(squareOutput, "triangle", triangleNames[i], 1)
					fmt.Printf("%s\n", squareOutput)
					squareOutput = temp
				}
			}

			output = ""
			squareOutput = ""
		}
	}

	fmt.Println("\n\n/* Query Processing */\nwriteln(T) :- write(T), nl.\nmain:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),\n\thalt.")
}

/*
Compiled Regex
Compiles the regex patterns for tokenPatterns
@param none
@return none
*/
func compiledRegex() map[string]*regexp.Regexp {
	compiledPatterns := make(map[string]*regexp.Regexp)
	for key, pattern := range tokenPatterns {
		compiledPatterns[key] = regexp.MustCompile(pattern)
	}

	return compiledPatterns
}

/*
Compiled Regex
Compiles the regex patterns for tokenPatterns2
@param none
@return none
*/
func compiledRegex2() map[string]*regexp.Regexp {
	compiledPatterns2 := make(map[string]*regexp.Regexp)
	for key, pattern := range tokenPatterns2 {
		compiledPatterns2[key] = regexp.MustCompile(pattern)
	}

	return compiledPatterns2
}

/*
stringType
Converts runes to string and determines if they're a digit, letter, or special character
@param rune
@return string
*/
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
