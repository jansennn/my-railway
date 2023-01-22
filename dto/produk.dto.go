package dto

type CreateProdukRequest struct {
	Name  		string `json:"name" form:"name" binding:"required,min=1"`
	Kategori    int64  `json:"kategori" form:"kategori"`
	Price 		string `json:"price" form:"email" binding:"required"`
}

type UpdateProdukRequest struct {
	ID    		int64  `json:"id" form:"id"`
	Name 		string `json:"name" form:"name" binding:"required,min=1"`
	Kategori    int64  `json:"kategori" form:"kategori"`
	Price 		string `json:"price" form:"email" binding:"required"`
}
