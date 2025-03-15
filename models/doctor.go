package models

type Doctor struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phoneNumber"`
	PoliType    string `json:"poliType"`
}
