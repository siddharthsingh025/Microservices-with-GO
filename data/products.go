package data

import "time"

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float32
	SKU         string
	CreatedOn   string
	UpdatedOn   string
	DeletedOn   string
}

// it will return list of all product
func getProduct() []*Product {
	return productList
}

// static list of products
var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "siddharth",
		Description: "strong black milky coffee",
		Price:       3.67,
		SKU:         "abd234",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "divu",
		Description: "light and  black milky coffee",
		Price:       1.99,
		SKU:         "xyz544",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}
