package models

type Poli struct {
	ID              uint64 `json:"id" gorm:"primaryKey"`
	PoliName        string `json:"poliName"`
	PoliDescription string `json:"poliDescription"`
}
