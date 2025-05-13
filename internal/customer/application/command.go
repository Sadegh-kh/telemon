package application

import (
	"context"

	"github.com/Sadegh-kh/telemon/internal/customer/domain"
)

type UpdatePhoneNumberCommand struct {
	CustomerID  domain.CustomerID
	PhoneNumber string
}

type UpdateAddressCommand struct {
	CustomerID domain.CustomerID
	Address    string
}

func (s *CustomerService) UpdatePhoneNumber(ctx context.Context, cmd UpdatePhoneNumberCommand) error {
	customer, err := s.repo.FindByID(ctx, cmd.CustomerID)
	if err != nil {
		return err
	}
	customer.UpdatePhoneNumber(cmd.PhoneNumber)
	return s.repo.Save(ctx, customer)
}

func (s *CustomerService) UpdateAddress(ctx context.Context, cmd UpdateAddressCommand) error {
	customer, err := s.repo.FindByID(ctx, cmd.CustomerID)
	if err != nil {
		return err
	}
	customer.UpdateAddress(cmd.Address)
	return s.repo.Save(ctx, customer)
}
