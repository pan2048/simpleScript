package parser

import (
	"github.com/stretchr/testify/assert"
	"simple/ast"
	"simple/lexer"
	"testing"
)

func TestCallExpressionParsing(t *testing.T) {
	input := "add(1, 2 * 3, 4 + 5);"
	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParserErrors(t, p)
	assert.Equal(t, 1, len(program.Statements))
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok)
	exp, ok := stmt.Expression.(*ast.CallExpression)
	assert.True(t, ok)
	if !testIdentifier(t, exp.Function, "add") {
		return
	}
	assert.Equal(t, 3, len(exp.Arguments))
	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}
