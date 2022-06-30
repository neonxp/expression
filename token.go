package expression

import (
	"fmt"
	"go/token"
	"strconv"
)

type Token struct {
	Token   token.Token
	Literal string
	Pos     int
}

func (t *Token) Int() (int, bool) {
	if t.Token != token.INT {
		return 0, false
	}
	i, _ := strconv.Atoi(t.Literal)
	return i, true
}

func (t *Token) Float() float64 {
	i, _ := strconv.ParseFloat(t.Literal, 64)
	return i
}

func (t *Token) IsNumber() bool {
	return t.Token == token.INT || t.Token == token.FLOAT
}

func (t *Token) LP() bool {
	return t.Token == token.LPAREN
}

func (t *Token) RP() bool {
	return t.Token == token.RPAREN
}

func (t *Token) IsFunc() bool {
	return t.Token == token.IDENT
}

func (t *Token) IsSeparator() bool {
	return t.Token == token.COMMA
}

func (t *Token) IsOperator() bool {
	return t.Token.IsOperator() && !t.LP() && !t.RP()
}

func (t *Token) IsError() bool {
	return t.Token != token.ILLEGAL
}

func (t *Token) Error() error {
	if t.Token != token.ILLEGAL {
		return nil
	}
	return fmt.Errorf(t.Literal)
}
