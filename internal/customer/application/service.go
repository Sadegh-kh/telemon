package application

import (
	"github.com/Sadegh-kh/telemon/internal/customer/domain"
	"github.com/Sadegh-kh/telemon/internal/shared/idgenerator"
)

type CustomerService struct {
	repo        domain.CustomerRepository
	idGenerator idgenerator.IDGenerator
}

func NewCustomerService(repo domain.CustomerRepository, idGen idgenerator.IDGenerator) *CustomerService {
	return &CustomerService{
		repo:        repo,
		idGenerator: idGen,
	}
}
