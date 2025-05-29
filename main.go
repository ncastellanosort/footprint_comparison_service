package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {  
	err := godotenv.Load()
	if err != nil {
		log.Fatal("err loading env")
  	}
	InitDB()
	http.HandleFunc("/compare_service", func (w http.ResponseWriter, r *http.Request) {
		CompareHandler(w, r)
	})	
	fmt.Printf("running on 8080")
	http.ListenAndServe(":8080", nil)
}

