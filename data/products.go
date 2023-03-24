package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` // this anotation use to avoid to add this into output
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

type Products []*Product

func (p *Products) ToJson(w io.Writer) error {
	//defining encoder
	e := json.NewEncoder(w)

	//pass our product list to encoder to write
	return e.Encode(p)
}

// change GetProduct to return our  custom type of Products
func GetProduct() Products {
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
