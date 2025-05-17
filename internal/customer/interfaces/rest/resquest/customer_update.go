package resquest

type UpdateCustomerRequest struct {
	Name        string `json:"name" `
	PhoneNumber string `json:"phone_number" `
	NationalID  string `json:"national_id" `
	Address     string `json:"address" `
	RegionID    string `json:"region_id" `
}
