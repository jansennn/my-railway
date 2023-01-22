package service

import (
	"fmt"
	"log"

	"github.com/mashingan/smapping"
	"my-railway/dto"
	"my-railway/entity"
	"my-railway/repo"

	_produk "my-railway/service/produk"
)

type ProdukService interface {
	All() (*[]_produk.ProdukResponse, error)
	CreateProduk(produkRequest dto.CreateProdukRequest) (*_produk.ProdukResponse, error)
	FindOneProdukById(produkID string) (*_produk.ProdukResponse, error)
	FindProdukByKategori(kategori string) (*[]_produk.ProdukResponse, error)
}

type produkService struct {
	produkRepo repo.ProdukRepository
}

func NewProdukService(produkRepo repo.ProdukRepository) ProdukService {
	return &produkService{
		produkRepo: produkRepo,
	}
}

func (c *produkService) All() (*[]_produk.ProdukResponse, error) {
	produks, err := c.produkRepo.All()
	if err != nil {
		return nil, err
	}

	prods := _produk.NewProdukArrayResponse(produks)
	return &prods, nil
}

func (c *produkService) CreateProduk(produkRequest dto.CreateProdukRequest) (*_produk.ProdukResponse, error) {
	produk := entity.Produk{}
	err := smapping.FillStruct(&produk, smapping.MapFields(&produkRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	p, err := c.produkRepo.InsertProduk(produk)
	if err != nil {
		return nil, err
	}

	res := _produk.NewProdukInsertResponse(p)
	return &res, nil
}

func (c *produkService) FindOneProdukById(produkID string) (*_produk.ProdukResponse, error) {
	produk, err := c.produkRepo.FindOneProdukById(produkID)
	fmt.Print("produk id :", produkID)
	if err != nil {
		return nil, err
	}

	res := _produk.NewProdukResponse(produk)
	return &res, nil
}

func (c *produkService) FindProdukByKategori(kategori string) (*[]_produk.ProdukResponse, error) {
	produk, err := c.produkRepo.FindProdukByKategori(kategori)
	if err != nil {
		return nil, err
	}

	res := _produk.NewProdukArrayResponse(produk)
	return &res, nil
}