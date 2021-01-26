package main

import (
	"log"
	"net/http"

	"github.com/doCodingWhileCoding/GO-API-CRUD/authorization"
	"github.com/doCodingWhileCoding/GO-API-CRUD/handler"
	"github.com/doCodingWhileCoding/GO-API-CRUD/storage"
)

func main() {
	err := authorization.LoadFiles("certificates/app.rsa", "certificates/app.rsa.pub")
	if err != nil {
		log.Fatalf("no se pudo cargar los certificados: %v", err)
	}

	store := storage.NewMemory()
	mux := http.NewServeMux()

	handler.RoutePerson(mux, &store)
	handler.RouteLogin(mux, &store)

	log.Println("Servidor iniciado en el puerto 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Printf("Error en el servidor: %v\n", err)
	}
}
