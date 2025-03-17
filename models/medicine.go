package models

type Medicine struct {
	ID      uint64 `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Package string `json:"package"`
	Price   string `json:"price"`
}
