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

	if r.Method == http.MethodPost {
		p.addProducts(w, r)
		return
	}

	//otherwise//
	w.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Product) getProducts(w http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	//fetching data from DataStore
	lp := data.GetProduct()
	// d, err := json.Marshal(lp)
	err := lp.ToJson(w)

	if err != nil {
		http.Error(w, "Unable to encode to jason", http.StatusInternalServerError)
	}
}

func (p *Product) addProducts(w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle POST Products")
	pdt := &data.Product{}
	err := pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
	if err != nil {
		http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
	}

	p.l.Printf("product : %#v", pdt) // it will print decoded data into nice format in logWindow

	data.AddProduct(pdt) // call func to add decoded data in list
}
