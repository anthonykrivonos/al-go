package expression

import (
	"github.com/anthonykrivonos/al-go/list"
	"github.com/anthonykrivonos/al-go/set"
	"github.com/anthonykrivonos/al-go/stack"
	"github.com/anthonykrivonos/al-go/tree"
	"github.com/anthonykrivonos/al-go/utils"
	"math"
	"strings"
)

type Expression interface {
	Prefix() string
	Infix() string
	Postfix() string
	Evaluate(values map[string]float64) float64
}

type expression struct {
	tree tree.BinaryTree
	prefix string
	infix string
	postfix string
}

func operators() set.Set {
	o := set.NewSet()
	o.Insert('^')
	o.Insert('*')
	o.Insert('/')
	o.Insert('+')
	o.Insert('-')
	return o
}

func brackets() set.Set {
	o := set.NewSet()
	o.Insert('(')
	o.Insert('[')
	o.Insert('{')
	o.Insert(')')
	o.Insert(']')
	o.Insert('}')
	return o
}

func oneSidedBrackets(open bool) set.Set {
	o := set.NewSet()
	if open {
		o.Insert('(')
		o.Insert('[')
		o.Insert('{')
	} else {
		o.Insert(')')
		o.Insert(']')
		o.Insert('}')
	}
	return o
}

func infixToPostfix(infix string) string {
	openBrackets := oneSidedBrackets(true)
	brackets := brackets()
	operators := operators()

	var operatorStack stack.Stack = list.NewArrayList()

	postfix := strings.Builder{}

	for _, c := range []rune(infix) {
		if operators.Has(c) || openBrackets.Has(c) {
			// Open bracket or operator
			operatorStack.Push(c)
		} else if !brackets.Has(c) && !operators.Has(c) {
			// Operand
			postfix.WriteRune(c)
		} else {
			// Closing bracket
			for !openBrackets.Has(operatorStack.Peek()) {
				nextC := operatorStack.Pop().(rune)
				postfix.WriteRune(nextC)
			}
			operatorStack.Pop()
		}
	}

	// Add remaining operands to the postfix expression
	for operatorStack.Length() > 0 {
		postfix.WriteRune(operatorStack.Pop().(rune))
	}

	return postfix.String()
}

func postfixToInfix(postfix string) string {
	brackets := brackets()
	operators := operators()

	var operandStack stack.Stack = list.NewArrayList()

	for _, c := range []rune(postfix) {
		if !brackets.Has(c) && !operators.Has(c) {
			// Operand
			operandStack.Push(string(c))
		} else {
			// Operator
			rightOperand := operandStack.Pop().(string)
			leftOperand := operandStack.Pop().(string)

			// Build expression
			expression := strings.Builder{}
			expression.WriteRune('(')
			expression.WriteString(leftOperand)
			expression.WriteRune(c)
			expression.WriteString(rightOperand)
			expression.WriteRune(')')

			operandStack.Push(expression.String())
		}
	}
	return operandStack.Pop().(string)
}

func postfixToTree(postfix string) tree.BinaryTree {
	brackets := brackets()
	operators := operators()

	var operandStack stack.Stack = list.NewArrayList()

	for _, c := range []rune(postfix) {
		if !brackets.Has(c) && !operators.Has(c) {
			// Operand
			operandStack.Push(tree.NewNode(string(c), nil, nil))
		} else {
			// Operator
			rightOperand := operandStack.Pop().(tree.BinaryTreeNode)
			leftOperand := operandStack.Pop().(tree.BinaryTreeNode)

			operandStack.Push(tree.NewNode(string(c), leftOperand, rightOperand))
		}
	}

	return tree.NewBinaryTree(operandStack.Pop().(tree.BinaryTreeNode))
}

func NewExpression(infix string) Expression {
	e := &expression{}
	e.postfix = infixToPostfix(infix)
	e.infix = postfixToInfix(e.postfix)
	e.tree = postfixToTree(e.postfix)

	prefix := &strings.Builder{}
	e.prefixHelper(e.tree.Root(), prefix)
	e.prefix = prefix.String()

	return e
}

func (e *expression) prefixHelper(n tree.BinaryTreeNode, builder *strings.Builder) {
	if !utils.IsNilInterface(n) {
		builder.WriteString(n.Value().(string))
		e.prefixHelper(n.Left(), builder)
		e.prefixHelper(n.Right(), builder)
	}
}

func (e *expression) Prefix() string {
	return e.prefix
}

func (e *expression) Infix() string {
	return e.infix
}

func (e *expression) Postfix() string {
	return e.postfix
}

func (e *expression) Evaluate(values map[string]float64) float64 {
	return e.evaluate(e.tree.Root(), values)
}

func (e *expression) evaluate(n tree.BinaryTreeNode, values map[string]float64) float64 {
	if utils.IsNilInterface(n) {
		return 0
	}

	if utils.IsNilInterface(n.Left()) && utils.IsNilInterface(n.Right()) {
		// We have a value
		return values[n.Value().(string)]
	}

	// Evaluate left subtree
	left := e.evaluate(n.Left(), values)

	// Evaluate right subtree
	right := e.evaluate(n.Right(), values)

	// Apply operator
	switch n.Value() {
		case "^":
			return math.Pow(left, right)
		case "/":
			return left / right
		case "+":
			return left + right
		case "-":
			return left - right
		default:
			return left * right
	}
}

var _ Expression = &expression{}
