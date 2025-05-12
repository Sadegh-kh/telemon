package domain

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
