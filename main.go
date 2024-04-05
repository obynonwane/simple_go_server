package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func main() {

	//define  routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)
	http.HandleFunc("/customersxml", getAllCustomersXML)

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

func getAllCustomersXML(w http.ResponseWriter, r *http.Request) {

	//create a slice of customers
	customers := []Customer{
		{Name: "Ahish", City: "New Delhi", Zipcode: "7895759"},
		{Name: "Obinna", City: "Logand", Zipcode: "78575"},
	}

	//make the response header to be json (if not it did return as text)
	w.Header().Add("Content-Type", "application/xml")

	//encode the slice of customers into json format (pass IO writer - w)
	xml.NewEncoder(w).Encode(customers)
}
