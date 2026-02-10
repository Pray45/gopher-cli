package parser

type TokenType int

const (
	WORD TokenType = iota
	PIPE
	REDIRECT_IN
	REDIRECT_OUT
	REDIRECT_APPEND
)

type Token struct {
	Type  TokenType
	Value string
}
