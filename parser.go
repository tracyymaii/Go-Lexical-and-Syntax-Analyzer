package main

import (
	"fmt"
	"os"
	"regexp"
	
	"bufio"
	
	//"strings"
)

// type point struct {
// 	variable string
// 	xValue   string
// 	yValue   string
// }



type tokenType string

const (
	POINT     tokenType = "POINT"
	ID        tokenType = "ID"
	NUM       tokenType = "NUM"
	SEMICOLON tokenType = "SEMICOLON"
	COMMA     tokenType = "COMMA"
	PERIOD    tokenType = "PERIOD"
	LPAREN    tokenType = "LPAREN"
	RPAREN    tokenType = "RPAREN"
	ASSIGN    tokenType = "ASSIGN"
	TRIANGLE  tokenType = "TRIANGLE"
	SQUARE    tokenType = "SQUARE"
	TEST      tokenType = "TEST"
)

type Token struct {
	Type  tokenType
	Value string
}

/*
a = point ( 2 , 3 ) ;
b = point ( 1 , 1 ) ;
c = point ( 1 , 3 ) ;
d = point ( 2 , 1 ) ;
test ( triangle , a , b , c ) ;
test ( square , a, b , c , d ) .
*/

var tokenRegex = map[tokenType]string{
	POINT:     `^point\b`,
	ID:        `^[a-z]+[a-z0-9]*`,
	NUM:       `\d+`,
	SEMICOLON: `;`,
	COMMA:     `,`,
	PERIOD:    `\.`,
	LPAREN:    `\(`,
	RPAREN:    `\)`,
	ASSIGN:    `=`,
	TRIANGLE:  `triangle\b`,
	SQUARE:    `square\b`,
	TEST:      `test\b`,
}



func main() {

	/*
		steps?:
		1.  parses through, tokenizing, keep track/store of var and lexeme
		2. needs to keep track of tokens, do in a map, key = id, value = #
		2b. keep track of shap, triangle, sqare
		3. while parsing need to watch for lexical and syntax errors
		4. stop and print if have errors
		4a. lexicode u can do
		4b. syntax not yet
		5. assume now no errors, then if-else for scheme or prolog format

	*/

	//lexicode error = invalid token, unrecognized ssymbols, caps,

	// not . () , = ;
	// syntax error: missing brackets, words that should be there
	/// pont instead of point
	// @ instead of number
	/**
		var fileName string
		var schemeOrProlog string

		fmt.Print("prompt> go run . ")
		fmt.Scan(&fileName)
		fmt.Scan(&schemeOrProlog)
	**/
	file, err := os.Open("test1.cpl")
	if err != nil {
		panic(err)
	}

	

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    	line := scanner.Text() // 'line' is of type string
    	tokenizer(line)    // Prints the line as a string
	}
 
		// append it if return a value

	// textFile := string(file)
	// tokenizer(textFile)

	
}

func tokenizer(line string) {
	/*
		you have 2 structs, tokenType string --> used more for printing so then youcan njsut say the var and it will print
		token regex, --> searches for pattern


		through file , if pattern is in tokenrege, print out tokentype,
							if token = id or num, needs to be saved
	*/

	// compiles entire map of pattern off the bat
	compiledPattern := make(map[tokenType]*regexp.Regexp)
	for key, pattern := range tokenRegex {
		compiledPattern[key] = regexp.MustCompile(pattern)		
	}

	matches := compiledPattern.FindAllString(line, -1)
	if len(matches) > 0 {
		fmt.Println(matches)
	}

	
	

/*
	for key, reg := range compiledPattern {
		if reg.MatchString(line) {
			fmt.Println(key)
		}
	}
*/
	




}
