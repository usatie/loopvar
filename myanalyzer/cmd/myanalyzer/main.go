package main

import (
	"myanalyzer"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(myanalyzer.Analyzer) }
