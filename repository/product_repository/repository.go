package product_repository

import (
	"go-jwt/entity"
	"go-jwt/package/errs"
)

type ProductRepository interface {
	CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById(productId int) (*entity.Product, errs.MessageErr)
	DeleteProduct(product *entity.Product) (*entity.Product, errs.MessageErr)
	UpdateProductById(payload entity.Product) errs.MessageErr
	GetProduct() ([]*entity.Product, errs.MessageErr)
}
