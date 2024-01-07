package main

import (
	"fmt"
	"net/http"

	"github.com/gaspartv/API-GO-integration-with-postgresql/configs"
	"github.com/gaspartv/API-GO-integration-with-postgresql/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)
	r.Get("/list", handlers.List)
	r.Get("/{id}", handlers.Get)

	// fmt.Println("Servidor escutando na porta 9000")
	// err = http.ListenAndServe(":9000", nil)
	// if err != nil {
	// 	fmt.Println("Erro ao iniciar o servidor:", err)
	// }
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
