package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// WAy of creating handler thats help in dependencies injection

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello { // creat new hello object of class Hellow with log l
	return &Hello{l} // and return it , that object containg ServeHttp method that we build
	// for handling
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Print("Hellow world")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Opss", http.StatusBadRequest)
		//above line can also  be replace with these 2 lines :
		// w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte("Opss"))
		return
	}

	fmt.Fprintf(w, "hellow %s\n", d)
}
