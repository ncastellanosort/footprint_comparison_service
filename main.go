package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {  
	err := godotenv.Load()
	if err != nil {
		log.Println("err loading env")
  	}
	InitDB()

	port := os.Getenv("PORT")
	http.HandleFunc("/compare_service", func (w http.ResponseWriter, r *http.Request) {
		CompareHandler(w, r)
	})	
	fmt.Printf("running on %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

