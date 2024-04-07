package app

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	"github.com/obynonwane/banking/service"
)

// concrete implementaion of customer service
type CustomerHandlers struct {
	service service.CustomerService
}
type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zip_code" xml:"zip_code"`
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		//make the response header to be json (if not it did return as text)
		w.Header().Add("Content-Type", "application/xml")

		//encode the slice of customers into json format (pass IO writer - w)
		xml.NewEncoder(w).Encode(customers)
	} else {
		//make the response header to be json (if not it did return as text)
		w.Header().Add("Content-Type", "application/json")

		//encode the slice of customers into json format (pass IO writer - w)
		json.NewEncoder(w).Encode(customers)
	}

}
