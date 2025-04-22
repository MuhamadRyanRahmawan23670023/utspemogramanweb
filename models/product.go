package models

type Produk struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	Nama       string   `json:"nama"`
	Harga      float64  `json:"harga"`
	KategoriID uint     `json:"kategori_id"`           // foreign key
	Kategori   Kategori `gorm:"foreignKey:KategoriID"` // relasi ke kategori
}
