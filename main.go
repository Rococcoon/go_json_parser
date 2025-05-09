package main

import (
	"fmt"
	"rococcoon/go_json_test/interpreter"
	"rococcoon/go_json_test/lexer"
	"rococcoon/go_json_test/parser"
)

func main() {
	input := ` 
  {
    "testCase": [123, "abc", true, null, {"nested": "yo"}]
  }
  `
	l := lexer.NewLexer(input)
	tokens := l.TokenizeInput()

	p := parser.NewParser(tokens)
	parsedJson := p.ParseRoot()
	goValue := interpreter.ConvertToGoValue(parsedJson.Value)
	data, _ := goValue.(map[string]interface{})
	testCase, ok := data["testCase"].([]interface{})
	if ok {
		fmt.Printf("%v\n", testCase[4])
	}
}
