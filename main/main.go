package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gophercli/parser"
	"gophercli/helper"
	"gophercli/external"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		dir, _ := os.Getwd()
		fmt.Printf("[Gopher %s] > ", dir[strings.LastIndex(dir, "/")+1:])

		if !reader.Scan() {
			fmt.Println()
			return
		}

		line := strings.TrimSpace(reader.Text())
		if line == "" {
			continue
		}

		tokens, err := parser.Lexcmd(line)
		if err != nil {
			fmt.Println(err)
			continue
		}

		commands, err := parser.ParseCommands(tokens)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if helper.IsBuiltin(commands[0].Args[0]) {
			switch commands[0].Args[0] {
			case "cd":
				helper.BuiltinCd(commands[0].Args[1:])
			case "pwd":
				helper.BuiltinPwd()
			case "exit":
				helper.BuiltinExit()
			}
			continue
		}

		external.Execute(commands)
	}
}
