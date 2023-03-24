package handler

import (
	"example/learn0/data"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)

		// read more about regexp - https://cs.opensource.google/go/go/+/go1.20.2:src/regexp/regexp.go;l=1197
		//expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id ")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group ")
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]

		id, err := strconv.Atoi(idString) // convert string id to integer

		if err != nil {
			p.l.Println("Invalid URI unable to convert to number ", idString)
			http.Error(w, "Invalid URI", http.StatusBadRequest)
			return
		}

		// p.l.Println("got id := ", id)

		p.updateProduct(id, w, r)
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

func (p Product) updateProduct(id int, w http.ResponseWriter, r *http.Request) {

	p.l.Println("Handle PUT Products")
	pdt := &data.Product{}

	err := pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
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
