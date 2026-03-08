// pkg/json.go

package utils

import "encoding/json"

type JsonMap map[string]any

func ToJson(obj JsonMap) string {
	res, _ := json.Marshal(obj)
	return string(res)
}
