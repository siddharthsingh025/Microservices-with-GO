package handler

import (
	"example/learn0/data"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		p.getProducts(w, r)
		return
	}

	//otherwise//
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	//Encoding to json
	lp := data.GetProduct()
	// d, err := json.Marshal(lp)
	err := lp.ToJson(w)

	if err != nil {
		http.Error(w, "Unable to encode to jason", http.StatusInternalServerError)
	}
}
