package main

import (
	"example/learn0/handler"
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hl := handler.NewHello(l) // that will creat Hello object with l loger
	//hl is our handler object with servehttp func act as handlerfunc

	//here we define new serverMUX and than register our above handler to it.
	sm := http.NewServeMux()
	sm.Handle("/", hl)

	http.ListenAndServe(":9090", sm)
}
