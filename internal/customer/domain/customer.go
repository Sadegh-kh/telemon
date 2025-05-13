package domain

import "errors"

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
	CreatedAt   string
}

func (c *Customer) UpdatePhoneNumber(newPhoneNumber string) error {
	if newPhoneNumber == "" {
		return errors.New("phone number cannot be empty")
	}
	c.PhoneNumber = newPhoneNumber
	return nil
}

func (c *Customer) UpdateAddress(newAddress string) error {
	if newAddress == "" {
		return errors.New("address cannot be empty")
	}
	c.Address = newAddress
	return nil
}

func (c *Customer) UpdateRegion(newRegionID RegionID) error {
	if newRegionID == "" {
		return errors.New("region ID cannot be empty")
	}
	c.RegionID = newRegionID
	return nil
}
