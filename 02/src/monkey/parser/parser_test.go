package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`
	/*input := `
	let x 5;
	let = 10;
	let 838383;
	`*/

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	//如果program 等於nil
	if program == nil {
		t.Fatalf("ParseProgram() return nil")
	}

	//如果statement 長度不為3
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	//測試的identifier
	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	//一樣沒有3個就失敗
	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contains 3 statement. got=%d",
			len(program.Statements))
	}

	for _, stmt := range program.Statements {
		//檢查是否為returnStatement
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		//如果不為return
		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		//如果他的tokenLiteral 不為return
		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got=%q",
				returnStmt.TokenLiteral())
		}

	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	//如果不為let
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	//檢查目前的Statement 是不是LetStatement
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
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

// 檢查錯誤
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
