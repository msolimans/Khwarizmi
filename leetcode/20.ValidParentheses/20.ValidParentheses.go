package ValidParantheses

import "container/list"
func isValid(s string) bool {
	
	l := list.New()

	for _, c := range s { 
		if c == '{' || c == '[' || c == '(' {
			l.PushBack(c)
			continue
		}
		back := l.Back()
		if back == nil {
			return false
		}
		l.Remove(back)

		switch c {
		case '}':
			if back.Value != '{' {
				return false
			}
		case ']':
			if back.Value != '[' {
				return false 
			}
		case ')':
			if back.Value != '('{
				return false
			}
		}

	}

	return true
}

