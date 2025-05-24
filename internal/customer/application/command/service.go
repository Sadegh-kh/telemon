package command

import (
	"github.com/Sadegh-kh/telemon/internal/customer/domain"
	"github.com/Sadegh-kh/telemon/internal/shared/idgenerator"
)

type CustomerCommandService struct {
	cmdRepo     domain.CustomerCommandRepository
	queryRepo   domain.CustomerQueryRepository
	idGenerator idgenerator.IDGenerator
}

func NewCustomerService(cmdRepo domain.CustomerCommandRepository, queryRepo domain.CustomerQueryRepository, idGen idgenerator.IDGenerator) *CustomerCommandService {
	return &CustomerCommandService{
		cmdRepo:     cmdRepo,
		queryRepo:   queryRepo,
		idGenerator: idGen,
	}
}
