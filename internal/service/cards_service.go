// internal/routes/cards_handler.go

package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/marciojalber/api.english/internal/handler"
	"github.com/marciojalber/api.english/internal/repo"
	"github.com/marciojalber/api.english/pkg/utils"
)

// SERVICE
func ApiCardsService(w http.ResponseWriter, r *http.Request) {
	ctx := r.URL.Query().Get("context")

	if ctx != "" {
		// Write the code to get all the cards
	}

	if ctx != "COUNTRIES" {
		getDataFromFile(w, ctx)
		return
	}

	getDataFromDB(w)
}

// CAPTURE DATA FROM FILE
func getDataFromFile(w http.ResponseWriter, ctx string) {
	fname := fmt.Sprintf("data/cards/%s.csv", ctx)
	_, err := os.Stat(fname)

	if err != nil {
		res := utils.ToJson(utils.JsonMap{
			"err": "invalid_arg",
			"txt": fmt.Sprintf("Unkown context [%s]", ctx),
		})
		fmt.Fprint(w, res)
		return
	}

	file, err := os.Open(fname)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	labels := records[0]
	lines := []utils.JsonMap{}

	for _, row := range records[1:] {
		line := utils.JsonMap{}

		for i, val := range row {
			line[labels[i]] = val
		}

		lines = append(lines, line)
	}

	res := utils.ToJson(utils.JsonMap{
		"total": len(lines),
		"items": lines,
	})

	fmt.Fprint(w, res)
}

// CAPTURE DATA FROM DB
func getDataFromDB(w http.ResponseWriter) {
	db, err := handler.DB.MyCon()
	if err != nil {
		panic(err)
	}

	sql := `
		SELECT 	id,
				continent,
				name,
				citizen,
				capital,
				language
		FROM 	country	`
	row, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer row.Close()

	countries := []repo.Country{}

	for {
		if !row.Next() {
			break
		}

		country := repo.Country{}
		err := row.Scan(
			&country.ID,
			&country.Continent,
			&country.Name,
			&country.Citizen,
			&country.Capital,
			&country.Language,
		)
		if err != nil {
			panic(err)
		}
		countries = append(countries, country)
	}

	type Response struct {
		Total int            `json:"total"`
		Items []repo.Country `json:"items"`
	}

	res := Response{
		Total: len(countries),
		Items: countries,
	}

	txt, _ := json.Marshal(res)
	fmt.Fprint(w, string(txt))
}
