package myanalyzer

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "myanalyzer is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "myanalyzer",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ForStmt:
			// pass.Reportf(n.Pos(), "for found")
			findLoopVar(pass, n)
		}
	})

	return nil, nil
}

func findLoopVar(pass *analysis.Pass, forstmt *ast.ForStmt) {
	assignStmt, ok := forstmt.Init.(*ast.AssignStmt)
	if !ok {
		return
	}

	if len(assignStmt.Lhs) == 0 {
		return
	}

	ident, ok := assignStmt.Lhs[0].(*ast.Ident)
	if !ok {
		return
	}

	if assignStmt != ident.Obj.Decl {
		return
	}
	pass.Reportf(assignStmt.Pos(), "%v found", ident)

	findPointerOfLoopVar(pass, assignStmt, forstmt.Body)
}

func findPointerOfLoopVar(pass *analysis.Pass, decl *ast.AssignStmt, body *ast.BlockStmt) {

	ast.Inspect(body, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		switch n := n.(type) {
		case *ast.UnaryExpr:
			// & じゃなかったら返す
			//if n.Op != token.AND {
			//	return true
			//}

			// x -> &の引数
			x, ok := n.X.(*ast.Ident)
			if !ok {
				return true
			}

			// xが宣言されている場所が、ループ変数と一致するときレポート
			if x.Obj.Decl == decl {
				pass.Reportf(n.Pos(), "unary expr found")
			}
		}
		return true
	})
}
