package main

import (
	timelint "github.com/tomoemon/go-time-lint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(timelint.Analyzer) }
