package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/LorenzoMinon/project2-api/db"
	"github.com/LorenzoMinon/project2-api/handlers"
)

func main() {
	conn, err := db.Connect()
	if err != nil {
		fmt.Println("Failed db init conn")
		os.Exit(1)
	}
	h := handlers.Handler{DB: conn}
	_ = h
	//http.HandleFunc("GET /products", h.GetProducts)
	http.ListenAndServe(":8080", nil)

}
