package helpers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Handlers struct{}

func JsonResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, "message")
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "applications/json")
	w.WriteHeader(code)
	w.Write(response)
}

// CORS ->
func (h Handlers) CORS(res http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	// CORS support for Preflighted requests
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, PATCH, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")

	next(res, req)
}