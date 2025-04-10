package models

type Medicine struct {
	ID              uint64  `json:"id" gorm:"primaryKey"`
	MedicineName    string  `json:"medicineName"`
	MedicinePackage string  `json:"medicinePackage"`
	MedicinePrice   float64 `json:"medicinePrice"`
}
