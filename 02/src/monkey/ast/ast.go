package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// statement 繼承node
type Statement interface {
	Node            // embedding，等於「繼承」這個 interface 的方法簽名
	statementNode() // 空函式，純粹拿來當「標記」，本身不做事
}

// expression 繼承 node
type Expression interface {
	Node
	expressionNode() // 一樣是空標記
}

type Program struct {
	Statements []Statement // 一串 Statement，本身就是一堆 interface
}

func (p *Program) TokenLiteral() string {
	//如果statements裡面不為空
	if len(p.Statements) > 0 {
		//回傳第一個statement的token literal
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// 把Statement 裡面的String做concat
func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// LetStatement class代表let語法
type LetStatement struct {
	Token token.Token // the token.LET token
	Name  *Identifier // 變數名: x
	Value Expression  // 值: 5 + 3 這個 expression（interface 型別！）
}

// 實作statement
func (ls *LetStatement) statementNode()       {} // 空的，只是打卡「我是 Statement」
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

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

// Identifier 代表變數名稱
type Identifier struct {
	Token token.Token // the token.IDENT token
	Value string
}

// 實作expression
func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// 用於完成 Statement 接口
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {} //佔位子
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// 用於完成 Statement 接口
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (es *ExpressionStatement) statementNode()       {} //佔位子
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
