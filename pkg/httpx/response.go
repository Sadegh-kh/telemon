package httpx

import (
	"encoding/json"
	"net/http"
)

func ResponsdJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func ResposndErrorJSON(w http.ResponseWriter, code int, err error) {
	ResponsdJSON(w, code, map[string]string{
		"error": err.Error(),
	})
}
