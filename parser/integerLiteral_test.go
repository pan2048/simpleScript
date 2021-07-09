package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"simple/ast"
	"simple/lexer"
	"testing"
)

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	assert.Equal(t, 1, len(program.Statements))
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	literal, ok := stmt.Expression.(*ast.IntegerLiteral)
	assert.True(t, ok)
	assert.Equal(t, int64(5), literal.Value)
	assert.Equal(t, "5", literal.TokenLiteral())
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value interface{}) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	assert.True(t, ok)
	assert.Equal(t, value, integ.Value)
	assert.Equal(t, fmt.Sprintf("%d", value), integ.TokenLiteral())
	return true
}
