package main

import (
	"GoCLI/coustom_cmd"
	"GoCLI/helper"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	read := bufio.NewScanner(os.Stdin)

	for {

		Printdefault()

		if !read.Scan() {
			fmt.Print()
			break
		}

		inputln := strings.TrimSpace(read.Text())

		if inputln == "" {
			continue
		}

		runcmd(inputln);

	}

}

func Printdefault() {

	fmt.Print("[gopher cli] > ")

}

func runcmd (inputln string) {

	parse := strings.Fields(inputln)

	if len(parse) == 0 {
		return
	}

	cmd := parse[0]
	params := parse[1:]

	if cmd == "exit" || cmd == "EXIT" {
		os.Exit(0)
	}

	if helper.IsBuiltin(cmd) {
		helper.Run(cmd, params)
		return
	}

	coustomcmd.Coustomcmd(cmd)

	helper.ExternalRun(cmd, params);
}