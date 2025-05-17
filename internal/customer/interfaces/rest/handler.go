package rest

import (
	"github.com/Sadegh-kh/telemon/internal/customer/application/command"
	"github.com/Sadegh-kh/telemon/internal/customer/application/query"
)

type CustomerHandler struct {
	commandService *command.CustomerCommandService
	queryService   *query.CusotmerQueryService
}

func NewCustomerHandler(cmd *command.CustomerCommandService, qry *query.CusotmerQueryService) *CustomerHandler {
	return &CustomerHandler{
		commandService: cmd,
		queryService:   qry,
	}
}
