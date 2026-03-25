// internal/routes/cards_service.go

package service

import (
	"github.com/marciojalber/api.english/internal/src"
)

var errorMsg = map[string]src.ApiError{
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
