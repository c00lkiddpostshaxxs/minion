package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: minion <filename>")
		os.Exit(1)
	}

	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	lexer := NewLexer(string(code))
	var tokens []Token
	for {
		tok := lexer.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == EOF {
			break
		}
	}

	parser := NewParser(tokens)
	ast := parser.Parse()

	interpreter := NewInterpreter()
	interpreter.Interpret(ast)
}
