package domain

import (
	"errors"
	"time"
)

type (
	CustomerID string
	NationalID string
	RegionID   string
)

type Customer struct {
	ID          CustomerID
	Name        string
	PhoneNumber string
	NationalID  NationalID
	Address     string
	RegionID    RegionID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewCustomer(id CustomerID, name, address, phoneNumber string, nationalID NationalID, regionID RegionID) (*Customer, error) {
	return &Customer{
		ID:          id,
		Name:        name,
		PhoneNumber: phoneNumber,
		NationalID:  nationalID,
		Address:     address,
		RegionID:    regionID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (c *Customer) UpdatePhoneNumber(newPhoneNumber string) error {
	if newPhoneNumber == "" {
		return errors.New("phone number cannot be empty")
	}
	c.PhoneNumber = newPhoneNumber
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Customer) UpdateAddress(newAddress string) error {
	if newAddress == "" {
		return errors.New("address cannot be empty")
	}
	c.Address = newAddress
	c.UpdatedAt = time.Now()
	return nil
}

func (c *Customer) UpdateRegion(newRegionID RegionID) error {
	if newRegionID == "" {
		return errors.New("region ID cannot be empty")
	}
	c.RegionID = newRegionID
	c.UpdatedAt = time.Now()
	return nil
}
