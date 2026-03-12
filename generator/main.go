// generator/main.go
//go:generate cmd /c echo - Executing the Go generating

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
)

func main() {
	dir := "../internal/repo"
	fset := token.NewFileSet()
	files, err := filepath.Glob(filepath.Join(dir, "*.go"))

	if err != nil {
		panic(err)
	}

	for _, file := range files {

		if file[0:2] == "__" {
			continue
		}

		node, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
		if err != nil {
			panic(err)
		}

		for _, decl := range node.Decls {

			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			if genDecl.Tok != token.TYPE {
				continue
			}

			for _, spec := range genDecl.Specs {

				typeSpec := spec.(*ast.TypeSpec)

				structType, ok := typeSpec.Type.(*ast.StructType)
				if !ok {
					continue
				}

				fmt.Println("Struct:", typeSpec.Name.Name)

				for _, field := range structType.Fields.List {

					for _, name := range field.Names {
						fmt.Println("  Field:", name.Name)
					}

				}

				fmt.Println()
			}
		}
	}
}
