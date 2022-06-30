package expression

import (
	"fmt"
	"go/token"
	"strconv"
)

var DefaultOperators = map[token.Token]Operator{
	token.ADD: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			if !a.IsNumber() {
				return fmt.Errorf("Token %s must be number", a.Literal)
			}
			if !b.IsNumber() {
				return fmt.Errorf("Token %s must be number", b.Literal)
			}
			n1, isInt1 := a.Int()
			n2, isInt2 := b.Int()
			switch {
			case isInt1 && isInt2:
				stack.Push(Token{
					Token:   token.INT,
					Literal: strconv.Itoa(n2 + n1),
					Pos:     b.Pos,
				})
			default:
				stack.Push(Token{
					Token:   token.FLOAT,
					Literal: strconv.FormatFloat((b.Float() + a.Float()), 'g', 5, 64),
					Pos:     b.Pos,
				})
			}
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.SUB: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			if !a.IsNumber() {
				return fmt.Errorf("Token %s must be number", a.Literal)
			}
			if !b.IsNumber() {
				return fmt.Errorf("Token %s must be number", b.Literal)
			}
			n1, isInt1 := a.Int()
			n2, isInt2 := b.Int()
			switch {
			case isInt1 && isInt2:
				stack.Push(Token{
					Token:   token.INT,
					Literal: strconv.Itoa(n2 - n1),
					Pos:     b.Pos,
				})
			default:
				stack.Push(Token{
					Token:   token.FLOAT,
					Literal: strconv.FormatFloat((b.Float() - a.Float()), 'g', 5, 64),
					Pos:     b.Pos,
				})
			}
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.MUL: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			if !a.IsNumber() {
				return fmt.Errorf("Token %s must be number", a.Literal)
			}
			if !b.IsNumber() {
				return fmt.Errorf("Token %s must be number", b.Literal)
			}
			n1, isInt1 := a.Int()
			n2, isInt2 := b.Int()
			switch {
			case isInt1 && isInt2:
				stack.Push(Token{
					Token:   token.INT,
					Literal: strconv.Itoa(n2 * n1),
					Pos:     b.Pos,
				})
			default:
				stack.Push(Token{
					Token:   token.FLOAT,
					Literal: strconv.FormatFloat((b.Float() * a.Float()), 'g', 5, 64),
					Pos:     b.Pos,
				})
			}
			return nil
		},
		priority:    20,
		isLeftAssoc: false,
	},
	token.QUO: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			if !a.IsNumber() {
				return fmt.Errorf("Token %s must be number", a.Literal)
			}
			if !b.IsNumber() {
				return fmt.Errorf("Token %s must be number", b.Literal)
			}
			n1, isInt1 := a.Int()
			n2, isInt2 := b.Int()
			switch {
			case isInt1 && isInt2:
				stack.Push(Token{
					Token:   token.INT,
					Literal: strconv.Itoa(n2 / n1),
					Pos:     b.Pos,
				})
			default:
				stack.Push(Token{
					Token:   token.FLOAT,
					Literal: strconv.FormatFloat((b.Float() / a.Float()), 'g', 5, 64),
					Pos:     b.Pos,
				})
			}
			return nil
		},
		priority:    20,
		isLeftAssoc: false,
	},
	token.REM: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			if !a.IsNumber() {
				return fmt.Errorf("Token %s must be number", a.Literal)
			}
			if !b.IsNumber() {
				return fmt.Errorf("Token %s must be number", b.Literal)
			}
			n1, isInt1 := a.Int()
			n2, isInt2 := b.Int()
			switch {
			case isInt1 && isInt2:
				stack.Push(Token{
					Token:   token.INT,
					Literal: strconv.Itoa(n2 % n1),
					Pos:     b.Pos,
				})
			default:
				return fmt.Errorf("rem operation valid only for ints")
			}
			return nil
		},
		priority:    20,
		isLeftAssoc: false,
	},

	token.AND: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(n2 & n1),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    20,
		isLeftAssoc: false,
	},
	token.OR: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(n2 | n1),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.XOR: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(n2 ^ n1),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.SHL: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(n2 << n1),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    30,
		isLeftAssoc: false,
	},
	token.SHR: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(n2 >> n1),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    30,
		isLeftAssoc: false,
	},

	token.LAND: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			r := 0
			if n1 != 0 && n2 != 0 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})

			return nil
		},
		priority:    20,
		isLeftAssoc: false,
	},
	token.LOR: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			r := 0
			if n1 != 0 || n2 != 0 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})

			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.EQL: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			r := 0
			if a.Literal == b.Literal {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.LSS: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			r := 0
			if n2 < n1 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.GTR: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			r := 0
			if n2 > n1 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.NEQ: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			r := 0
			if a.Literal != b.Literal {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.LEQ: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			r := 0
			if n2 <= n1 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.GEQ: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			b := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			n2, isInt2 := b.Int()
			if !isInt2 {
				return fmt.Errorf("Token %s must be integer", b.Literal)
			}
			r := 0
			if n2 >= n1 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     b.Pos,
			})
			return nil
		},
		priority:    10,
		isLeftAssoc: false,
	},
	token.NOT: {
		fn: func(stack *Stack) error {
			a := stack.Pop()
			n1, isInt1 := a.Int()
			if !isInt1 {
				return fmt.Errorf("Token %s must be integer", a.Literal)
			}
			r := 0
			if n1 == 0 {
				r = 1
			}
			stack.Push(Token{
				Token:   token.INT,
				Literal: strconv.Itoa(r),
				Pos:     a.Pos,
			})

			return nil
		},
		priority:    40,
		isLeftAssoc: false,
	},
}

var DefaultFunctions = map[string]func(stack *Stack) error{
	"max": func(stack *Stack) error {
		a := stack.Pop()
		b := stack.Pop()
		n1, isInt1 := a.Int()
		if !isInt1 {
			return fmt.Errorf("Token %s must be integer", a.Literal)
		}
		n2, isInt2 := b.Int()
		if !isInt2 {
			return fmt.Errorf("Token %s must be integer", b.Literal)
		}
		r := n2
		if n2 < n1 {
			r = n1
		}
		stack.Push(Token{
			Token:   token.INT,
			Literal: strconv.Itoa(r),
			Pos:     b.Pos,
		})
		return nil
	},
	"min": func(stack *Stack) error {
		a := stack.Pop()
		b := stack.Pop()
		n1, isInt1 := a.Int()
		if !isInt1 {
			return fmt.Errorf("Token %s must be integer", a.Literal)
		}
		n2, isInt2 := b.Int()
		if !isInt2 {
			return fmt.Errorf("Token %s must be integer", b.Literal)
		}
		r := n2
		if n2 > n1 {
			r = n1
		}
		stack.Push(Token{
			Token:   token.INT,
			Literal: strconv.Itoa(r),
			Pos:     b.Pos,
		})
		return nil
	},
}
