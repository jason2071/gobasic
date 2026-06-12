# gobasic — Go workshop repo

Collection of ~30 independent Go packages that teach Go concepts step by step. Each
runnable lesson lives in its own subdirectory with `package main` and `func main()`.

## Run a lesson

```bash
go run ./arrays/        # any lesson dir with package main
go run ./goroutine/channel/
go run ./workshop/FizzBuzz/
```

Library packages (`greet/`, `hr/`, `jsonDemo/`, `models/`, `money/`, `numbers/`,
`outlineNodes/`, `users/`) have named packages and are consumed by `package main`
lessons — they are not run directly.

## Root `main.go`

A non-functional launchpad — all imports and calls are commented out. Do not try
to `go run .` or `go build .` without adding imports first.

## Verify compilation

```bash
go build ./...          # builds all packages
go vet ./...            # static analysis
```

There are **no tests**, **no linter config**, and **no CI/CD** in this repo.

## External deps

- `github.com/rivo/uniseg` — grapheme clusters in `workshop/reverse/`
- `golang.org/x/tour` — exercise validators in `exercise/` subdirs

## Notes

- Code comments are in Thai (target audience is Thai-speaking learners).
- `.claude/settings.local.json` pre-authorises `go run *` for AI tool use.
- Module name: `gobasic`, Go 1.26.2, no workspace (`go.work` is gitignored).
- No `opencode.json` found in repo.
