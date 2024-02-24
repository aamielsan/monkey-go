package repl

import (
	"bufio"
	"fmt"
	"io"

	"kwago.dev/monkey/lexer"
	"kwago.dev/monkey/token"
)

func Start(in io.Reader, out io.Writer) {
	fmt.Fprintln(out, "Welcome to Monkey")
	fmt.Fprintln(out, "To exit, press Ctrl+C")
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, ">> ")
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexer := lexer.New(line)

		for tok := lexer.NextToken(); tok.Type != token.Eof.Type; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
