package parser

import "fmt"

func Lex(input string) ([]Token, error) {
	var tokens []Token
	var current []rune

	inSingle := false
	inDouble := false
	escaped := false

	flushWord := func() {
		if len(current) > 0 {
			tokens = append(tokens, Token{
				Type:  WORD,
				Value: string(current),
			})
			current = nil
		}
	}

	for _, ch := range input {

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
				flushWord()
				continue
			case '|':
				flushWord()
				tokens = append(tokens, Token{Type: PIPE, Value: "|"})
				continue
			case '>':
				flushWord()
				tokens = append(tokens, Token{Type: REDIRECT_OUT, Value: ">"})
				continue
			case '<':
				flushWord()
				tokens = append(tokens, Token{Type: REDIRECT_IN, Value: "<"})
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

	flushWord()
	return tokens, nil
}
