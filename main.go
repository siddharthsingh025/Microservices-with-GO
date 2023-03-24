package main

import (
	"context"
	"example/learn0/handler"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handler.NewProduct(l) // that will creat Product handler object with l loger
	//ph is our handler object with servehttp func act as handlerfunc

	//here we define new Router using gorilla mux
	sm := mux.NewRouter()

	//GET request
	getRouter := sm.Methods(http.MethodGet).Subrouter() // methods registers a new route with matcher for HTTP methods
	getRouter.HandleFunc("/", ph.GetProducts)

	//PUT request
	getRouter = sm.Methods(http.MethodPut).Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	//this is how we define above is regexp for extrcting id form URI

	//POST request
	getRouter = sm.Methods(http.MethodPost).Subrouter()
	getRouter.HandleFunc("/", ph.AddProducts)
	//this is how we define above is regexp for extrcting id form URI

	// sm.Handle("/", ph)

	//creating our own server , for tutining our application for better performance
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm, // our servermux
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//handle my listenAndServe so that it not gonna block
	go func() { // calls goroutines
		fmt.Println("\n Starting The Server ....\n")
		time.Sleep(2 * time.Second)
		fmt.Println("\n $ ACTIVATED $\n")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}() //its a function call

	sigChan := make(chan os.Signal)

	// signal.notify is going to broadcast a message on this channel ( sigchan ) whenever  an operating system
	// kill command or os intrupts is receiveded
	// obviouslly we block here to consume the msg from the  channel
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("Recieved terminate, graceful shutdown", sig)

	// shuting down server , by creating  a context with timedelay of 30s wait for all previous task to comlete and than go to shutdown
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

	//default server by http
	//http.ListenAndServe(":9090", sm)
}
