package models

type Menu struct {
	Id       int64   `gorm:"primaryKey" json:"id"`
	NamaMenu string  `gorm:"type:varchar(300)" json:"menu"`
	Harga    float64 `gorm:"type:decimal(14.2)" json:"harga"`
	Stok     int32   `gorm:"type:int(5)" json:"stok"`
}
