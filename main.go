package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string
	City    string
	Zipcode string
}

func main() {

	//define  routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//starting a server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hellow World!")
}
func getAllCustomers(w http.ResponseWriter, r *http.Request) {

	//create a slice of customers
	customers := []Customer{
		{Name: "Ahish", City: "New Delhi", Zipcode: "7895759"},
		{Name: "Obinna", City: "Logand", Zipcode: "78575"},
	}

	//make the response header to be json (if not it did return as text)
	w.Header().Add("Content-Type", "application/json")

	//encode the slice of customers into json format (pass IO writer - w)
	json.NewEncoder(w).Encode(customers)
}
