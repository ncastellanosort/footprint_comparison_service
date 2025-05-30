package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("AWS_RDS_URL")

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connect db", err)
	}

}

func InitDB () {
	Connect()
}

func FindResults(id int) Carbon_results {
	var d Carbon_results
	res := DB.First(&d, id)
	if res.Error != nil {
		panic(res.Error)
	}
	return d
}

func FindWastes(id int) Wastes {
	var d Wastes
	res := DB.First(&d, id)
	if res.Error != nil {
		panic(res.Error)
	}
	return d
}

func FindEnergies(id int) Energies {
	var d Energies
	res := DB.First(&d, id)
	if res.Error != nil {
		panic(res.Error)
	}
	return d
}

func FindTransports(id int) Transports {
	var d Transports
	res := DB.First(&d, id)
	if res.Error != nil {
		panic(res.Error)
	}
	return d
}

func FindFoods(id int) Foods {
	var d Foods
	res := DB.First(&d, id)
	if res.Error != nil {
		panic(res.Error)
	}
	return d
}
