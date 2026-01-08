package main

import (
	"GoCLI/helper"
	"GoCLI/parser"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	read := bufio.NewScanner(os.Stdin)

	for {
		printPrompt()

		if !read.Scan() {
			fmt.Println()
			break
		}

		inputln := strings.TrimSpace(read.Text())
		if inputln == "" {
			continue
		}

		runCmd(inputln)
	}
}

func printPrompt() {
	fmt.Print("[gopher cli] > ")
}


func hasPipe(tokens []parser.Token) bool {
	for _, t := range tokens {
		if t.Type == parser.PIPE {
			return true
		}
	}
	return false
}

func splitByPipe(tokens []parser.Token) [][]string {
	var cmds [][]string
	var current []string

	for _, t := range tokens {
		if t.Type == parser.PIPE {
			cmds = append(cmds, current)
			current = nil
			continue
		}
		if t.Type == parser.WORD {
			current = append(current, t.Value)
		}
	}

	if len(current) > 0 {
		cmds = append(cmds, current)
	}
	return cmds
}


func runCmd(inputln string) {

	tokens, err := parser.Lex(inputln)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	if hasPipe(tokens) {
		commands := splitByPipe(tokens)
		helper.RunPipelineCmds(commands)
		return
	}

	if len(tokens) == 0 {
		return
	}

	if tokens[0].Type != parser.WORD {
		fmt.Println("syntax error")
		return
	}

	cmd := tokens[0].Value
	var params []string

	for _, t := range tokens[1:] {
		if t.Type == parser.WORD {
			params = append(params, t.Value)
		}
	}

	if helper.IsBuiltin(cmd) {
		helper.Run(cmd, params)
		return
	}

	helper.ExternalRun(cmd, params)
}
