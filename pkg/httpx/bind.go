package httpx

import (
	"encoding/json"
	"net/http"
)

func BindJSON(r *http.Request, req interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	return nil
}
