package parser

import (
	"github.com/stretchr/testify/assert"
	"simple/ast"
	"simple/lexer"
	"testing"
)

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
	assert.Equal(t, 3, len(program.Statements))
	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		assert.True(t, ok)
		assert.Equal(t, "return", returnStmt.TokenLiteral())
	}
}
