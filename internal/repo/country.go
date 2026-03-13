// internal/repo/country.go

package repo

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
