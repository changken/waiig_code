package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	//創建一個scanner，從in讀取輸入
	scanner := bufio.NewScanner(in)

	for {
		//在out輸出PROMPT
		fmt.Printf(PROMPT)
		//如果scanner掃描不到下一個字元，則回傳
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		//讀取scanner的文字
		line := scanner.Text()
		l := lexer.New(line)

		//使用NextToken()取得下一個token，直到遇到EOF
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
