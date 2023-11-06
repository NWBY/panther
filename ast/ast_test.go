package ast

import (
	"panther/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&VolatileStatement{
				Token: token.Token{Type: token.VOLATILE, Literal: "vol"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "vol myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
