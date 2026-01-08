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

func runCmd(inputln string) {

	tokens, err := parser.Tokenize(inputln)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	if strings.Contains(inputln, "|") {
		helper.RunPipeline(inputln)
		return
	}

	cmd := tokens[0]
	params := tokens[1:]

	if helper.IsBuiltin(cmd) {
		helper.Run(cmd, params)
		return
	}
	
	helper.ExternalRun(cmd, params)
}
