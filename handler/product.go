package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Product {
	return &Product{l}
}

func (p *Product) serveHttp(w http.ResponseWriter, r *http.Request) {

	//Encoding to json

	lp := data.getProduct()
	d, err := json.Marshal(lp)
	if err != nil {
		http.Error(w, "Unable to encode to jason", http.StatusInternalServerError)
	}

	// send back with respone
	w.Write(d)

}