package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/obynonwane/banking/domain"
	"github.com/obynonwane/banking/service"
)

func Start() {
	//custom request multiplexer
	// mux := http.NewServeMux()
	router := mux.NewRouter()
	// define  routes

	//wiring
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// starting a server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
