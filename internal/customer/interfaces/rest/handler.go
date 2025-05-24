package rest

import (
	"github.com/Sadegh-kh/telemon/internal/customer/application/command"
	"github.com/Sadegh-kh/telemon/internal/customer/application/query"
)

type CustomerCommandHandler struct {
	commandService *command.CustomerCommandService
}

type CustomerQueryHandler struct {
	queryService *query.CusotmerQueryService
}

func NewCustomerCommandHandler(cmd *command.CustomerCommandService) *CustomerCommandHandler {
	return &CustomerCommandHandler{
		commandService: cmd,
	}
}

func NewCustomerQueryHandler(qry *query.CusotmerQueryService) *CustomerQueryHandler {
	return &CustomerQueryHandler{
		queryService: qry,
	}
}
