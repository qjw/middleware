package middleware

import (
	"net/http"
	"strconv"
	"strings"
)

func normalize(values []string) []string {
	if values == nil {
		return nil
	}
	distinctMap := make(map[string]bool, len(values))
	normalized := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.ToLower(value)
		if _, seen := distinctMap[value]; !seen {
			normalized = append(normalized, value)
			distinctMap[value] = true
		}
	}
	return normalized
}

type converter func(string) string

func convert(s []string, c converter) []string {
	var out []string
	for _, i := range s {
		out = append(out, c(i))
	}
	return out
}

func setHttpCode(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Write([]byte(strconv.Itoa(code) + " - " + http.StatusText(code)))
}
