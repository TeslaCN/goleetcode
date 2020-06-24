package p20_1

func IsValid(s string) bool {
	if s == "" {
		return true
	}
	st := &stack{
		a: make([]uint8, len(s)),
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			st.push(s[i])
			continue
		}
		if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			if st.empty() {
				return false
			}
			pop := st.pop()
			switch s[i] {
			case ')':
				if pop != '(' {
					return false
				}
			case ']':
				if pop != '[' {
					return false
				}
			case '}':
				if pop != '{' {
					return false
				}
			}
		}
	}
	return st.empty()
}

type stack struct {
	a   []uint8
	cur int
}

func (s *stack) empty() bool {
	return s.cur < 1
}

func (s *stack) push(e uint8) {
	s.a[s.cur] = e
	s.cur++
}

func (s *stack) pop() uint8 {
	if s.cur >= 0 {
		s.cur--
		return s.a[s.cur]
	}
	panic("stack empty")
}
