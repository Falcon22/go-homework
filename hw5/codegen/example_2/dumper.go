package main

// go build dumper.go && dumper.exe input.go

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, os.Args[1], nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range node.Decls {
		g, ok := f.(*ast.FuncDecl)
		if !ok {
			fmt.Printf("SKIP %T is not *ast.GenDecl\n", f)
			continue
		}

		fmt.Printf("%#v\n", g.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident).Name)

	}

}
