package main

import (
	"bufio"
	"fmt"
	"gophercli/external"
	"gophercli/helper"
	"gophercli/parser"
	"os"
	"strings"
)

func main() {

	render := bufio.NewScanner(os.Stdin)

	for {
		dir, _ := os.Getwd()
		basedir := dir[strings.LastIndex(dir, "/")+1:]
		fmt.Printf("[Gopher $%s ] > ", basedir)
		if !render.Scan() {
			fmt.Println()
			break
		}
		line := render.Text()
		if line == "exit" {
			break
		}

		runCommand(line)

	}
}

func runCommand(line string) {

	line = strings.TrimSpace(line)
	if line == "" {
		return
	}

	tokens, err := parser.Lexcmd(line)
	if err != nil {
		fmt.Println("parse error:", err)
		return
	}

	cmd := tokens[0].Value
	args := []string{}
	for _, t := range tokens[1:] {
		args = append(args, t.Value)
	}

	if helper.IsBuiltin(cmd) {

		switch cmd {

		case "cd":
			helper.BuiltinCd(args)
		case "pwd":
			helper.BuiltinPwd()
		case "exit":
			helper.BuiltinExit()
		}

		return
	}

	external.Externalcmd(cmd, args)

}
