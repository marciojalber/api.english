// internal/routes/cards_handler.go

package routes

import (
    "net/http"
    "fmt"
    "os"
    "encoding/csv"

    "github.com/marciojalber/api.english/pkg/utils"
    "github.com/marciojalber/api.english/internal/db"
)

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

    fname       := fmt.Sprintf("data/cards/%s.csv", ctx)
    _, err      := os.Stat(fname)

    if err != nil {
        res := utils.ToJson(utils.JsonMap{
            "err": "invalid_arg",
            "txt": fmt.Sprintf("Unkown context [%s]", ctx),
        })
        fmt.Fprint(w, res)
        return
    }

    file, err   := os.Open(fname)
    
    if err != nil {
        panic(err)
    }
    
    defer file.Close()

    reader          := csv.NewReader(file)
    reader.Comma    = ';'
    records, err    := reader.ReadAll()
    
    if err != nil {
        panic(err)
    }

    labels          := records[0]
    var res_obj []utils.JsonMap

    for _, row := range records[1:] {
        line := utils.JsonMap{}

        for i, val := range row {
            line[labels[i]] = val
        }

        res_obj = append(res_obj, line)
    }

    db, err := db.MyCon()
    if err != nil {
        panic(err)
    }

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    res := utils.ToJson(utils.JsonMap{
        "total": len(res_obj),
        "my_con": "OK",
        "items": res_obj,
    })

    fmt.Fprint(w, res)
}
