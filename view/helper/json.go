package helper

import (
	"encoding/json"
	"net/http"
)

func RenderJson(obj interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(obj)
}
