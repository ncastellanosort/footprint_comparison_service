package main

import (
	"encoding/json"
	"net/http"
)

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	var d CompareData
	
	json.NewDecoder(r.Body).Decode(&d)

	w.Header().Add("Content-type", "applicattion/json")
	w.Write([]byte("HOLA"))
}
