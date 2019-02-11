package utility

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message error) {
	RespondWithJSON(w, code, map[string]interface{}{"error": message.Error()})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	resp, e := json.Marshal(payload)
	if e != nil {
		log.Printf("Failed to marshal message %v\n", e.Error())
		resp = []byte(e.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}
