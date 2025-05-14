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

type CreateCustomerCommand struct {
	Name        string
	PhoneNumber string
	NationalID  domain.NationalID
	Address     string
	RegionID    domain.RegionID
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

func (s *CustomerService) CreateCustomer(ctx context.Context, cmd CreateCustomerCommand) error {
	customer, err := domain.NewCustomer(
		domain.CustomerID(s.idGenerator.GenerateID()),
		cmd.Name,
		cmd.Address,
		cmd.PhoneNumber,
		cmd.NationalID,
		cmd.RegionID,
	)
	if err != nil {
		return err
	}

	return s.repo.Create(ctx, customer)
}
