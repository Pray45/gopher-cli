package external

import (
	"fmt"
	"os"
	"os/exec"
)

func Externalcmd(cmd string, args []string) {
	command := exec.Command(cmd, args...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
