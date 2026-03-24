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
	fname := fmt.Sprintf("internal/data/cards/%s.csv", ctx)
	_, err := os.Stat(fname)

	if err != nil {
		sendError(w, "file_not_found", ctx)
		return
	}

	file, err := os.Open(fname)
	if err != nil {
		sendError(w, "file_not_accessable", fname)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'

	records, err := reader.ReadAll()
	if err != nil {
		sendError(w, "file_not_readable", fname)
		return
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
		sendError(w, "db_error")
	}

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

	countries, err := countryModel.Scan(rows)
	if err != nil {
	    panic(err)
	}
	
	cards := struct {
		Total int            `json:"total"`
		Items []repo.Country `json:"items"`
	}{
		Total: len(countries),
		Items: countries,
	}

	txt, _ := json.Marshal(cards)
	sendRes(w, txt)
}

type apiError struct {
    Status int
    Err    string
    Txt    string
}

var errorMsg = map[string]apiError{
	"file_not_found" 		: {
		Status 	: http.StatusNotFound,
		Err 	: "invalid_arg",
		Txt 	: "Unkown context [%s]",
	},
	"file_not_accessable" 	: {
		Status 	: http.StatusForbidden,
		Err 	: "file_access",
		Txt 	: "Arquivo [%s] não encontrado.",
	},
	"file_not_readable" 	: {
		Status 	: http.StatusInternalServerError,
		Err 	: "file_access",
		Txt 	: "Não foi possível ler o arquivo [%s].",
	},
	"db_error" 				: {
		Status 	: http.StatusInternalServerError,
		Err 	: "db_not_accessable",
		Txt 	: "Não foi possível conectar ao banco-de-dados.",
	},
}

func sendError(w http.ResponseWriter, key string, args ...any) {
    msg, ok := errorMsg[key]
    if !ok {
        msg = apiError{
            Status: http.StatusInternalServerError,
            Err:    "unknown_error",
            Txt:    "The error informed is invalid",
        }
    }

    response := map[string]string{
        "err": msg.Err,
        "txt": fmt.Sprintf(msg.Txt, args...),
    }

	res, err := json.Marshal(response)
    if err != nil {
        http.Error(w, "internal error", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(msg.Status)
    w.Write(res)
}

func sendRes(w http.ResponseWriter, txt []byte) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(txt))
}
