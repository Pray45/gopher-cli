package external

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func Externalcmd(input string) {
	parts := strings.Fields(input)

	if len(parts) == 0{
		return
	}

	cmd := exec.Command(parts[0], parts[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println("error:", err)
	}

}
