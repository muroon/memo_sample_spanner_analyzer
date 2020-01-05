package memo_sample_spanner_analyzer

import (

	"golang.org/x/tools/go/analysis"

	memounclosetx "github.com/muroon/memo_sample_spanner_analyzer/passes/unclosetx"
	"github.com/gcpug/zagane/passes/unclosetx"
	"github.com/gcpug/zagane/passes/unstopiter"
	"github.com/gcpug/zagane/passes/wraperr"
)

// Analyzers returns analyzers of zagane.
func Analyzers() []*analysis.Analyzer {
	ucla := unclosetx.Analyzer
	ucla.Requires = append(ucla.Requires, memounclosetx.Analyzer)

	return []*analysis.Analyzer{
		unstopiter.Analyzer,
		ucla,
		wraperr.Analyzer,
	}
}




