package _produk

import (
	"my-railway/entity"
	_image "my-railway/service/image"
)

type ProdukResponse struct {
	ID          int64              		`json:"id"`
	Name 		string             		`json:"name"`
	Kategori	int64					`json:"kategori"`
	Price       string             		`json:"price"`
	Image       []_image.ImageResponse 	`json:"image,omitempty"`
}

func NewProdukResponse(produk entity.Produk) ProdukResponse {
	return ProdukResponse{
		ID:          produk.ID,
		Name: 		 produk.Name,
		Kategori: 	 produk.Kategori,
		Price:       produk.Price,
		Image:       _image.NewImageArrayResponse(produk.Images),
	}
}

func NewProdukInsertResponse(produk entity.Produk) ProdukResponse {
	return ProdukResponse{
		ID:          produk.ID,
		Name: 		 produk.Name,
		Kategori: 	 produk.Kategori,
		Price:       produk.Price,
	}
}

func NewProdukArrayResponse(produks []entity.Produk) []ProdukResponse {
	productRes := []ProdukResponse{}
	for _, v := range produks {
		p := ProdukResponse{
			ID    		: v.ID,
			Name  		: v.Name,
			Kategori	: v.Kategori,
			Price 		: v.Price,
			Image 		: _image.NewImageArrayResponse(v.Images),
		}
		productRes = append(productRes, p)
	}
	return productRes
}
