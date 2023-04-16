package dto

import "time"

type NewProductRequest struct {
	Title string `json:"title" valid:"required~title cannot be empty" example:"buku dongeng"`
	Price int    `json:"price" valid:"required~price cannot be empty" example:"20000"`
}

type NewProductResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

type GetProductByIdResponse struct {
	Result     string `json:"result"`
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
	Data       ProductResponse
}

type GetProductsResponse struct {
	Result     string            `json:"result"`
	Message    string            `json:"message"`
	StatusCode int               `json:"statusCode"`
	Data       []ProductResponse `json:"data"`
}

type ProductResponse struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	UserId    int       `json:"userId"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
