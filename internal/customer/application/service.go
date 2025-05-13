package application

import "github.com/Sadegh-kh/telemon/internal/customer/domain"

type CustomerService struct {
	repo domain.CustomerRepository
}

func NewCustomerService(repo domain.CustomerRepository) *CustomerService {
	return &CustomerService{
		repo: repo,
	}
}
