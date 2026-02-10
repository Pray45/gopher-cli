package parser

import "fmt"

type redirectExpect int

const (
	expectNone redirectExpect = iota
	expectIn
	expectOut
	expectAppend
)

func ParseCommands(tokens []Token) ([]Command, error) {
	var commands []Command
	var current Command
	expect := expectNone

	for _, tok := range tokens {

		switch tok.Type {

		case WORD:
			if expect != expectNone {
				switch expect {
				case expectIn:
					current.Stdin = tok.Value
				case expectOut:
					current.Stdout = tok.Value
					current.Append = false
				case expectAppend:
					current.Stdout = tok.Value
					current.Append = true
				}
				expect = expectNone
			} else {
				current.Args = append(current.Args, tok.Value)
			}

		case REDIRECT_IN:
			if expect != expectNone {
				return nil, fmt.Errorf("syntax error")
			}
			expect = expectIn

		case REDIRECT_OUT:
			if expect != expectNone {
				return nil, fmt.Errorf("syntax error")
			}
			expect = expectOut

		case REDIRECT_APPEND:
			if expect != expectNone {
				return nil, fmt.Errorf("syntax error")
			}
			expect = expectAppend

		case PIPE:
			if expect != expectNone || len(current.Args) == 0 {
				return nil, fmt.Errorf("invalid pipe usage")
			}
			commands = append(commands, current)
			current = Command{}
		}
	}

	if expect != expectNone {
		return nil, fmt.Errorf("redirection missing file")
	}

	if len(current.Args) > 0 {
		commands = append(commands, current)
	}

	return commands, nil
}
