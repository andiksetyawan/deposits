package web

import (
	"encoding/json"
	"log"
	"net/http"
)

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}, headMap map[string]string) {
	w.Header().Add("Content-Type", "application/json")
	if headMap != nil && len(headMap) > 0 {
		for key, val := range headMap {
			w.Header().Add(key, val)
		}
	}
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(data)
	w.Write(jsonData)
}

func WriteFailResponse(w http.ResponseWriter, statusCode int, error interface{}) {
	log.Println(error)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonData, _ := json.Marshal(error)
	w.Write(jsonData)
}
