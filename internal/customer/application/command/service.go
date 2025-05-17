package command

import (
	"github.com/Sadegh-kh/telemon/internal/customer/domain"
	"github.com/Sadegh-kh/telemon/internal/shared/idgenerator"
)

type CustomerCommandService struct {
	repo        domain.CustomerRepository
	idGenerator idgenerator.IDGenerator
}

func NewCustomerService(repo domain.CustomerRepository, idGen idgenerator.IDGenerator) *CustomerCommandService {
	return &CustomerCommandService{
		repo:        repo,
		idGenerator: idGen,
	}
}
