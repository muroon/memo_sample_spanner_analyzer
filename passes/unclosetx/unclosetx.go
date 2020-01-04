package unclosetx

import (

	zaganeunclosetx "github.com/gcpug/zagane/passes/unclosetx"

	"github.com/gostaticanalysis/comment/passes/commentmap"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

var closeMethods = "Close"

var Analyzer = &analysis.Analyzer{
	Name: "unclosetx",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		buildssa.Analyzer,
		commentmap.Analyzer,
	},
}

const Doc = "unclosetx for memo_sample_spanner finds transactions which does not close"

func run(pass *analysis.Pass) (interface{}, error) {
	funcs := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs

	newFuncs := make([]*ssa.Function, 0, len(funcs))
	for _, f := range funcs {
		if ignoreFunc(f) {
			continue
		}
		newFuncs = append(newFuncs, f)
	}

	pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA).SrcFuncs = newFuncs

	return zaganeunclosetx.Analyzer.Run(pass)
}

var ignorePkgFuncNames = map[string]map[string]bool{
	"cloudspanner" : map[string]bool {
		"ReadRow": true,
		"Read": true,
		"ReadUsingIndex": true,
		"Query": true,
		"Apply": true,
	},
}


func ignoreFunc(f *ssa.Function) bool {
	funcs, ok := ignorePkgFuncNames[f.Pkg.Pkg.Name()]
	if !ok {
		return false
	}

	_, ok = funcs[f.Name()]
	return ok
}