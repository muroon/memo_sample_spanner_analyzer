package main

import (
	"github.com/muroon/memo_sample_spanner_analyzer"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	unitchecker.Main(memo_sample_spanner_analyzer.Analyzers()...)
}
