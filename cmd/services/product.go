package services

import (
	"context"
	"grpc_go/pb/pagination"
	productPb "grpc_go/pb/product"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type ProductService struct {
	productPb.UnimplementedProductServiceServer
	*gorm.DB
}

func (p *ProductService) GetProducts(context.Context, *productPb.Empty) (*productPb.Products, error) {
	var products []*productPb.Product

	rows, err := p.DB.Table("products AS p").
		Joins("LEFT JOIN categories AS c ON c.id = p.category_id").
		Select("p.id", "p.name", "p.price", "p.stock", "c.id as category_id", "c.name as category_name").
		Rows()

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var product productPb.Product
		var category productPb.Category

		if err := rows.Scan(&product.Id, &product.Name, &product.Price,
			&product.Stock, &category.Id, &category.Name); err != nil {
			log.Fatalf("Row data failed %v", err.Error())
		}
		product.Category = &category
		products = append(products, &product)
	}

	response := &productPb.Products{
		Pagination: &pagination.Pagination{
			Total:       2,
			PerPage:     1,
			CurrentPage: 1,
			LastPage:    1,
		},
		Data: products,
	}

	return response, nil
}
