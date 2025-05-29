package main

import (
	"encoding/json"
	"net/http"
	"strconv"
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

func CompareHandler(w http.ResponseWriter, r *http.Request) {
	// cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Add("Content-type", "applicattion/json")
	
	prev := r.URL.Query().Get("prev")
	curr := r.URL.Query().Get("curr")

	prevId, err1 := strconv.Atoi(prev)
	currId, err2 := strconv.Atoi(curr)

	if err1 != nil || err2 != nil {
		panic(err1)
	}

	resultPrev := FindResults(prevId)
	resultCurr := FindResults(currId)

	wastePrev := FindWastes(prevId)
	wasteCurr := FindWastes(currId)

	transportPrev := FindTransports(prevId)
	transportCurr := FindTransports(currId)

	energyPrev := FindEnergies(prevId)
	energyCurr := FindEnergies(currId)

	foodPrev := FindFoods(prevId)
	foodCurr := FindFoods(currId)

	tl, vt := comparison(resultPrev.Total, resultCurr.Total)
	ws, vw := comparison(wastePrev.Total, wasteCurr.Total)
	fd, vf := comparison(foodPrev.Total, foodCurr.Total)
	tr, vtr := comparison(transportPrev.Total, transportCurr.Total)
	en, ve := comparison(energyPrev.Total, energyCurr.Total)

	energyPrevious := map[string]float32{
		"appliance_hours": energyPrev.Appliance_hours,
		"light_bulbs": energyPrev.Light_bulbs,
		"gas_tanks": energyPrev.Gas_tanks,
		"hvac_hours": energyPrev.Hvac_hours,
	}

	energyCurrent := map[string]float32{
		"appliance_hours": energyCurr.Appliance_hours,
		"light_bulbs": energyCurr.Light_bulbs,
		"gas_tanks": energyCurr.Gas_tanks,
		"hvac_hours": energyCurr.Hvac_hours,
	}

	foodPrevious := map[string]float32{
		"red_meat": foodPrev.Red_meat,
		"white_meat": foodPrev.White_meat,
		"dairy": foodPrev.Dairy,
		"vegetarian": foodPrev.Vegetarian,
	}

	foodCurrent := map[string]float32{
		"red_meat": foodCurr.Red_meat,
		"white_meat": foodCurr.White_meat,
		"dairy": foodCurr.Dairy,
		"vegetarian": foodCurr.Vegetarian,
	}

	transportPrevious := map[string]float32{
		"car_km": transportPrev.Car_km,
		"public_km": transportPrev.Public_km,
		"domestic_flights": transportPrev.Domestic_flights,
		"international_fligts": transportPrev.International_flights,
	}

	transportCurrent := map[string]float32{
		"car_km": transportCurr.Car_km,
		"public_km": transportCurr.Public_km,
		"domestic_flights": transportCurr.Domestic_flights,
		"international_fligts": transportCurr.International_flights,
	}

	wastePrevious := map[string]float32{
		"trash_bags": wastePrev.Trash_bags,
		"food_waste": wastePrev.Food_waste,
		"plastic_bottles": wastePrev.Plastic_bottles,
		"paper_packages": wastePrev.Paper_packages,
	}

	wasteCurrent := map[string]float32{
		"trash_bags": wasteCurr.Trash_bags,
		"food_waste": wasteCurr.Food_waste,
		"plastic_bottles": wasteCurr.Plastic_bottles,
		"paper_packages": wasteCurr.Paper_packages,
	}


	previous := Data{
		Date: "",
		Energy: energyPrevious,
		Food: foodPrevious,
		Transport: transportPrevious,
		Waste: wastePrevious,
		Result: resultPrev.Total,
	}

	current := Data{
		Date: "",
		Energy: energyCurrent,
		Food: foodCurrent,
		Transport: transportCurrent,
		Waste: wasteCurrent,
		Result: resultCurr.Total,
	}
	resp := CompareResponse{
		Previous: previous,
		Current: current,
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
