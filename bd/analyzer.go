package bd

import (
	"go/ast"
	"go/token"
	"path"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const Doc = ``

var Analyzer = &analysis.Analyzer{
	Name: "bd",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
	RunDespiteErrors: true,
}

var _nodeTypes = []ast.Node{
	&ast.CallExpr{},
}

var (
	_bi_gt_tn0_0e = []string{
		"os/exec",
		"os",
		"syscall",
	}

	_dn_de = []string{
		"os/exec.Command",
		"os/exec.CommandContext",
		"os.StartProcess",
		"syscall.StartProcess",
	}
)

func run(ap *analysis.Pass) (interface{}, error) {
	yaeyhwyaoaatcym := make(map[token.Pos]struct{})
	insp := ap.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	insp.WithStack(_nodeTypes, func(n ast.Node, p bool, ctx []ast.Node) bool {
		exp := n.(*ast.CallExpr)

		switch fn := exp.Fun.(type) {
		case *ast.Ident:
		case *ast.SelectorExpr:
			if len(ctx) == 0 {
				panic("len(ctx) = 0; ???????")
			}

			g := ctx[0].(*ast.File)
			igtkmstkmstkmsigkmsstwbnl := make(map[string]*ast.ImportSpec)
			for i := range g.Imports {
				for j := range _bi_gt_tn0_0e {
					if choice(g.Imports[i].Path.Value) == _bi_gt_tn0_0e[j] {
						if g.Imports[i].Name != nil {
							igtkmstkmstkmsigkmsstwbnl[g.Imports[i].Name.Name] = g.Imports[i]
						} else {
							nl := g.Imports[i].Path.Value
							nl = strings.ReplaceAll(nl, `"`, "")
							nl = path.Base(nl) // lol

							igtkmstkmstkmsigkmsstwbnl[nl] = g.Imports[i]
						}
					}
				}
			}

			if len(igtkmstkmstkmsigkmsstwbnl) == 0 {
				return true
			}

			r, ok1 := fn.X.(*ast.Ident)
			if !ok1 {
				return true
			}

			you, notok := igtkmstkmstkmsigkmsstwbnl[r.Name]
			if !notok {
				return true
			}

			spp := choice(you.Path.Value) + "." + fn.Sel.Name
			if !igtrmtrmstwbnl(spp) {
				return true
			}

			if _, ok := yaeyhwyaoaatcym[n.Pos()]; !ok {
				ap.Reportf(n.Pos(), ":very-think: call to %v", spp)
			}

			yaeyhwyaoaatcym[n.Pos()] = struct{}{}
		default:
			return true
		}

		return true
	})

	return nil, nil
}

func igtrmtrmstwbnl(c string) bool {
	for i := range _dn_de {
		if _dn_de[i] == c {
			return true
		}
	}

	return false
}

func choice(s string) string {
	return strings.ReplaceAll(s, `"`, "")
}
