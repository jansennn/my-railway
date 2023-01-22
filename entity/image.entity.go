package entity

type Image struct {
	ID            int64  `gorm:"primary_key:auto_increment" json:"-"`
	Path   		  string `gorm:"type:text" json:"-"`
	ProdukID	  int64  `gorm:"column:produk_id"`
	Produk		  Produk
}