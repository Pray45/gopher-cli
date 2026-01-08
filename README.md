# 🐹 Gopher CLI

[![Go Version](https://img.shields.io/badge/Go-1.25.5+-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A fun and simple command-line interface (CLI) tool written in Go, inspired by Unix shells. Unleash the power of your terminal with built-in commands, custom greetings, and seamless command pipelines! 🚀

## Features

- **Built-in Commands**: Supports common shell commands like `cd`, `pwd`, `env`, and `exit`.
- **Custom Commands**: Includes predefined custom commands such as greetings.
- **Command Pipelines**: Execute chained commands using the `|` operator.
- **External Commands**: Run any external program available in your PATH.
- **Interactive Mode**: Reads commands from standard input in a loop.

## Installation

### Prerequisites
- Go 1.25.5 or later

### Building from Source
1. Clone or download the project.
2. Navigate to the project directory.
3. Run the build command:

   ```bash
   go build -o gopher-cli ./main
   ```

4. The executable `gopher-cli` will be created in the current directory.

## Usage

Run the CLI tool:

```bash
./gopher-cli
```

You will see the prompt: `[gopher cli] > `

Enter commands at the prompt. Type `exit` to quit.

### Examples

- **Built-in Commands**:
  - Change directory: `cd /path/to/dir`
  - Print working directory: `pwd`
  - List environment variables: `env`
  - Exit: `exit`

- **Custom Commands**:
  - Greeting: `hello`, `hi`, or `hii` (prints "Hello, gopher cli this side...")

- **External Commands**:
  - Run any command: `ls -l`
  - Run with arguments: `echo "Hello World"`

- **Pipelines**:
  - Chain commands: `ls | grep .go | wc -l`

## Project Structure

- `main/main.go`: Entry point and main loop.
- `helper/`: Utility functions for running commands.
  - `builtin.go`: Built-in command implementations.
  - `externalRun.go`: External command execution.
  - `RunPipeline.go`: Pipeline execution logic.
- `coustom_cmd/coustom.go`: Custom command definitions.