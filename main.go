package main

import (
	"os"

	"kwago.dev/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
