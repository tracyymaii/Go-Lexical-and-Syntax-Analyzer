package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/elliotchance/orderedmap/v2"
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

	orderedMap := orderedmap.NewOrderedMap[string, string]()

	orderedMap.Set("ID", `^[a-z]+[a-z0-9]*`)
	orderedMap.Set("ASSIGN", `=`)
	orderedMap.Set("POINT", `point`)
	orderedMap.Set("LPAREN", `\(`)
	orderedMap.Set("NUM", `[0-9]`)
	orderedMap.Set("RPAREN", `\)`)
	orderedMap.Set("SEMICOLON", `;`)
	orderedMap.Set("PERIOD", `\.`)

	compiledPattern := orderedmap.NewOrderedMap[string, *regexp.Regexp]()

	yuh := orderedmap.NewOrderedMap[int, int]()
	yuh.Set(7, 2)
	yuh.Set(3, 4)
	yuh.Set(5, 6)
	yuh.Set(1, 8)

	/*
		re := regexp.MustCompile(`^[a-z]+[a-z0-9]*`)
			id := re.FindString(line)
			if id == "test" {
				fmt.Println("TEST")
			} else {
				fmt.Println("ID", id)
			}
	*/

	for element := orderedMap.Front(); element != nil; element = element.Next() {
		compiledPattern.Set(element.Key, regexp.MustCompile(element.Value))
	}

	/*
		for item := compiledPattern.Front(); item != nil; item = item.Next() {
			fmt.Println(item.Key, item.Value)
			fmt.Printf("Type of 'item.Value': %T\n", item.Value)
		}
	*/

	file, err := os.ReadFile("test1.cpl")
	if err != nil {
		panic(err)
	}

	//var fileBuild strings.Builder
	// scanner := bufio.NewScanner(file)
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// }

	//entireFile := fileBuild.String()
	entireFile := string(file)

	tokens(entireFile)

	//defer file.Close()
}

func tokens(entireFile string) {

	print(entireFile)

	trimmedFile := strings.TrimSpace(entireFile)

	orderedMap := orderedmap.NewOrderedMap[string, string]()

	orderedMap.Set("ID", `^[a-z]+[a-z0-9]*`)
	orderedMap.Set("ASSIGN", `=`)
	orderedMap.Set("POINT", `point`)
	orderedMap.Set("LPAREN", `\(`)
	orderedMap.Set("NUM", `[0-9]`)
	orderedMap.Set("COMMA", `,`)
	orderedMap.Set("RPAREN", `\)`)
	orderedMap.Set("SEMICOLON", `;`)
	orderedMap.Set("PERIOD", `\.`)

	//compiledPattern := orderedmap.NewOrderedMap[string, *regexp.Regexp]()

	yuh := orderedmap.NewOrderedMap[int, int]()
	yuh.Set(7, 2)
	yuh.Set(3, 4)
	yuh.Set(5, 6)
	yuh.Set(1, 8)

	/*
		re := regexp.MustCompile(`^[a-z]+[a-z0-9]*`)
			id := re.FindString(line)
			if id == "test" {
				fmt.Println("TEST")
			} else {
				fmt.Println("ID", id)
			}
	*/

	for len(trimmedFile) > 0 {
		for element := orderedMap.Front(); element != nil; element = element.Next() {
			matched, _ := regexp.MatchString(element.Value, trimmedFile)
			re := regexp.MustCompile(element.Value)
			found := re.FindString(trimmedFile)
			if matched {
				fmt.Println(element.Key)
				fmt.Printf("%q\n", re.FindString(trimmedFile))
				trimmedFile = strings.Replace(trimmedFile, found, "", 1)
				//break
			}

			//compiledPattern.Set(element.Key, regexp.MustCompile(element.Value))
		}
	}

	// for item := compiledPattern.Front(); item != nil; item = item.Next() {
	// 	matched, _ := regexp.MatchString(item.Value, line)
	// }

	/*
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
			//

		// tri := regexp.MustCompile(`(,\s*\d+)`)
		// triVars := tri.FindAllStringSubmatch(line, -1)

		// tri1 := os.DevNull
		// tri2 := os.DevNull
		// tri3 := os.DevNull

		// if len(triVars) > 0

		// 	tri1 = match[1]
		// 	fmt.Println("ID", tri1)
		// 	tri2 = match[2]
		// 	fmt.Print

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

	*/
}
