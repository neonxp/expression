package expression

import (
	"fmt"
	"go/token"
	"strings"
)

func (e *Evaluator) execute(tokens chan Token) (any, error) {
	stack := Stack{}
	for tok := range tokens {
		switch {
		case tok.IsNumber():
			stack.Push(tok)
		case tok.IsOperator():
			op := e.operators[tok.Token]
			if err := op.fn(&stack); err != nil {
				return nil, err
			}
		case tok.IsFunc():
			fn, fnEsist := e.functions[strings.ToLower(tok.Literal)]
			if !fnEsist {
				return nil, fmt.Errorf("unknown function %s at %d", tok.Literal, tok.Pos)
			}
			if err := fn(&stack); err != nil {
				return nil, err
			}
		case tok.IsError():
			return nil, fmt.Errorf("Error at token %d: %w", tok.Pos, tok.Error())
		}
	}
	if len(stack) != 1 {
		return nil, fmt.Errorf("Expected exact one returning value, go %+v", stack)
	}
	result := stack.Pop()
	switch result.Token {
	case token.INT:
		n, _ := result.Int()
		return n, nil
	case token.FLOAT:
		return result.Float(), nil
	default:
		return result.Literal, nil
	}
}
