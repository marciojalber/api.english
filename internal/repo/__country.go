// internal/repo/__country.go

package repo

import (
	"database/sql"
	"github.com/marciojalber/api.english/internal/src"
	"fmt"
)

// Capture the fields pointers to populate the instance
func (repo *Country) ScanPointers(fields []string) ([]any, error) {
    ptrs := make([]any, len(fields))

    for i, f := range fields {
        switch f {
        case "id": ptrs[i] = &repo.ID
        case "continent": ptrs[i] = &repo.Continent
        case "name": ptrs[i] = &repo.Name
        case "citizen": ptrs[i] = &repo.Citizen
        case "capital": ptrs[i] = &repo.Capital
        case "language": ptrs[i] = &repo.Language
        default:
            return nil, fmt.Errorf("column %s does not exist in [Country]", f)
        }
    }

    return ptrs, nil
}

// Proceed the scan
func (Country) Scan(rows *sql.Rows) ([]Country, error) {
    fields, err := rows.Columns()
    if err != nil {
        return nil, err
    }
    
    collection := []Country{}

    for rows.Next() {
        instance := Country{}
        
        ptrs, err := instance.ScanPointers(fields)

        if err := rows.Scan(ptrs...); err != nil {
            return nil, err
        }

        collection = append(collection, instance)
    }

    return collection, nil
}
