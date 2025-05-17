package resquest

type CreateCustomerRequest struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	NationalID  string `json:"national_id" validate:"required"`
	Address     string `json:"address" validate:"required"`
	RegionID    string `json:"region_id" validate:"required"`
}
