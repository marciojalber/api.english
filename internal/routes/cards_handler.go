// internal/routes/cards_handler.go

package routes

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/marciojalber/api.english/internal/dbhelper"
	"github.com/marciojalber/api.english/internal/entities"
	"github.com/marciojalber/api.english/pkg/utils"
)

// HANDLER
func apiCardsHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.URL.Query().Get("context")

	if ctx == "" {
		res := utils.ToJson(utils.JsonMap{
			"err": "missing_arg",
			"txt": "Context UNDEFINDED for the cards",
		})
		fmt.Fprint(w, res)
		return
	}

	if ctx != "COUNTRIES" {
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

		getDataFromFile(w, fname)
		return
	} else {
		getDataFromDB(w)
		return
	}
}

// CAPTURE DATA FROM FILE
func getDataFromFile(w http.ResponseWriter, fname string) {
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
	res_obj := []utils.JsonMap{}

	for _, row := range records[1:] {
		line := utils.JsonMap{}

		for i, val := range row {
			line[labels[i]] = val
		}

		res_obj = append(res_obj, line)
	}

	res := utils.ToJson(utils.JsonMap{
		"total": len(res_obj),
		"items": res_obj,
	})

	fmt.Fprint(w, res)
}

// CAPTURE DATA FROM DB
func getDataFromDB(w http.ResponseWriter) {
	db, err := dbhelper.MyCon()
	if err != nil {
		panic(err)
	}

	cols := []string{"id", "continent", "name", "citizen", "capital", "language"}
	sql := dbhelper.Select("SELECT :cols FROM country", cols)
	row, db_err := db.Query(sql)
	if db_err != nil {
		panic(db_err)
	}
	defer row.Close()

	res_obj := []entities.Country{}

	for {
		if row.Next() == false {
			break
		}
		country := entities.Country{}
		s_err := row.Scan(&country.ID, &country.Continent, &country.Name, &country.Citizen, &country.Capital, &country.Language)
		if s_err != nil {
			panic(s_err)
		}
		res_obj = append(res_obj, country)
	}

	type Response struct {
		Total int                `json:"total"`
		Items []entities.Country `json:"items"`
	}

	res := Response{
		Total: len(res_obj),
		Items: res_obj,
	}

	txt, _ := json.Marshal(res)

	fmt.Fprint(w, string(txt))
	// err := row.Scan(BuildScanPointers(user, columns)...)
}
