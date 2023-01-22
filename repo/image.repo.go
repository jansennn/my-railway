package repo

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"my-railway/entity"
)

type ImageRepository interface {
	All() ([]entity.Image, error)
	InsertImage(image entity.Image) (entity.Image, error)
}

type imageRepo struct {
	connection *gorm.DB
}

func NewImageRepo(connection *gorm.DB) ImageRepository{
	return &imageRepo{
		connection: connection,
	}
}

func (c *imageRepo) All() ([]entity.Image, error) {
	images := []entity.Image{}
	c.connection.Preload(clause.Associations).Find(&images)
	return images, nil
}

func (c *imageRepo) InsertImage(image entity.Image) (entity.Image, error) {
	c.connection.Save(&image)
	c.connection.Find(&image)
	return image, nil
}