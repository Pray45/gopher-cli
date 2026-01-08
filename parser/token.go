package parser

type TokenType int

const (
	WORD TokenType = iota
	PIPE
	REDIRECT_OUT
	REDIRECT_IN
)

type Token struct {
	Type  TokenType
	Value string
}
