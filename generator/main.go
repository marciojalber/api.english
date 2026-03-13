// generator/main.go
//go:generate cmd /c echo - Executing the Go generating

package main

import (
	"fmt"
	"os"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/marciojalber/api.english/internal/src"
)

type Field struct {
	Name string
	Col  string
}

type StructInfo struct {
	Name   string
	Fields []Field
}

func main() {
	dir := src.CurrentDir() + "/internal/repo"
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

				info := StructInfo{
					Name: typeSpec.Name.Name,
				}

				fmt.Println("Struct:", info.Name)

				for _, field := range structType.Fields.List {

					col := ""
				
					if field.Tag != nil {
				        raw := field.Tag.Value
				        tag := reflect.StructTag(strings.Trim(raw, "`"))
				        col = tag.Get("col")
					}

					for _, name := range field.Names {
						info.Fields = append(info.Fields, Field{
							Name: name.Name,
							Col:  col,
						})
					}
				}

				var cases strings.Builder

				for _, f := range info.Fields {

					fmt.Fprintf(
						&cases,
						`        case "%s": ptrs[i] = &repo.%s
`,
						f.Col,
						f.Name,
					)
				}
				var out strings.Builder

				filename := "__" + strings.ToLower(info.Name) + ".go"
				path := filepath.Join(dir, filename)
				fmt.Fprintf(&out, `// internal/repo/%s

package repo

import (
	"database/sql"
	"github.com/marciojalber/api.english/internal/src"
	"fmt"
)

// Capture the fields pointers to populate the instance
func (repo *%s) ScanPointers(fields []string) ([]any, error) {
    ptrs := make([]any, len(fields))

    for i, f := range fields {
        switch f {
%s        default:
            return nil, fmt.Errorf("column %%s does not exist in [%s]", f)
        }
    }

    return ptrs, nil
}

// Proceed the scan
func (%s) Scan(rows *sql.Rows) ([]%s, error) {
    fields, err := rows.Columns()
    if err != nil {
        return nil, err
    }
    
    collection := []%s{}

    for rows.Next() {
        instance := %s{}
        
        ptrs, err := instance.ScanPointers(fields)

        if err := rows.Scan(ptrs...); err != nil {
            return nil, err
        }

        collection = append(collection, instance)
    }

    return collection, nil
}
`, filename, info.Name, cases.String(), info.Name, info.Name, info.Name, info.Name, info.Name)

				err := os.WriteFile(path, []byte(out.String()), 0644)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}
