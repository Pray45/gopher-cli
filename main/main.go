package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"gophercli/external"
	"gophercli/helper"
	"gophercli/parser"
)

func main() {
	fmt.Println(`
   _____             _                  _____ _      _____ 
  / ____|           | |                / ____| |    |_   _|
 | |  __  ___  _ __ | |__   ___ _ __  | |    | |      | |  
 | | |_ |/ _ \| '_ \| '_ \ / _ \ '__| | |    | |      | |  
 | |__| | (_) | |_) | | | |  __/ |    | |____| |____ _| |_ 
  \_____|\___/| .__/|_| |_|\___|_|     \_____|______|_____|
              | |                                           
              |_|                                           

  A Minimal Unix-like Shell Implemented in Go
------------------------------------------------------`)
	fmt.Println("")

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
			case "help":
				helper.BuiltinHelp()
			}
			continue
		}

		external.Execute(commands)
	}
}
