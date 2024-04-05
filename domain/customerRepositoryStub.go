package domain

// define the stub adaptor
type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {

	customers := []Customer{
		{Id: "1001", Name: "Ahish", City: "New Delhi", Zipcode: "7895759", DateofBirth: "01-01-2020", Status: "Active"},
		{Id: "1002", Name: "Zainab", City: "Lagos", Zipcode: "735474", DateofBirth: "01-01-2010", Status: "Active"},
	}

	return CustomerRepositoryStub{
		customers: customers,
	}
}
