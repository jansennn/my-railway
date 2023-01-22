package repo

import (
	"gorm.io/gorm/clause"
	"my-railway/entity"
	"gorm.io/gorm"
)

type ProdukRepository interface {
	All() ([]entity.Produk, error)
	InsertProduk(produk entity.Produk) (entity.Produk, error)
	FindOneProdukById(ID string) (entity.Produk, error)
	FindProdukByKategori(kategori string) ([]entity.Produk, error)
}

type produkRepo struct {
	connection *gorm.DB
}

func NewProdukRepo(connection *gorm.DB) ProdukRepository{
	return &produkRepo{
		connection: connection,
	}
}

func (c *produkRepo) All() ([]entity.Produk, error) {
	produks := []entity.Produk{}
	c.connection.Preload(clause.Associations).Find(&produks)
	return produks, nil
}

func (c *produkRepo) InsertProduk(produk entity.Produk) (entity.Produk, error) {
	c.connection.Save(&produk)
	c.connection.Find(&produk)
	return produk, nil
}

func (c *produkRepo) FindOneProdukById(produkID string) (entity.Produk, error) {
	produks := entity.Produk{}
	c.connection.Preload(clause.Associations).Where("id = ?", produkID).Find(&produks)
	return produks, nil
}

func (c *produkRepo) FindProdukByKategori(kategori string) ([]entity.Produk, error) {
	produks := []entity.Produk{}
	c.connection.Preload(clause.Associations).Where("kategori = ?", kategori).Find(&produks)
	return produks, nil
}