package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/parthin-baraiya/user-mongodb/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.Router(router)
	http.Handle("/", router)

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
