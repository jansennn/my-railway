package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ratu-melamine-be/common/obj"
	"ratu-melamine-be/common/response"
	"ratu-melamine-be/dto"
	"ratu-melamine-be/service"
	"net/http"
)


type ProdukHandler interface {
	All(ctx *gin.Context)
	CreateProduk(ctx *gin.Context)
	FindOneProdukById(ctx *gin.Context)
	FindProdukByKategori(ctx *gin.Context)
}

type produkHandler struct {
	produkService service.ProdukService

}

func NewProdukHandler(produkService service.ProdukService) ProdukHandler {
	return &produkHandler{
		produkService: produkService,
	}
}

func (c *produkHandler) All(ctx *gin.Context) {
	produks, err := c.produkService.All()
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", produks)
	ctx.JSON(http.StatusOK, response)
}


func (c *produkHandler) CreateProduk(ctx *gin.Context) {
	var createProdukReq dto.CreateProdukRequest
	err := ctx.ShouldBind(&createProdukReq)

	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := c.produkService.CreateProduk(createProdukReq)
	if err != nil {
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusCreated, response)
}

func (c *produkHandler) FindOneProdukById(ctx *gin.Context) {
	id := ctx.Param("id")
	fmt.Print("produk id :", id)
	res, err := c.produkService.FindOneProdukById(id)
	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}

func (c *produkHandler) FindProdukByKategori(ctx *gin.Context) {
	kategori := ctx.Param("kategori")
	fmt.Print("kategori : ", kategori)
	res, err := c.produkService.FindProdukByKategori(kategori)
	if err != nil{
		response := response.BuildErrorResponse("Failed to process request", err.Error(), obj.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := response.BuildResponse(true, "OK!", res)
	ctx.JSON(http.StatusOK, response)
}
