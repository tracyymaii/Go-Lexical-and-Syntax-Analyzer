package main

import (
	"fmt"
	"regexp"

	"bufio"
	"os"
	"strings"
)

// type point struct {
// 	variable string
// 	xValue   string
// 	yValue   string
// }

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
		line := scanner.Text()

		tokens(line)
	}

	defer file.Close()
}

func tokens(line string) {
	// a = point(2, 3);

	// do id last bc i feel like it,

	// change tokenizing to take care of errors

	// ask about diff between lexeme and syntax error

	// Tokenizes ID
	re := regexp.MustCompile(`^[a-z]+[a-z0-9]*`)
	id := re.FindString(line)
	fmt.Println(id)

	// Tokenizes =
	if strings.Contains(line, "=") {
		fmt.Println("ASSIGN")
	}

	// Tokenizes ,
	if strings.Contains(line, "point") {
		fmt.Println("POINT")
	}

	// Tokenizes (
	if strings.Contains(line, "(") {
		fmt.Println("LPAREN")
	}

	// Tokenizes x value
	x := regexp.MustCompile(`(\d+)`)
	xVal := x.FindString(line)
	fmt.Println(xVal)

	// Tokenizes ,
	if strings.Contains(line, ",") {
		fmt.Println("COMMA")
	}

	// Tokenizes y value
	// y := regexp.MustCompile(`\,(.*?)\)`)
	// ySeen := y.FindStringSubmatch(line)
	// if len(ySeen) > 1 {
	// 	fmt.Println("NUM", ySeen[1])

	// }

	y := regexp.MustCompile(`(,\s*\d+)`)
	yComma := y.FindString(line)
	y2 := regexp.MustCompile(`\d+`)
	yVal := y2.FindString(yComma)
	fmt.Println(yVal)

	//fmt.Println(created)

	// Tokenizes )
	if strings.Contains(line, ")") {
		fmt.Println("RPAREN")
	}

	// Tokenizes ;
	if strings.Contains(line, ";") {
		fmt.Println("SEMICOLON")
	}

}
