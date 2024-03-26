// abstract syntax tree package
package ast

import (
	"bytes"

	"github.com/wtran29/go-interpreter/src/token"
)

// Node is the interface that all nodes in the AST implement.
type Node interface {
	TokenLiteral() string // Returns the literal value of the token associated with the node
	String() string       // Returns a string representation of the node
}

// Statement represents a statement in the program.
type Statement interface {
	Node
	statementNode() // Placeholder method to distinguish statement nodes from others
}

// Expression represents an expression in the program.
type Expression interface {
	Node
	expressionNode() // Placeholder method to distinguish expression nodes from others
}

// Program represents the entire program as a series of statements.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the token literal of the first statement in the program.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// String returns a string representation of the entire program by concatenating
// the string representations of each statement.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement represents a 'let' statement.
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // Identifier for the variable being declared
	Value Expression  // Expression representing the value being assigned
}

func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier represents an identifier (variable name).
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string      // Value of the identifier
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement represents a 'return' statement.
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ExpressionStatement represents an expression used as a statement.
type ExpressionStatement struct {
	Token      token.Token // the first token of expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// IntegerLiteral represents an integer literal.
type IntegerLiteral struct {
	Token token.Token // The integer token
	Value int64       // Value of the integer
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression represents a prefix expression (e.g. !x, -x).
type PrefixExpression struct {
	Token    token.Token // prefix token, e.g. ! or -
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// InfixExpression represents an infix expression (e.g. x + y, x < y).
type InfixExpression struct {
	Token    token.Token // The operator token, e.g. +
	Left     Expression
	Operator string
	Right    Expression
}

func (oe *InfixExpression) expressionNode() {}
func (oe *InfixExpression) TokenLiteral() string {
	return oe.Token.Literal
}
func (oe *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(oe.Left.String())
	out.WriteString(" " + oe.Operator + " ")
	out.WriteString(oe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Boolean represents booleans
type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

// IfExpression represents the if statements
type IfExpression struct {
	Token       token.Token // if token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

// BlockStatement represents the block of code separated by brackets
type BlockStatement struct {
	Token      token.Token // the { } token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}
