package utilities

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func ParseInt64(s string) int64 {
	dec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return dec
}

func Response(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorResponse(w http.ResponseWriter, code int, msg string) {
	Response(w, code, map[string]string{"message": msg})
}
