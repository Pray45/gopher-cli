package parser

type Command struct {
	Args   []string
	Stdin  string
	Stdout string
	Append bool
}
