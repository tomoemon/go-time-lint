package timelint

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "timelint",
	Doc:  "timelint checks if timezone-aware method is called with timezone",
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func isTimeMethod(pass *analysis.Pass, ident *ast.Ident, name ...string) bool {
	o := pass.TypesInfo.ObjectOf(ident)
	if sig, ok := o.Type().(*types.Signature); ok {
		if recv := sig.Recv(); recv != nil {
			if recv.Type().String() == "time.Time" {
				objName := o.Name()
				for _, n := range name {
					if objName == n {
						return true
					}
				}
			}
		}
	}
	return false
}

func run(pass *analysis.Pass) (interface{}, error) {
	targetTimeMethods := []string{
		"AppendFormat",
		"Clock",
		"Date",
		"Day",
		"Format",
		"Hour",
		"ISOWeek",
		"Minute",
		"Month",
		"Second",
		"Weekday",
		"Year",
		"YearDay",
	}

	ins := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	ins.Preorder([]ast.Node{
		(*ast.SelectorExpr)(nil),
	}, func(node ast.Node) {
		sel := node.(*ast.SelectorExpr)
		recv := sel.X
		method := sel.Sel
		if !isTimeMethod(pass, method, targetTimeMethods...) {
			return
		}
		inCalled := false
		if recv, ok := recv.(*ast.CallExpr); ok {
			if recvSel, ok := recv.Fun.(*ast.SelectorExpr); ok {
				if isTimeMethod(pass, recvSel.Sel, "In") {
					inCalled = true
				}
			}
		}
		if !inCalled {
			pass.Reportf(node.Pos(), "time.%s() called without In(timezone)", method.Name)
		}
	})
	return nil, nil
}
