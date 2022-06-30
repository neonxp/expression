package expression

import (
	"go/scanner"
	"go/token"
)

type Evaluator struct {
	operators map[token.Token]Operator
	functions map[string]func(stack *Stack) error
}

func New() *Evaluator {
	return &Evaluator{
		operators: DefaultOperators,
		functions: DefaultFunctions,
	}
}

func (e *Evaluator) Eval(expression string) (any, error) {
	s := scanner.Scanner{}
	fset := token.NewFileSet()
	file := fset.AddFile("", fset.Base(), len(expression))
	s.Init(file, []byte(expression), nil, scanner.ScanComments)
	tokens := make(chan Token)
	go func() {
		for {
			pos, tok, lit := s.Scan()
			if tok == token.SEMICOLON {
				continue
			}
			if tok == token.EOF {
				break
			}
			tokens <- Token{
				Token:   tok,
				Literal: lit,
				Pos:     int(pos),
			}
		}
		close(tokens)
	}()
	rpnTokens := e.ToPRN(tokens)
	return e.execute(rpnTokens)
}

type Operator struct {
	fn          func(stack *Stack) error
	priority    int
	isLeftAssoc bool
}
