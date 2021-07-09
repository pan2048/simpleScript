package parser

import (
	"github.com/stretchr/testify/assert"
	"simple/ast"
	"simple/lexer"
	"testing"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}
	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)
		assert.Equal(t, 1, len(program.Statements))
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		assert.True(t, ok)
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		assert.True(t, ok)
		if !testLiteralExpression(t, exp.Left, tt.leftValue) {
			return
		}
		assert.Equal(t, tt.operator, exp.Operator)
		if !testLiteralExpression(t, exp.Right, tt.rightValue) {
			return
		}
	}
}
