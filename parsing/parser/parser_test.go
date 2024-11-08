package parser

import (
	"fmt"
	"log"
	"testing"

	"github.com/ajtroup1/interpreters/parsing/ast"
	"github.com/ajtroup1/interpreters/parsing/lexer"
	"github.com/ajtroup1/interpreters/util"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 838383;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	if program == nil {
		t.Fatalf(util.RedText("ParseProgram() returned nil"))
	}

	expectedStatements := 3
	if len(program.Statements) != expectedStatements {
		t.Fatalf(util.RedText(fmt.Sprintf("program.Statements does not contain %d statements. got=%d",
			expectedStatements, len(program.Statements))))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	numPassed := 0
	for i, tt := range tests {
		stmt := program.Statements[i]
		if testLetStatement(t, stmt, tt.expectedIdentifier) {
			numPassed++
		}
	}

	totalTests := len(tests)
	if numPassed == totalTests {
		log.Println(util.GreenText(fmt.Sprintf("%d / %d PARSER TESTS PASSED", numPassed, totalTests)))
	} else {
		log.Println(util.RedText(fmt.Sprintf("%d / %d PARSER TESTS PASSED", numPassed, totalTests)))
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf(util.RedText(fmt.Sprintf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())))
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf(util.RedText(fmt.Sprintf("s not *ast.LetStatement. got=%T", s)))
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf(util.RedText(fmt.Sprintf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)))
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf(util.RedText(fmt.Sprintf("s.Name not '%s'. got=%s", name, letStmt.Name)))
		return false
	}
	return true
}
