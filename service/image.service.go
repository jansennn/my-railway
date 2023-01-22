package service

import (
	"github.com/mashingan/smapping"
	"log"
	"ratu-melamine-be/dto"
	"ratu-melamine-be/entity"
	"ratu-melamine-be/repo"
	_image "ratu-melamine-be/service/image"
)

type ImageService interface {
	All() (*[]_image.ImageResponse, error)
	CreateImage(imageRequest dto.CreateImageRequest) (*_image.ImageResponse, error)
}

type imageService struct {
	imageRepo repo.ImageRepository
}

func NewImageService(imageRepo repo.ImageRepository) ImageService {
	return &imageService{
		imageRepo: imageRepo,
	}
}

func (c *imageService) All() (*[]_image.ImageResponse, error) {
	images, err := c.imageRepo.All()
	if err != nil {
		return nil, err
	}

	imags := _image.NewImageArrayResponse(images)
	return &imags, nil
}

func (c *imageService) CreateImage(imageRequest dto.CreateImageRequest) (*_image.ImageResponse, error) {
	image := entity.Image{}
	err := smapping.FillStruct(&image, smapping.MapFields(&imageRequest))

	if err != nil {
		log.Fatalf("Failed map %v", err)
		return nil, err
	}

	p, err := c.imageRepo.InsertImage(image)
	if err != nil {
		return nil, err
	}

	res := _image.NewImageResponse(p)
	return &res, nil
}