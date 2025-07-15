package api

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	h := &http.Server{
		Addr:    ":8080",
		Handler: setupRouter(),
	}

	fmt.Println("Server Started on ", h.Addr)
	log.Fatal(h.ListenAndServe())
}
