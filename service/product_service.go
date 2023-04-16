package service

import (
	"go-jwt/dto"
	"go-jwt/entity"
	"go-jwt/package/errs"
	"go-jwt/package/helpers"
	"go-jwt/repository/product_repository"
	"net/http"
)

type ProductService interface {
	CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr)
	// GetProductById
	GetProductById(productId int) (*dto.GetProductByIdResponse, errs.MessageErr)
	GetProduct() (*dto.GetProductsResponse, errs.MessageErr)
}

type productService struct {
	productRepo product_repository.ProductRepository
}

func NewProductService(productRepo product_repository.ProductRepository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (m *productService) UpdateProductById(productId int, productRequest dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {

	err := helpers.ValidateStruct(productRequest)

	if err != nil {
		return nil, err
	}

	payload := entity.Product{
		Id:    productId,
		Title: productRequest.Title,
		Price: productRequest.Price,
	}

	err = m.productRepo.UpdateProductById(payload)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusOK,
		Result:     "success",
		Message:    "your Product data successfully updated",
	}

	return &response, nil
}

func (m *productService) CreateProduct(userId int, payload dto.NewProductRequest) (*dto.NewProductResponse, errs.MessageErr) {
	productRequest := &entity.Product{
		Title:  payload.Title,
		Price:  payload.Price,
		UserId: userId,
	}

	_, err := m.productRepo.CreateProduct(productRequest)

	if err != nil {
		return nil, err
	}

	response := dto.NewProductResponse{
		StatusCode: http.StatusCreated,
		Result:     "success",
		Message:    "congratulation your product data has been successfully created",
	}

	return &response, err
}
func (p *productService) GetProduct() (*dto.GetProductsResponse, errs.MessageErr) {
	products, err := p.productRepo.GetProduct()

	if err != nil {
		return nil, err
	}

	var productsResponse []dto.ProductResponse

	for _, eachProdcuts := range products {
		productsResponse = append(productsResponse, eachProdcuts.EntityToProductResponseDto())
	}

	result := dto.GetProductsResponse{
		Result:     "success",
		Message:    "product data successfully sent",
		Data:       productsResponse,
		StatusCode: http.StatusOK,
	}

	return &result, nil
}

// func (m *productService) GetProduct() (*dto.GetProductsResponse, errs.MessageErr) {
// 	products, err := m.productRepo.GetProduct()

// 	if err != nil {
// 		return nil, err
// 	}

// 	productResponse := []dto.ProductResponse{}

// 	for _, eachProduct := range products {
// 		productResponse = append(productResponse, eachProduct.EntityToProductResponseDto())
// 	}

// 	response := dto.GetProductsResponse{
// 		Result:     "success",
// 		StatusCode: http.StatusOK,
// 		Message:    "product data have been sent successfully",
// 		Data:       productResponse,
// 	}

// 	return &response, nil
// }

func (m *productService) GetProductById(productId int) (*dto.GetProductByIdResponse, errs.MessageErr) {
	result, err := m.productRepo.GetProductById(productId)

	if err != nil {
		return nil, err
	}

	response := dto.GetProductByIdResponse{
		Result:     "success",
		StatusCode: http.StatusOK,
		Message:    "product data have been sent successfully",
		Data: dto.ProductResponse{
			Id:        result.Id,
			Title:     result.Title,
			Price:     result.Price,
			UserId:    result.UserId,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
		},
	}

	return &response, nil
}
