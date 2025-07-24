// models/supplier.go
package models

type Supplier struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Logo         string    `json:"logo"`
	SupplierName string    `json:"supplier_name"`
	NickName     string    `json:"nick_name"`
	Address      string    `json:"address"`
	Status       string    `json:"status"`
	Contacts     []Contact `gorm:"foreignKey:SupplierID" json:"contacts"`
}
