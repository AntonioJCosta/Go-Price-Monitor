package models

type Product struct {
	// Auto increment ID
	ID        uint `gorm:"primary_key;auto_increment;not_null"`
	Name      string
	FullPrice float64
	Link      string
}
