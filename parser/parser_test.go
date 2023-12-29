package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLeftStatement(t *testing.T) {
	input := `
let x 5;
let = 10;
let 838 383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
	if len(program.Statements) == 1 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}
	testCases := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
	}

	for i, tc := range testCases {
		stmt := program.Statements[i]
		if !testLeftStatement(t, stmt, tc.expectedIdentifier) {
			return
		}
	}
}

func testLeftStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	leftStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LeftStatement. got=%T", s)
		return false
	}

	if leftStmt.Name.Value != name {
		t.Errorf("leftStmt.Name.Value not '%s'. got=%s", name, leftStmt.Name.Value)
		return false
	}

	if leftStmt.Name.TokenLiteral() != name {
		t.Errorf("leftStmt.Name.TokenLiteral() not '%s'. got=%s", name, leftStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
