package image

import (
	"my-railway/entity"
)

type ImageResponse struct {
	ID    		int64  `json:"id"`
	Path  		string `json:"path"`
	ProdukID  	int64 `json:"produk_id"`
}

func NewImageResponse(image entity.Image) ImageResponse {
	return ImageResponse{
		ID:    		image.ID,
		Path:  		image.Path,
		ProdukID: 	image.ProdukID,
	}
}

func NewImageArrayResponse(images []entity.Image) []ImageResponse {
	imageRes := []ImageResponse{}
	for _, v := range images {
		p := ImageResponse{
			ID:          v.ID,
			Path:        v.Path,
			ProdukID:    v.ProdukID,
		}
		imageRes = append(imageRes, p)
	}
	return imageRes
}

