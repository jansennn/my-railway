package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ratu-melamine-be/common/obj"
	"ratu-melamine-be/common/response"
	"ratu-melamine-be/dto"
	"ratu-melamine-be/service"
)

type ImageHandler interface {
	All(ctx *gin.Context)
	CreateImage(ctx *gin.Context)
}

type imageHandler struct {
	imageService service.ImageService

}

func NewImageHandler(imageService service.ImageService) ImageHandler {
	return &imageHandler{
		imageService: imageService,
	}
}

func (c *imageHandler) All(ctx *gin.Context) {
	images, err := c.imageService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", images)
	ctx.JSON(http.StatusOK, response)
}


func (c *imageHandler) CreateImage(ctx *gin.Context) {
	var createImageReq dto.CreateImageRequest
	err := ctx.ShouldBind(&createImageReq)

	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.imageService.CreateImage(createImageReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}