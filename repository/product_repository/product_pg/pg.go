package product_pg

import (
	"database/sql"
	"errors"
	"fmt"
	"go-jwt/entity"
	"go-jwt/package/errs"
	"go-jwt/repository/product_repository"
)

const (
	getProductByIdQuery = `
		SELECT id, title, userId, price, createdAt, updatedAt from "products"
		WHERE id = $1;
	`

	updateProductByIdQuery = `
		UPDATE "products"
		SET title = $2,
		price = $3
		WHERE id = $1;
	`
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) product_repository.ProductRepository {
	return &productPG{
		db: db,
	}
}

func (m *productPG) UpdateProductById(payload entity.Product) errs.MessageErr {
	_, err := m.db.Exec(updateProductByIdQuery, payload.Id, payload.Title, payload.Price)

	if err != nil {

		return errs.NewInternalServerError("something went wrong")
	}

	return nil
}

func (m *productPG) GetProductById(productId int) (*entity.Product, errs.MessageErr) {
	row := m.db.QueryRow(getProductByIdQuery, productId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.UserId, &product.Price, &product.CreatedAt, &product.UpdatedAt)

	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errs.NewNotFoundError("product not found please check again your product")
		}

		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil
}

func (m *productPG) GetProduct() ([]*entity.Product, errs.MessageErr) {
	return nil, nil
}

func (m *productPG) CreateProduct(productPayload *entity.Product) (*entity.Product, errs.MessageErr) {
	createProductQuery := `
		INSERT INTO "products"
		(
			title,
			price,
			userId
		)
		VALUES($1, $2, $3)
		RETURNING id,title, price, userId;
	`
	row := m.db.QueryRow(createProductQuery, productPayload.Title, productPayload.Price, productPayload.UserId)

	var product entity.Product

	err := row.Scan(&product.Id, &product.Title, &product.Price, &product.UserId)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return nil, errs.NewInternalServerError("something went wrong")
	}

	return &product, nil

}
