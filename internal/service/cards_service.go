// internal/routes/cards_service.go

package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/marciojalber/api.english/internal/src"
	"github.com/marciojalber/api.english/internal/repo"
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
		res, _ := json.Marshal(map[string]string{
			"err": "invalid_arg",
			"txt": fmt.Sprintf("Unkown context [%s]", ctx),
		})
		fmt.Fprint(w, string(res))
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
	lines := []map[string]string{}

	for _, row := range records[1:] {
		line := map[string]string{}

		for i, val := range row {
			line[labels[i]] = val
		}

		lines = append(lines, line)
	}

	res, _ := json.Marshal(map[string]any{
		"total": len(lines),
		"items": lines,
	})

	fmt.Fprint(w, string(res))
}

// CAPTURE DATA FROM DB
func getDataFromDB(w http.ResponseWriter) {
	db, err := src.DB.MyCon()
	if err != nil {
		panic(err)
	}

	var countryModel repo.Country
	sql := `
		SELECT
			id,
			continent,
			name,
			citizen,
			capital,
			language,
		FROM	` + countryModel.TableName()
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	countries, err := countryModel.Scan(rows)
	if err != nil {
	    panic(err)
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
