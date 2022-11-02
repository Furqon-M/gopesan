package models

type Pesandb struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	Nopesanan string  `gorm:"type:varchar(300)" json:"nopesanan"`
	Pesanan   string  `gorm:"type:varchar(300)" json:"pesanan"`
	Harga     float64 `gorm:"type:decimal(14.2)" json:"harga"`
	Jumlah    int     `gorm:"type:int(5)" json:"jumlah"`
	Total     float64 `gorm:"type:decimal(14.2)" json:"total"`
	Meja      int     `gorm:"type:int(5)" json:"meja"`
	Status    string  `gorm:"type:varchar(300)" json:"status"`
	Stok      int32   `gorm:"type:int(5)" json:"stok"`
}
