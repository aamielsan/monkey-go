package main

import (
	"fmt"

	"kwago.dev/monkey/lexer"
)

func main() {
	token := lexer.New(`
        let five = 5;
    `).NextToken()
	fmt.Printf("token is %+v", token)
}
