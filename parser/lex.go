package parser

import "fmt"

func Lexcmd(input string) ([]Token, error) {
	var tokens []Token
	var current []rune

	inSingle := false
	inDouble := false
	escaped := false

	flush := func() {
		if len(current) > 0 {
			tokens = append(tokens, Token{Type: WORD, Value: string(current)})
			current = nil
		}
	}

	runes := []rune(input)

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if escaped {
			current = append(current, ch)
			escaped = false
			continue
		}

		if ch == '\\' && !inSingle {
			escaped = true
			continue
		}

		if ch == '\'' && !inDouble {
			inSingle = !inSingle
			continue
		}

		if ch == '"' && !inSingle {
			inDouble = !inDouble
			continue
		}

		if !inSingle && !inDouble {
			switch ch {
			case ' ':
				flush()
				continue
			case '|':
				flush()
				tokens = append(tokens, Token{Type: PIPE, Value: "|"})
				continue
			case '<':
				flush()
				tokens = append(tokens, Token{Type: REDIRECT_IN, Value: "<"})
				continue
			case '>':
				flush()
				if i+1 < len(runes) && runes[i+1] == '>' {
					tokens = append(tokens, Token{Type: REDIRECT_APPEND, Value: ">>"})
					i++
				} else {
					tokens = append(tokens, Token{Type: REDIRECT_OUT, Value: ">"})
				}
				continue
			}
		}

		current = append(current, ch)
	}

	if escaped {
		return nil, fmt.Errorf("unfinished escape")
	}
	if inSingle || inDouble {
		return nil, fmt.Errorf("unclosed quote")
	}

	flush()
	return tokens, nil
}
