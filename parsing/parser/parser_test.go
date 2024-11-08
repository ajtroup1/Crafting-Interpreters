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
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf(util.RedText("ParseProgram() returned nil"))
	}

	expectedStatements := 3
	if len(program.Statements) != expectedStatements {
		t.Fatalf(util.RedText(fmt.Sprintf("Expected %d statements, got %d",
			expectedStatements, len(program.Statements))))
	}

	tests := []string{"x", "y", "foobar"}
	passedTests := 0

	for i, expectedIdentifier := range tests {
		if testLetStatement(t, program.Statements[i], expectedIdentifier) {
			passedTests++
		}
	}

	totalTests := len(tests)
	logResult := util.GreenText
	if passedTests != totalTests {
		logResult = util.RedText
	}
	log.Println(logResult(fmt.Sprintf("%d / %d PARSER TESTS PASSED", passedTests, totalTests)))
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 993322;
	`
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("stmt not *ast.returnStatement. got=%T", stmt)
			continue
		}
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q",
				returnStmt.TokenLiteral())
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	letStmt, ok := s.(*ast.LetStatement)
	if !ok || letStmt.Name.Value != name || letStmt.Name.TokenLiteral() != name {
		t.Errorf(util.RedText(fmt.Sprintf("Incorrect let statement: got=%T, value=%s, literal=%s",
			s, letStmt.Name.Value, letStmt.Name.TokenLiteral())))
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf(util.RedText(fmt.Sprintf("parser has %d errors", len(errors))))
	for _, msg := range errors {
		t.Errorf(util.RedText(fmt.Sprintf("parser error: %q", msg)))
	}
	t.FailNow()
}
