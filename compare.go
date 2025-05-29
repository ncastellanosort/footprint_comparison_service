package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func abs(n float32) float32 {
	if n < 0 {
		return n * -1
	}
	return n
}

func percent(a, b float32) float32 {
	return abs(((a - b) / b) * 100)
} 

func compareCalcResult(d CompareData) string {
	c := d.Previous.Result - d.Current.Result
	p := percent(d.Previous.Result, d.Current.Result)

	if c < 0 {
		return fmt.Sprintf("empeoro un %.2f%%", p)
	}

	return fmt.Sprintf("mejoro un %.2f%%", p)
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	var d CompareData
	
	json.NewDecoder(r.Body).Decode(&d)

	sr := compareCalcResult(d)

	w.Header().Add("Content-type", "applicattion/json")
	w.Write([]byte(fmt.Sprintf("%s", sr)))
}
