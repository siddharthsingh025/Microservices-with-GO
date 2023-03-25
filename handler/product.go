package handler

import (
	"context"
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
	pdt := r.Context().Value(KeyProduct{}).(data.Product) //getting product from context

	p.l.Printf("product : %#v", pdt) // it will print decoded data into nice format in logWindow

	data.AddProduct(&pdt) // call func to add decoded data in list
}

// for PUT req
func (p Product) UpdateProducts(w http.ResponseWriter, r *http.Request) {
	// mux extract id from URL using Vars()
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "enable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT Products", id)
	pdt := r.Context().Value(KeyProduct{}).(data.Product) // getting our decoded json data from request context and cast it into Product type

	err = data.UpdateProduct(id, &pdt)
	if err == data.ErrProductNotFound {
		http.Error(w, "Product not found", http.StatusBadRequest)
		return
	}

	if err != nil {
		http.Error(w, "Product not found", http.StatusInternalServerError)
		return
	}
}

//creating MiddleWare for validation our request

type KeyProduct struct{}

func (p Product) MiddlewxareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		pdt := data.Product{}

		err := pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
		if err != nil {
			p.l.Println("[ERROR] deseializing product", err)
			http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
			return
		}

		//add the product (pdt) to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, pdt)
		req := r.WithContext(ctx)

		//call the handler , which can be another middleware in the chain , or the final handler
		next.ServeHTTP(w, req)
	})
}
