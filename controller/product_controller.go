package controller

import (
	"go-api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type productController struct {
	//Usecase
}

func NewProductController() productController {
	return productController{}

}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID: 1,
			Name: ""
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, product)
}