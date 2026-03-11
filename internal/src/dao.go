// internal/src/dao.go

package src

import (
	"strings"
    "errors"
)

/* @todo To activate
func ExtractDBFields(v any) map[string]string {

    t := reflect.TypeOf(v)

    if t.Kind() == reflect.Pointer {
        t = t.Elem()
    }

    fields := make(map[string]string, t.NumField())

    for i := 0; i < t.NumField(); i++ {

        field := t.Field(i)

        tag := field.Tag.Get("db")
        if tag != "" {
            fields[field.Name] = tag
        }
    }

    return fields
}

func GetModelCols(map[string]string)
*/

// Validate the existence of fields once and returns the fields as string
func JoinFields(fields []string, fieldMap map[string]any, repoName string) (string, error) {
    for _, field := range fields {
        if _, ok := fieldMap[field]; !ok {
            return "", errors.New("Column " + field + " does not exist in " + repoName)
        }
    }

    return strings.Join(fields, ","), nil
}

// Returns pointers to be used with Scan
func GetScanPointer(fields []string, fieldMap map[string]any, repoName string) ([]any, error) {
    cols := make([]any, len(fields))

	for i, field := range fields {
        ptr, ok := fieldMap[field]
        if !ok {
            return nil, errors.New("Column " + field + " does not exist in " + repoName)
        }
        cols[i] = ptr
	}

	return cols, nil
}
