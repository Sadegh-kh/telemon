package query

import (
	"context"

	"github.com/Sadegh-kh/telemon/internal/customer/application/dto"
	"github.com/Sadegh-kh/telemon/internal/customer/domain"
)

func (c *CusotmerQueryService) GetCustomerByID(ctx context.Context, id domain.CustomerID) (*dto.CustomerDTO, error) {
	customer, err := c.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	dto := dto.ToCustomerDTO(customer)

	return dto, nil
}

func (c *CusotmerQueryService) GetCustomerByNatinalID(ctx context.Context, nationalID domain.NationalID) (*dto.CustomerDTO, error) {
	customer, err := c.repo.FindByNationalID(ctx, nationalID)
	if err != nil {
		return nil, err
	}

	dto := dto.ToCustomerDTO(customer)

	return dto, nil
}
