package dto


type CreateImageRequest struct {
	Path  		string `json:"path" form:"path" binding:"required,min=1"`
	ProdukID 	int64 `json:"produk_id" form:"produk_id" binding:"required"`
}

type UpdateImageRequest struct {
	ID    		int64  `json:"id" form:"id"`
	Path  		string `json:"path" form:"path" binding:"required,min=1"`
	ProdukID 	int64 `json:"produk_id" form:"produk_id" binding:"required"`
}
