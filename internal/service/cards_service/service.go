// internal/routes/cards_service.go

package service

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

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
	
	// Test the file
	fname := fmt.Sprintf("internal/data/cards/%s.csv", ctx)

	_, err := os.Stat(fname)
	if err != nil {
		src.SendError(w, "file_not_found", ctx)
		return
	}

	// Open the file
	file, err := os.Open(fname)
	if err != nil {
		src.SendError(w, "file_not_accessable", fname)
		return
	}
	defer file.Close()

	// Read the file
	reader 		 := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		src.SendError(w, "file_not_readable", fname)
		return
	}

	// Capture the label
	labels := records[0]
	lines  := []map[string]string{}

	// Capture the lines
	for _, row := range records[1:] {
		line := map[string]string{}

		for i, val := range row {
			line[labels[i]] = val
		}

		lines = append(lines, line)
	}

	// Send response
	res, _ := json.Marshal(map[string]any{
		"total": len(lines),
		"items": lines,
	})

	fmt.Fprint(w, string(res))
}

// CAPTURE DATA FROM DB
func getDataFromDB(w http.ResponseWriter) {

	// Open the database
	db, err := src.DB.MyCon()
	if err != nil {
		src.SendError(w, "db_error")
	}

	// Execute the query
	var countryModel repo.Country
	sql := `
		SELECT
			id,
			continent,
			name,
			citizen,
			capital,
			language
		FROM	` + countryModel.TableName()
	
	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Capture the values
	countries, err := countryModel.Scan(rows)
	if err != nil {
	    panic(err)
	}
	
	// Send response
	cards := struct {
		Total int            `json:"total"`
		Items []repo.Country `json:"items"`
	}{
		Total: len(countries),
		Items: countries,
	}

	txt, _ := json.Marshal(cards)
	src.SendRes(w, txt)
}
