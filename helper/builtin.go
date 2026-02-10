package helper

import (
	"fmt"
	"os"
)

func IsBuiltin(cmd string) bool {
	return cmd == "cd" || cmd == "pwd" || cmd == "exit"
}

func BuiltinCd(args []string) {
	dir := os.Getenv("HOME")
	if len(args) > 0 {
		dir = args[0]
	}
	if err := os.Chdir(dir); err != nil {
		fmt.Println(err)
	}
}

func BuiltinPwd() {
	dir, _ := os.Getwd()
	fmt.Println(dir)
}

func BuiltinExit() {
	os.Exit(0)
}
