package helper

import (
	"fmt"
	"os"
	"os/exec"
)

func RunPipelineCmds(cmds [][]string) {

	if len(cmds) < 2 {
		fmt.Println("pipeline error: not enough commands")
		return
	}

	var processes []*exec.Cmd
	var prevReader *os.File

	for i, cmdParts := range cmds {

		if len(cmdParts) == 0 {
			return
		}

		cmd := exec.Command(cmdParts[0], cmdParts[1:]...)
		cmd.Stderr = os.Stderr

		// stdin
		if i == 0 {
			cmd.Stdin = os.Stdin
		} else {
			cmd.Stdin = prevReader
		}

		// stdout
		if i == len(cmds)-1 {
			cmd.Stdout = os.Stdout
		} else {
			reader, writer, err := os.Pipe()
			if err != nil {
				fmt.Println("pipe error:", err)
				return
			}

			cmd.Stdout = writer
			prevReader = reader
		}

		processes = append(processes, cmd)
	}

	for _, p := range processes {
		if err := p.Start(); err != nil {
			fmt.Println("start error:", err)
			return
		}
	}

	for _, p := range processes {
		p.Wait()
	}
}
