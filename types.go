package main

type Data struct {
	Date      string         `json:"date"`
	Energy    map[string]int `json:"energy"`
	Food      map[string]int `json:"food"`
	Transport map[string]int `json:"transport"`
	Waste     map[string]int `json:"waste"`
	Result float32 `json:"result"`
}

type CompareData struct {
	Previous Data `json:"previous"`
	Current Data `json:"current"`
}


