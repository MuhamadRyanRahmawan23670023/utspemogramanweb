package models

type Kategori struct {
	ID     uint     `json:"id" gorm:"primaryKey"`
	Nama   string   `json:"nama"`
	Produk []Produk `gorm:"foreignKey:KategoriID"` // relasi dengan produk
}
