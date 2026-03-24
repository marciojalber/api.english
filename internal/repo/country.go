// internal/repo/country.go

package repo

import (
    // "github.com/marciojalber/api.english/internal/src"
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
