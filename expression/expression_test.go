package expression

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInfixToPostfix(t *testing.T) {
	infix := "(a+b)-(c*d)"
	postfix := infixToPostfix(infix)
	assert.Equal(t, "ab+cd*-", postfix)
}

func TestPostfixToInfix(t *testing.T) {
	postfix := "ab+cd*-"
	infix := postfixToInfix(postfix)
	assert.Equal(t, "((a+b)-(c*d))", infix)
}

func TestNewExpression(t *testing.T) {
	input := "(a+b)-(c*d)"
	exp := NewExpression(input)
	assert.Equal(t, "-+ab*cd", exp.Prefix())
	assert.Equal(t, "((a+b)-(c*d))", exp.Infix())
	assert.Equal(t, "ab+cd*-", exp.Postfix())
}

func TestExpression_Evaluate(t *testing.T) {
	input := "(a+b)-(c*d)"
	exp := NewExpression(input)

	values := map[string] float64 {
		"a": 12,
		"b": 8,
		"c": 2,
		"d": -1,
	}

	assert.Equal(t, float64(22), exp.Evaluate(values))
}
