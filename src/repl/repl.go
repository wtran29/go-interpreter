// REPL - Read input, sneds to interpreter to Eval, Print result, Loop to start again
package repl

import (
	"bufio"
	"fmt"

	"io"

	"github.com/wtran29/go-interpreter/src/lexer"
	"github.com/wtran29/go-interpreter/src/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
