package data

import "time"

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// it will return list of all product
func GetProduct() []*Product {
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
