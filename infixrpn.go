package expression

import (
	"fmt"
	"go/token"
)

func (e *Evaluator) ToPRN(in <-chan Token) chan Token {
	out := make(chan Token)
	stack := &Stack{}

	go func() {
		defer func() {
			for !stack.Empty() {
				tok := stack.Pop()
				if tok.LP() {
					out <- Token{
						Token:   token.ILLEGAL,
						Literal: "no closing parenthesis",
						Pos:     tok.Pos,
					}
				} else {
					out <- tok
				}
			}
			close(out)
		}()
		for tok := range in {
			switch {
			case tok.Token == token.ILLEGAL:
				return
			case tok.IsNumber():
				out <- tok
			case tok.IsFunc():
				stack.Push(tok)
			case tok.IsSeparator():
				for {
					if stack.Empty() {
						out <- Token{
							Token:   token.ILLEGAL,
							Literal: "no opening parenthesis",
							Pos:     tok.Pos,
						}
						return
					}
					if stack.Head().LP() {
						break
					}
					out <- tok
				}
			case tok.IsOperator():
				op1 := e.operators[tok.Token]
				for {
					if stack.Empty() {
						break
					}
					if stack.Head().IsOperator() {
						op2, hasOp := e.operators[stack.Head().Token]
						if !hasOp {
							out <- Token{
								Token:   token.ILLEGAL,
								Literal: fmt.Sprintf("unknown operator: %s", stack.Head().Literal),
								Pos:     tok.Pos,
							}
							return
						}
						if op2.priority > op1.priority {
							out <- stack.Pop()
							continue
						} else {
							break
						}
					} else {
						break
					}
				}
				stack.Push(tok)
			case tok.LP():
				stack.Push(tok)
			case tok.RP():
				for {
					if stack.Empty() {
						out <- Token{
							Token:   token.ILLEGAL,
							Literal: "no opening parenthesis",
							Pos:     tok.Pos,
						}
						return
					}
					if stack.Head().LP() {
						break
					}
					out <- tok
				}
				stack.Pop()
				if stack.Head().IsFunc() {
					out <- stack.Pop()
				}
			}
		}
	}()

	return out
}
