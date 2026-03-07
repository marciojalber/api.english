package routes

import (
    "net/http"
	"fmt"
	"os"
	"encoding/csv"
	"encoding/json"
)

func apiCardsHandler(w http.ResponseWriter, r *http.Request) {
    ctx := r.URL.Query().Get("context")

    if ctx == "" {
	    msg := `{"err": "missing_arg", "txt": "Context UNDEFINDED for the cards"}`
        fmt.Println(w, msg)
        return
    }    

	fname 		:= fmt.Sprintf("data/cards/%s.csv", ctx)
	_, err 		:= os.Stat(fname)

	if err != nil {
	    msg := fmt.Sprintf(`{"err": "invalid_arg", "txt": "Unkown context [%s]."}`, ctx)
        fmt.Fprintln(w, msg)
        return
	}

	file, err 	:= os.Open(fname)
	
	if err != nil {
		panic(err)
	}
	
	defer file.Close()

	reader 			:= csv.NewReader(file)
	reader.Comma 	= ';'
	records, err 	:= reader.ReadAll()
	
	if err != nil {
		panic(err)
	}

	labels 			:= records[0]
	var res_obj []map[string]string

	for _, row := range records[1:] {
		line := map[string]string{}

		for i, val := range row {
			line[labels[i]] = val
		}

		res_obj = append(res_obj, line)
	}

	json, err 		:= json.Marshal(res_obj)
	
	if err != nil {
		panic(err)
	}

	res := fmt.Sprintf(`{"total": %d, "items": %s}`, len(res_obj), string(json))
	fmt.Fprintln(w, string(res))
}
