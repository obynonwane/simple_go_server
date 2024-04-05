package domain

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

// define repository
type CustomerRepository interface {
	FindAll() ([]Customer, error)
}
