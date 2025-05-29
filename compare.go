package main

import (
	"encoding/json"
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

func comparison(a, b float32) (string, float32) {
	c := a - b
	p := percent(a, b)

	if c < 0 {
		return "empeoro", p
	}

	return "mejoro", p
}

func totalArea(d map[string]int) float32 {
	var n float32
	for _, v := range d {
		n += float32(v)
	}
	return n
}

func compareAreas(a, b map[string]int) (string, float32) {
	prev := totalArea(a)
	curr := totalArea(b)

	return comparison(prev, curr)
}

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	var d CompareData

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Add("Content-type", "applicattion/json")
	
	json.NewDecoder(r.Body).Decode(&d)

	tl, vt := comparison(d.Previous.Result, d.Current.Result)
	
	ws, vw := compareAreas(d.Previous.Waste, d.Current.Waste)
	fd, vf := compareAreas(d.Previous.Food, d.Current.Food)
	tr, vtr := compareAreas(d.Previous.Transport, d.Current.Transport)
	en, ve := compareAreas(d.Previous.Energy, d.Current.Energy)

	resp := CompareResponse{
		Previous: d.Previous,
		Current: d.Current,
		Comparison: map[string]AreaComparison{
			"total":     {Status: tl, Diff: vt},
			"waste":     {Status: ws, Diff: vw},
			"food":      {Status: fd, Diff: vf},
			"transport": {Status: tr, Diff: vtr},
			"energy":    {Status: en, Diff: ve},
		},
	}

	json.NewEncoder(w).Encode(resp)
}
