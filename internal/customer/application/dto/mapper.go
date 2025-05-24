package dto

import (
	"github.com/Sadegh-kh/telemon/internal/customer/domain"
)

func ToCustomerDTO(c *domain.Customer) *CustomerDTO {
	return &CustomerDTO{
		ID:          string(c.ID),
		Name:        c.Name,
		PhoneNumber: c.PhoneNumber,
		Address:     c.Address,
		NationalID:  string(c.NationalID),
		RegionID:    string(c.RegionID),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func FromCustomerDTO(dto *CustomerDTO) *domain.Customer {
	return &domain.Customer{
		ID:          domain.CustomerID(dto.ID),
		Name:        dto.Name,
		PhoneNumber: dto.PhoneNumber,
		Address:     dto.Address,
		NationalID:  domain.NationalID(dto.NationalID),
		RegionID:    domain.RegionID(dto.RegionID),
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}
}
