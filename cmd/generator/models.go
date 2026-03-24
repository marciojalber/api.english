// cmd/generator/models.go

package main

var caseStatementModel  string = `        case "%s": ptrs[i] = &repo.%s
`

var repoModel           string = `// internal/repo/%s

package repo

import (
    "database/sql"
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
        repo := %s{}

        ptrs, err := repo.ScanPointers(fields)
        if err != nil {
            return nil, err
        }

        if err := rows.Scan(ptrs...); err != nil {
            return nil, err
        }

        collection = append(collection, repo)
    }

    return collection, nil
}
`