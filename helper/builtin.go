package helper

import (
	"fmt"
	"os"
)

func IsBuiltin(cmd string) bool {
	return cmd == "cd" || cmd == "pwd" || cmd == "exit" || cmd == "help"
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

func BuiltinHelp() {
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
	fmt.Println("Built-in commands:")
	fmt.Println("")
	fmt.Println("  cd [dir]   Change the current directory (defaults to $HOME)")
	fmt.Println("  pwd        Print the current working directory")
	fmt.Println("  exit       Exit the shell")
	fmt.Println("  help       Show this help message")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("External commands:")
	fmt.Println("")
	fmt.Println("  Any executable available in your PATH can be run.")
	fmt.Println("  Common examples:")
	fmt.Println("    echo   Print text to standard output")
	fmt.Println("    ls     List directory contents")
	fmt.Println("    cat    Concatenate and print files")
	fmt.Println("    grep   Search text matching a pattern")
	fmt.Println("    wc     Count lines, words, and bytes")
	fmt.Println("    mkdir  Create directories")
	fmt.Println("    rm     Remove files or directories")
	fmt.Println("    mv     Move or rename files/directories")
	fmt.Println("    cp     Copy files or directories")
}
