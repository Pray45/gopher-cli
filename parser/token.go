package parser

type TokenType int

const (
	WORD TokenType = iota
)

type Token struct {
	Type  TokenType
	Value string
}
