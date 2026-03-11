// internal/repo/country.go

package repo

import (
	"database/sql"
	"github.com/marciojalber/api.english/internal/src"
)

type Country struct {
	ID        uint `col: "id"`
	Continent string `col: "continent"`
	Name      string `col: "name"`
	Citizen   string `col: "citizen"`
	Capital   string `col: "capital"`
	Language  string `col: "language"`
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

/* @ To pass part to the DAO and remain here only the cases
func (c *Country) ScanPointers(fields []string) ([]any, error) {

    ptrs := make([]any, len(fields))

    for i, f := range fields {
        switch f {
	        case "id": ptrs[i] = &c.ID
	        case "continent": ptrs[i] = &c.Continent
	        case "name": ptrs[i] = &c.Name
	        case "citizen": ptrs[i] = &c.Citizen
	        case "capital": ptrs[i] = &c.Capital
	        case "language": ptrs[i] = &c.Language
	        default: return nil, fmt.Errorf("column %s does not exist in Country", f)
        }
    }

    return ptrs, nil
}
*/

// Validate the existence of fields once and returns the fields as string
func (country *Country) JoinFields(fields []string) (string, error) {
    return src.JoinFields(fields, country.FieldMap(), country.RepoName())
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
