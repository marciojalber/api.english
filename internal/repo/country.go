// internal/repo/country.go

package repo

import (
    "database/sql"
    "github.com/marciojalber/api.english/internal/src"
)

// @todo To capture tableName and repoName directly from here
// @todo To create a script to mount automaticly from all repos [FieldMap] and [Scan]
type Country struct {
    ID        uint `col:"id"`
    Continent string `col:"continent"`
    Name      string `col:"name"`
    Citizen   string `col:"citizen"`
    Capital   string `col:"capital"`
    Language  string `col:"language"`
}

func (Country) RepoName() string {
    return "Country"
}

func (Country) TableName()string {
    return "country"
}

// @todo To replace by the reflection on the struct
func (country *Country) FieldMap() map[string]any{
    return map[string]any{
        "id":        &country.ID,
        "continent": &country.Continent,
        "name":      &country.Name,
        "citizen":   &country.Citizen,
        "capital":   &country.Capital,
        "language":  &country.Language,
    }
}

// @todo Replace src.GetScanPointer for country.ScanPointers(fields)
// @todo To transfer part of it into DAO
// Runs the Scan method
func (Country) Scan(rows *sql.Rows) ([]Country, error) {
    fields, err := rows.Columns()
    if err != nil {
        return nil, err
    }
    
    countries := []Country{}

    for rows.Next() {
        country := Country{}
        
        // ptrs, err := country.ScanPointers(fields)
        ptrs, err := src.GetScanPointer(fields, country.FieldMap(), country.RepoName())
        if err != nil {
            return nil, err
        }

        if err := rows.Scan(ptrs...); err != nil {
            return nil, err
        }

        countries = append(countries, country)
    }

    return countries, nil
}