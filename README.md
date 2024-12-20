# Go Programming: 4Point Grammar Lexical and Syntax Analysis


## Motivation
Go is a programming language designed at Google. From the Go [website](https://golang.org/ref/spec):
> Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies.
> The grammar is compact and simple to parse, allowing for easy analysis by automatic tools such as integrated development environments.

This project consists in the development of the front end of a compiler. By programming the Lexical Analyzer (Scanner) and Syntax Analyzer (Parser) for the 4Point grammar you will gain further understanding of the lexical analysis and the production of tokens needed for the Syntax Analyzer (Parser), and how to consume those tokens by the Parser to verify that the syntax is correct.


## Description
A program in Go that takes a program written in 4Point, and outputs:
1. If the code has lexical or syntax errors, the error that was found. Use panic version of error handling (once an error is found report the error and stop the process).
1. If the code is OK, depending on a command line flag the program will produce:
   1.	If the flag is `-s` the program will output function calls in Scheme that is going to be called by a program in Scheme that will calculate properties of those four points.
   1. If the flag is `-p` the program will output a series of queries about those four points.

The program should run like this:
```
prompt>go run . input.txt -s
; Processing Input File input.txt
; Lexical and Syntax analysis passed
; Generating Scheme Code
(process-triangle (make-point 2 3) (make-point 1 4) (make-point 3 4))
prompt>
```

If there is a problem with the program invocation, your program should report it. For example:
```
prompt>go run . file.html
Missing parameter, usage:
go run . filename -flag
flag can be p for prolog generation
flag can be s for scheme generation
```

## Grammar

```
START      --> STMT_LIST
STMT_LIST  --> STMT. |
               STMT; STMT_LIST
STMT       --> POINT_DEF |
               TEST
POINT_DEF  --> ID = point(NUM, NUM)
TEST       --> test(OPTION, POINT_LIST)
ID         --> LETTER+
NUM        --> DIGIT+
OPTION     --> triangle |
               square
POINT_LIST --> ID |
               ID, POINT_LIST
LETTER     --> a | b | c | d | e | f | g | ... | z
DIGIT      --> 0 | 1 | 2 | 3 | 4 | 5 | 6 | ... | 9

```

The tokens of this grammar are:

Token | Lexeme
------ | ------
`POINT` | `point`
`ID` | `identifier`
`NUM` | `234`
`SEMICOLON` | `;`
`COMMA` | `,`
`PERIOD` | `.`
`LPAREN` | `(`
`RPAREN` | `)`
`ASSIGN` | `=`
`TRIANGLE` | `triangle`
`SQUARE` | `square`
`TEST` | `test`

Given the following program written in this language:
```
a = point(2, 3);
b = point(1, 1);
c = point(1, 3);
d = point(0, 0);
test(square, a, b, c, d);
test(triangle, a, b, c).
```
The tokens that it would generate are:
```
ID  a
ASSIGN
POINT
LPAREN
NUM 2
COMMA
NUM 3
RPAREN
SEMICOLON
ID  b
ASSIGN
POINT
LPAREN
NUM 1
COMMA
NUM 1
RPAREN
SEMICOLON
ID  c
ASSIGN
POINT
LPAREN
NUM 1
COMMA
NUM 3
RPAREN
SEMICOLON
ID  d
ASSIGN
POINT
LPAREN
NUM 0
COMMA
NUM 0
RPAREN
SEMICOLON
TEST
LPAREN
SQUARE
COMMA
ID a
COMMA
ID b
COMMA
ID c
COMMA
ID d
RPAREN
SEMICOLON
TEST
LPAREN
TRIANGLE
COMMA
ID a
COMMA
ID b
COMMA
ID c
RPAREN
PERIOD
```

Notice that the ID and NUM tokens have their lexeme associated. Also notice that in the language the elements do not need to be separated by space. 
The tokens are generated within the program, and will not be shown as an output. This is to demonstrate what is happening in the backend.

---

### How to run the program

`go run input.txt`

The following examples assume that `input.txt` contains the following code:
```
a = point(2, 3);
b = point(1, 1);
c = point(1, 3);
d = point(0, 0);
test(square, a, b, c, d);
test(triangle, a, b, c).
```

### Scheme Output
The following Scheme output will be generated when you add the `-s` flag at the end of the command:
```
prompt> go run . input.txt -s
; processing input file input.txt
; Lexical and Syntax analysis passed
; Generating Scheme Code
(process-square (make-point 2 3) (make-point 1 1) (make-point 1 3) (make-point 0 0))
(process-triangle (make-point 2 3) (make-point 1 1) (make-point 1 3))
```

### Prolog Output
To generate prolog output you will add the `-p` flag at the end of the command:
```
prompt> go run .  input.txt -p
/* processing input file input.txt
   Lexical and Syntax analysis passed
   Generating Prolog Code */

 /* Processing test(square, a, b, c, d) */
 query(square(point2d(2, 3), point2d(1, 1), point2d(1, 3), point2d(0, 0))).

 /* Processing test(triangle, a, b, c) */
 query(line(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(triangle(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(vertical(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(horizontal(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(equilateral(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(isosceles(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(right(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(scalene(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(acute(point2d(2,3), point2d(1,1), point2d(1, 3))).
 query(obtuse(point2d(2,3), point2d(1,1), point2d(1, 3))).
 
 /* Query Processing */
 writeln(T) :- write(T), nl.
 main:- forall(query(Q), Q-> (writeln(‘yes’)) ; (writeln(‘no’))),
       halt.

```


