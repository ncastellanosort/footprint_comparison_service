package main

import "gorm.io/gorm"

type Data struct {
	Date      string         `json:"date"`
	Energy    map[string]float32 `json:"energy"`
	Food      map[string]float32 `json:"food"`
	Transport map[string]float32 `json:"transport"`
	Waste     map[string]float32 `json:"waste"`
	Result float32 `json:"result"`
}

type Carbon_results struct {
	gorm.Model
	Total float32
	User_id int
}

type Transports struct {
	gorm.Model
	Car_km      float32
	Public_km      float32
	Domestic_flights      float32
	International_flights      float32
	Total          float32
	User_id        int
}

type Energies struct {
	gorm.Model
	Appliance_hours      float32
	Light_bulbs      float32
	Gas_tanks      float32
	Hvac_hours      float32
	Total          float32
	User_id        int
}

type Foods struct {
	gorm.Model
	Red_meat      float32
	White_meat      float32
	Dairy      float32
	Vegetarian      float32
	Total          float32
	User_id        int
}

type Wastes struct {
	gorm.Model
	Trash_bags      float32
	Food_waste      float32
	Plastic_bottles      float32
	Paper_packages      float32
	Total          float32
	User_id        int
}

type WastesResponse struct {
	Current Wastes
	Previous Wastes
}

type CompareData struct {
	Previous Data `json:"previous"`
	Current Data `json:"current"`
}

type AreaComparison struct {
    Status string  `json:"status"`
    Diff   float32 `json:"diff"`
}

type CompareResponse struct {
    Previous   Data                 `json:"previous"`
    Current    Data                 `json:"current"`
    Comparison map[string]AreaComparison `json:"comparison"`
}


