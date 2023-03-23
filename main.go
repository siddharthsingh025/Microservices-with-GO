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
)

func main() {

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handler.NewProduct(l) // that will creat Product handler object with l loger
	//ph is our handler object with servehttp func act as handlerfunc

	//here we define new serverMUX and than register our above handler to it.
	sm := http.NewServeMux()
	sm.Handle("/", ph)

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
