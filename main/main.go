package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

	if cmd == "hello" || cmd == "hii" || cmd == "hi" {
		println("Hello, gopher cli this side...")
		return
	}

	command := exec.Command(cmd, params...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	

	if err := command.Run(); err != nil {
		fmt.Println("error:", err)
	}

}

