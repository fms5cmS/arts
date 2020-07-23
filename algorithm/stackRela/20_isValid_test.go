package stackRela

func isValid(s string) bool {
	mem := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make([]rune, 0)
	top := -1
	for _, v := range s {
		switch v {
		case '[', '(', '{':
			stack = append(stack, v)
			top++
		case ']', ')', '}':
			if top < 0 {
				return false
			}
			if mem[v] != stack[top] {
				return false
			}
			stack = stack[:top]
			top--
		}
	}
	return top == -1
}
