package external

import (
	"os"
	"os/exec"
	"sync"

	"gophercli/parser"
)

func Execute(commands []parser.Command) {

	var wg sync.WaitGroup
	var prevReader *os.File

	for i, cmd := range commands {

		execCmd := exec.Command(cmd.Args[0], cmd.Args[1:]...)
		execCmd.Stderr = os.Stderr

		// ---- stdin wiring ----
		if cmd.Stdin != "" {
			file, err := os.Open(cmd.Stdin)
			if err != nil {
				panic(err)
			}
			execCmd.Stdin = file
		} else if prevReader != nil {
			execCmd.Stdin = prevReader
		} else {
			execCmd.Stdin = os.Stdin
		}

		// ---- stdout wiring ----
		if cmd.Stdout != "" {
			flags := os.O_CREATE | os.O_WRONLY
			if cmd.Append {
				flags |= os.O_APPEND
			} else {
				flags |= os.O_TRUNC
			}
			file, err := os.OpenFile(cmd.Stdout, flags, 0644)
			if err != nil {
				panic(err)
			}
			execCmd.Stdout = file
		} else if i < len(commands)-1 {
			reader, writer, err := os.Pipe()
			if err != nil {
				panic(err)
			}
			execCmd.Stdout = writer
			prevReader = reader
		} else {
			execCmd.Stdout = os.Stdout
		}

		// ---- start process ----
		if err := execCmd.Start(); err != nil {
			panic(err)
		}

		// ---- wait concurrently ----
		wg.Add(1)
		go func(c *exec.Cmd) {
			defer wg.Done()
			c.Wait()
		}(execCmd)
	}

	// ---- wait for ALL processes ----
	wg.Wait()
}
