// REPL - Read input, sends to interpreter to Eval, Print result, Loop to start again
package repl

import (
	"bufio"
	"fmt"

	"io"

	"github.com/wtran29/go-interpreter/src/compiler"
	"github.com/wtran29/go-interpreter/src/lexer"
	"github.com/wtran29/go-interpreter/src/parser"
	"github.com/wtran29/go-interpreter/src/vm"
)

const PROMPT = ">> "

const GOPHER_FACE = `         ,_---~~~~~----._         
  _,,_,*^____       ____''*g*\"*, 
 / __/  /'    ^.   /     \ ^@q   f 
[  @f  | @))   |  | @))   l  0 _/  
 \ /    \~____/ __ \_____/    \   
  |           _l__l_           I   
  }          [______]          I  
  ]            | | |           |  
  ]             ~ ~            |  
  |                            |   
   |                           |   
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	// env := object.NewEnvironment()

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		// evaluated := evaluator.Eval(program, env)
		// if evaluated != nil {
		// 	io.WriteString(out, evaluated.Inspect())
		// 	io.WriteString(out, "\n")
		// }

		comp := compiler.New()
		err := comp.Compile(program)
		if err != nil {
			fmt.Fprintf(out, "Uh oh! Compilation failed:\n %s\n", err)
			continue
		}

		machine := vm.New(comp.Bytecode())
		err = machine.Run()
		if err != nil {
			fmt.Fprintf(out, "Uh oh! Excuting bytecode failed:\n %s\n", err)
			continue
		}
		stackTop := machine.StackTop()
		io.WriteString(out, stackTop.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, GOPHER_FACE)
	io.WriteString(out, "Oops! We ran into some funckey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
