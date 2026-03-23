// generator/main.go
//go:generate cmd /c echo - Executing the Go generating

package main

import (
    "log"
    "regexp"
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
    dir     := src.DirBase() + "/internal/repo"
    fset    := token.NewFileSet()
    files, err := filepath.Glob(filepath.Join(dir, "*.go"))

    if err != nil {
        panic(err)
    }

    for _, file := range files {

        if file[0:2] == "__" {
            continue
        }

        re := regexp.MustCompile(`__`)
        if re.MatchString(file) {
            continue
        }

        file_dest := dir + "/__" + file[len(dir)+1:]

        infoDest, err := os.Stat(file_dest)
        if err == nil {
            infoOrig, err := os.Stat(file)
            if err != nil {
                log.Fatal("[cmd/generator/main.go] Erro: ", err)
            }

            iOrig := infoOrig.ModTime()
            if iOrig.Before(infoDest.ModTime()) || iOrig.Equal(infoDest.ModTime()) {
                continue
            }
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

                fmt.Println("Repo created for:", info.Name)

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
                        caseStatementModel,
                        f.Col,
                        f.Name,
                    )
                }
                var out strings.Builder

                filename := "__" + strings.ToLower(info.Name) + ".go"
                path := filepath.Join(dir, filename)
                fmt.Fprintf(&out, repoModel, filename, info.Name, cases.String(), info.Name, info.Name, info.Name, info.Name, info.Name)

                err := os.WriteFile(path, []byte(out.String()), 0644)
                if err != nil {
                    panic(err)
                }
            }
        }
    }
}