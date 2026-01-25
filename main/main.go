package main

import (
	"bufio"
	"fmt"
	"gophercli/external"
	"gophercli/helper"
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
	parts := strings.Fields(line)
	if len(parts) == 0 {
		return
	}

	cmd := parts[0]
	args := parts[1:]

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

	external.Externalcmd(line)
}
