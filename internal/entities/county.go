package entities

type Country struct {
	ID        uint   `json: "id"`
	Continent string `json: continent`
	Name      string `json: ciziten`
	Citizen   string `json: ciziten`
	Capital   string `json: capital`
	Language  string `json: language`
}
