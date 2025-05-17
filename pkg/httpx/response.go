package httpx

import (
	"encoding/json"
	"net/http"
)

func ResponsdJson(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func ResposndErrorJson(w http.ResponseWriter, code int, err error) {
	ResponsdJson(w, code, map[string]string{
		"error": err.Error(),
	})
}
