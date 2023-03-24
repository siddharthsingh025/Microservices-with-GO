package handler

import (
	"example/learn0/data"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

// for GET req
func (p *Product) GetProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetching data from DataStore
	lp := data.GetProduct()
	// d, err := json.Marshal(lp)
	err := lp.ToJson(w)

	if err != nil {
		http.Error(w, "Unable to encode to jason", http.StatusInternalServerError)
	}
}

// for POST req
func (p *Product) AddProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST Products")
	pdt := &data.Product{}
	err := pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
	}

	p.l.Printf("product : %#v", pdt) // it will print decoded data into nice format in logWindow

	data.AddProduct(pdt) // call func to add decoded data in list
}

// for PUT req
func (p Product) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	// mux extract id from URL using Vars()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "enable to convert id", http.StatusBadRequest)
	}

	p.l.Println("Handle PUT Products", id)
	pdt := &data.Product{}

	err = pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
	}

	err = data.UpdateProduct(id, pdt)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}
