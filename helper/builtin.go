package helper

import (
	"fmt"
	"os"
)

func IsBuiltin(cmd string) bool {
	switch cmd {
	case "cd", "pwd", "exit":
		return true
	default:
		return false
	}
}

func BuiltinCd(args []string) {

	var target string

	if len(args) == 0 {

		target = os.Getenv("HOME")
		if target == "" {
			fmt.Println("cd: HOME not set")
			return
		}
	} else {
		target = args[0]
	}

	if err := os.Chdir(target); err != nil {
		fmt.Println("cd error:", err)
	}

}

func BuiltinPwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("pwd error:", err)
		return
	}
	fmt.Println(dir)
}

func BuiltinExit() {
	os.Exit(0)
}
