package dto

import "time"

type CustomerDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Address     string    `json:"address"`
	NationalID  string    `json:"national_id"`
	RegionID    string    `json:"region_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
