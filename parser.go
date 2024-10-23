package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	re := regexp.MustCompile(`^[a-z]+[a-z0-9]*`)
	id := re.FindString(line)
	if id == "test" {
		fmt.Println("TEST")
	} else {
		fmt.Println("ID", id)
	}

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
	if _, err := strconv.Atoi(xVal); err == nil {
		fmt.Println("NUM ", xVal)

		if strings.Contains(line, ",") {
			fmt.Println("COMMA")
		}
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
	if _, err := strconv.Atoi(yVal); err == nil {
		fmt.Println("NUM ", yVal)
	}

	// Tokenizes triangle
	if strings.Contains(line, "triangle") {
		fmt.Println("TRIANGLE")
	}

	/**
	tri := regexp.MustCompile(`,\s*[a-z0-9]+`)
	tri1Comma := tri.FindString(line)
	tri1 := regexp.MustCompile(`[a-z0-9]+`)
	tri1Val := tri1.FindString(tri1Comma)
	if id == "test" {
		ID:
	}
	**/
	/*
	triangle, find biggerm if triangle, then print, nad repeat
	t is the id, an will print out as id, but will save as t1, t2, t3 = ids and then the ids have values assigne to it
	do matchString to get a list of it values after comma, then preset each to null, so its jsut re-assigned in statement, then can be used outside of function
	*/

	tri := regexp.MustCompile(`(,\s*\d+)`)
	triVars := tri.FindAllStringSubmatch(line, -1)

	tri1 := os.DevNull
	tri2 := os.DevNull
	tri3 := os.DevNull

	if len(triVars) > 0 && 


	
		tri1 = match[1]
		fmt.Println("ID", tri1)
		tri2 = match[2]
		fmt.Print
	



	// Tokenizes square
	if strings.Contains(line, "square") {
		fmt.Println("SQUARE")
	}

	// Tokenizes ,
	if strings.Contains(line, ",") {
		fmt.Println("COMMA")
	}

	// Tokenizes )
	if strings.Contains(line, ")") {
		fmt.Println("RPAREN")
	}

	// Tokenizes ;
	if strings.Contains(line, ";") {
		fmt.Println("SEMICOLON")
	}

	// Tokenizes .
	if strings.Contains(line, ".") {
		fmt.Println("PERIOD")
	}
}
