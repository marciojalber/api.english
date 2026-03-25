// internal/src/utils.go

package src

import (
    "net/http"
)

type apiError struct {
    Status int
    Err    string
    Txt    string
}

func SendError(w http.ResponseWriter, key string, args ...any) {
    msg, ok := errorMsg[key]
    if !ok {
        msg = ApiError{
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

func src.SendRes(w http.ResponseWriter, txt []byte) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprint(w, string(txt))
}
