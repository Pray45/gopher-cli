package helper

import (
	"fmt"
	"os"
)

func IsBuiltin(cmd string) bool {

	switch cmd {
	case "cd", "pwd", "env", "exit":
		return true
	default:
		return false
	}

}

func Run(cmd string, params []string) {

	switch cmd {
	case "cd":
		cd(params)
	case "pwd":
		pwd()
	case "env":
		env()
	case "exit":
		exit()
	}

}

func cd(params []string) {

	var target string

	if len(params) == 0 {
		target = os.Getenv("HOME")
		if target == "" {
			fmt.Println("cd: HOME not set")
			return
		}
	} else {
		target = params[0]
	}

	if err := os.Chdir(target); err != nil {
		fmt.Println("cd error:", err)
	}

}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd error:", err)
		return
	}
	fmt.Println(dir)
}

func env() {
	for _, e := range os.Environ() {
		fmt.Println(e)
	}
}

func exit() {
	os.Exit(0)
}
