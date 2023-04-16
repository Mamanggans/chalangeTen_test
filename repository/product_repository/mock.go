package product_repository

import (
	"go-jwt/entity"
	"go-jwt/package/errs"
)

var (
	CreateProduct     func(productPayload *entity.Product) (*entity.Product, errs.MessageErr)
	GetProductById    func(productId int) (*entity.Product, errs.MessageErr)
	UpdateProductById func(payload entity.Product) errs.MessageErr
	GetProducts       func() ([]*entity.Product, errs.MessageErr)
)

type productRepoMock struct{}

func NewProductRepoMock() ProductRepository {
	return &productRepoMock{}
}

func (m *productRepoMock) GetProduct() ([]*entity.Product, errs.MessageErr) {
	return GetProducts()
}

func (m *productRepoMock) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	return CreateProduct(productPayload)
}
func (m *productRepoMock) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	return GetProductById(productId)
}
func (m *productRepoMock) UpdateProductById(payload entity.Product) errs.MessageErr {
	return UpdateProductById(payload)
}
