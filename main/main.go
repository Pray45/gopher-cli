package main

import (
	"bufio"
	"fmt"
	"gophercli/external"
	"os"
)

func main() {

	render := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("[Go Shell] > ")
		if !render.Scan() {
			fmt.Println()
			break
		}
		line := render.Text()
		if line == "exit" {
			break
		}

		external.Externalcmd(line)

	}
}
