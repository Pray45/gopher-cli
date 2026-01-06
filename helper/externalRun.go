package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func ExternalRun(cmd string, params []string) {
	command := exec.Command(cmd, params...)

	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	if err := command.Run(); err != nil {
		fmt.Println("error:", err)
	}
}
