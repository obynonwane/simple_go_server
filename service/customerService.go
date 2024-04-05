package service

import "github.com/obynonwane/banking/domain"

type CustomerService interface {
	GetAllCustomers() ([]domain.Customer, error)
}
