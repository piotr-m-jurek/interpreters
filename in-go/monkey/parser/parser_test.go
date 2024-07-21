package parser

import (
	"in-go/monkey/ast"
	"in-go/monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;

	let foobar = 838384;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) < 3 {
		t.Fatalf("Program did not return 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testLetStatements(t, stmt, tt.expectedIdentifier) {
			return
		}
	}

}

func TestReturnStatements(t *testing.T) {
	input := `
	return 1337;
	return 69;
	return 42

	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram returned nil")
	}

	if len(program.Statements) < 3 {
		t.Fatalf("Program did not return 3 return statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedReturnValue string
	}{
		{"1337"},
		{"69"},
		{"42"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if !testReturnStatements(t, stmt, tt.expectedReturnValue) {
			return
		}
	}

}

func checkParserErrors(t *testing.T, p *Parser) {

	errors := p.errors

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)

	}
	t.FailNow()
}

func testLetStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {

		t.Fatalf("Not Let statement, got=%s", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'. got=%s",
			name, letStmt.Name.TokenLiteral())
		return false
	}
	return true
}

func testReturnStatements(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "return" {
		t.Fatalf("Not return statement, got=%s", s.TokenLiteral())
	}

	returnStmt, ok := s.(*ast.ReturnStatement)

	if !ok {
		t.Errorf("s not *ast.ReturnStatement. got=%T", s)
		return false
	}

	if returnStmt.ReturnValue.TokenLiteral() != name {
		t.Errorf("returnStmt.ReturnValue.TokenLiteral not '%s'. got=%s",
			name,
			returnStmt.ReturnValue.TokenLiteral())
		return false
	}
	return true
}
