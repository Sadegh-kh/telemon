package command

import (
	"context"

	"github.com/Sadegh-kh/telemon/internal/customer/domain"
)

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

func (s *CustomerCommandService) CreateCustomer(ctx context.Context, cmd CreateCustomerCommand) error {
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

	return s.cmdRepo.Create(ctx, customer)
}

func (s *CustomerCommandService) UpdateCustomer(ctx context.Context, cmd UpdateCustomerInfoCommand) error {
	customer, err := s.queryRepo.FindByID(ctx, cmd.CustomerID)
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

	return s.cmdRepo.Save(ctx, customer)
}

func (s *CustomerCommandService) DeleteCustomer(ctx context.Context, id domain.CustomerID) error {

	return s.cmdRepo.Delete(ctx, id)
}
