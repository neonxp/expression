package expression

type Stack []Token

func (s *Stack) Push(item Token) {
	*s = append(*s, item)
}

func (s *Stack) Pop() (item Token) {
	if len(*s) == 0 {
		return
	}

	*s, item = (*s)[:len(*s)-1], (*s)[len(*s)-1]
	return item
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Head() (item *Token) {
	if s.Empty() {
		return nil
	}
	return &((*s)[len(*s)-1])
}
