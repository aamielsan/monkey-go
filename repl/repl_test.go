package repl

import (
	"bytes"
	"strings"
	"testing"
)

func TestStart(t *testing.T) {
	given := `let x = 5;`

	expected := `Welcome to Monkey
To exit, press Ctrl+C
>> {Type:LET Literal:let}
{Type:IDENT Literal:x}
{Type:ASSIGN Literal:=}
{Type:INT Literal:5}
{Type:SEMICOLON Literal:;}
>> `

	var output bytes.Buffer

	Start(strings.NewReader(given), &output)

	actual := output.String()
	if actual != expected {
		t.Fatalf("Print output mismatch. \nexpected=%q \nactual=%q", expected, actual)
	}
}
