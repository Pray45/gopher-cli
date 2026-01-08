package parser

import "fmt"

func Tokenize(input string) ([]string, error) {
	
	var tokens []string
	var current []rune

	inSingle := false
	inDouble := false
	escaped := false

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

		if ch == ' ' && !inSingle && !inDouble {
			if len(current) > 0 {
				tokens = append(tokens, string(current))
				current = nil
			}
			continue
		}

		current = append(current, ch)
	}

	if escaped {
		return nil, fmt.Errorf("unfinished escape")
	}

	if inSingle || inDouble {
		return nil, fmt.Errorf("unclosed quote")
	}

	if len(current) > 0 {
		tokens = append(tokens, string(current))
	}

	return tokens, nil
}
