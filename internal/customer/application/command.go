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

type UpdateCustomerInfoCommand struct {
	CustomerID  domain.CustomerID
	Name        string
	PhoneNumber string
	Address     string
	NationalID  domain.NationalID
	RegionID    domain.RegionID
}

type CreateCustomerCommand struct {
	Name        string
	PhoneNumber string
	NationalID  domain.NationalID
	Address     string
	RegionID    domain.RegionID
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

func (s *CustomerService) UpdateCustomer(ctx context.Context, cmd UpdateCustomerInfoCommand) error {
	customer, err := s.repo.FindByID(ctx, cmd.CustomerID)
	if err != nil {
		return err
	}

	if cmd.Name != "" {
		err = customer.UpdateName(cmd.Name)
		if err != nil {
			return err
		}
	}

	if cmd.NationalID != "" {
		err = customer.UpdateNationalID(cmd.NationalID)
		if err != nil {
			return err
		}
	}

	if cmd.PhoneNumber != "" {
		err = customer.UpdatePhoneNumber(cmd.PhoneNumber)
		if err != nil {
			return err
		}
	}

	if cmd.Address != "" {
		err = customer.UpdateAddress(cmd.Address)
		if err != nil {
			return err
		}
	}

	if cmd.RegionID != "" {
		err = customer.UpdateRegion(cmd.RegionID)
		if err != nil {
			return err
		}
	}

	return s.repo.Save(ctx, customer)
}
