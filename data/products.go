package data

import (
	"encoding/json"
	"fmt"
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

//here as a function parameter we get body of post request with our json content
//and we decode it to structure of our Product struct

func (p *Product) FromJson(r io.Reader) error {
	// defining decoder
	e := json.NewDecoder(r)

	// return decoded data to structure of Products type
	return e.Decode(p)
}

type Products []*Product

// ToJson serializes the contents of the collection to JSON
// NewEncoder provides better performance than json.Unmarshal as it does not
// have to buffer the output into an im memory slice of bytes
// this reduces allocations and the overheads of the service
func (p *Products) ToJson(w io.Writer) error {
	//defining encoder
	e := json.NewEncoder(w)

	//pass our product list to encoder to write
	return e.Encode(p)
}

// fumnction to generate integer for our ID
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1 // increment by 1 of Last product Id in the list
}

var ErrProductNotFound = fmt.Errorf("Product not found")

// to find product in our database list by id
func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

// change GetProduct to return our  custom type of Products
func GetProduct() Products {
	return productList
}

// add new product to our static list
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p) // appended our new Product structured data to our existing List
}

func UpdateProduct(id int, p *Product) error {

	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil

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
