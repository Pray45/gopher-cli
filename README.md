# Gopher CLI
*A minimal Unix-like shell implemented in Go.*

## Overview

Gopher CLI is a minimal Unix-like command-line shell written in Go. The project emphasizes **correct shell architecture** (lexing, parsing, process execution, pipelines, and redirection) over feature breadth.

This is a **systems-level learning project** that mirrors Unix process and file-descriptor semantics.

**Current version:** `v1.0.0` (Core Shell Engine)

---

## Features

### Core functionality
- Interactive Read–Eval–Print Loop (REPL)
- Foreground command execution
- Built-in commands:
  - `cd`
  - `pwd`
  - `exit`

### Parsing and syntax handling
- Character-level lexer with support for:
  - Single quotes (`'...'`)
  - Double quotes (`"..."`)
  - Escaped characters (`\\`)
- No string-based command execution
- Input is fully tokenized and parsed before execution

### Pipelines
- Support for multi-stage pipelines using `|`
- All pipeline processes start concurrently
- Deadlock-safe execution

### Input and output redirection
- Input redirection (`<`)
- Output redirection (`>`)
- Output append redirection (`>>`)

### Architecture
- Clear separation of responsibilities:
  - **Lexer** → **Parser** → **Executor**
- Operators are resolved during parsing, not execution
- Shell state is maintained in the shell process, not child processes

---

## Design principles

This project emphasizes **correctness and extensibility** over convenience.

Key design decisions:
- Parsing is implemented as a state machine rather than string splitting
- Raw input strings are never used after parsing
- Execution logic is intentionally simple and relies entirely on parsed structures
- Built-in commands execute within the shell process
- Pipelines are executed concurrently, following Unix semantics

This structure allows new features to be added without architectural changes.

---

## Work Flow

![Gopher Workflow](github/Gopher%20Workflow.png)

---

## Project structure

```
gopher-cli/
├── go.mod
├── README.md
├── main/
│   └── main.go         # Shell REPL and orchestration
├── parser/
│   ├── command.go      # Command representation
│   ├── lex.go          # Character-level lexer
│   ├── parse.go        # Grammar and syntax validation
│   └── token.go        # Token definitions
├── external/
│   └── external.go     # Process execution, pipes, redirection
└── helper/
  └── builtin.go        # Built-in shell commands
```


---

## Getting started

### Requirements
- Go 1.20+ (earlier versions may work, but are untested)

### Build
```bash
go build -o gopher-cli ./main
```

### Run
```bash
./gopher-cli
```

## Example usage

```bash
[Gopher main] > echo "hello world"
hello world

[Gopher main] > ls | wc -l
12

[Gopher main] > cat < input.txt | grep foo >> output.txt
```

## Known limitations

The following features are intentionally not included in version 1.0.0:
- Signal handling (e.g., Ctrl-C / SIGINT)
- Background job execution (`&`)
- Job control (`jobs`, `fg`, `bg`)
- Environment variable expansion (`$VAR`)
- Non-interactive scripting mode
- Command history and line editing

These limitations are planned for future releases.

## Roadmap

- v1.1.0 — Signal handling and Ctrl-C safety
- v1.2.0 — Background jobs
- v1.3.0 — Job control
- v1.4.0 — Environment variables and expansion
- v2.0.0 — Scripting support and advanced parsing

## Why Go?

Go was chosen for its suitability in systems programming:
- `os/exec` maps cleanly to Unix process creation
- Goroutines and `sync.WaitGroup` support correct pipeline concurrency
- Strong typing improves parser and executor reliability
- Simple build and deployment workflow

Go is used deliberately and conservatively, only where it aligns with Unix semantics.