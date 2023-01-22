package entity

type Produk struct {
	ID     		int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name   		string `gorm:"type:varchar(100)" json:"-"`
	Kategori 	int64
	Price  		string `gorm:"type:varchar(100)" json:"-"`
	Images 		[]Image `gorm:"ForeignKey:ProdukID"`
}
