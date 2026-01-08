package helper

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunPipeline(input string) {
	segments := strings.Split(input, "|")

	var commands []*exec.Cmd
	var prevStdout *os.File

	for i, segment := range segments {
		segment = strings.TrimSpace(segment)
		parts := strings.Fields(segment)

		if len(parts) == 0 {
			return
		}

		cmd := exec.Command(parts[0], parts[1:]...)
		cmd.Stderr = os.Stderr

		if i == 0 {
			cmd.Stdin = os.Stdin
		} else {
			cmd.Stdin = prevStdout
		}

		if i == len(segments)-1 {
			cmd.Stdout = os.Stdout
		} else {
			pipeReader, pipeWriter, err := os.Pipe()
			if err != nil {
				fmt.Println("pipe error:", err)
				return
			}

			cmd.Stdout = pipeWriter
			prevStdout = pipeReader
		}

		commands = append(commands, cmd)
	}

	for _, cmd := range commands {
		if err := cmd.Start(); err != nil {
			fmt.Println("start error:", err)
			return
		}
	}

	for _, cmd := range commands {
		cmd.Wait()
	}
}
