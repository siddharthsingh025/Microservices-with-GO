# # Microservices-with-GO
 <p align="center">
   <a>
   <img height="300" width="400" src="https://github.com/siddharthsingh025/Microservices-with-GO/blob/main/imgs/micro.png">
   <img height="300" width="200" src="https://github.com/siddharthsingh025/Microservices-with-GO/blob/main/imgs/golang.png">
   </a>
</p> 

## `#Basics of API develpment with Go [ REST ]`

ListenAndServe - establize http  server with port for serving and handlerFunction to handle coming request 

ServeMux - that register a path to and handler 
   - http package has default handler 
   - HandleFunc ( default handler ) - we use to handle http request by connecting path to function , http-Handler is a interface with function " serverHttp () "
   - that Function has to parameters - responseWritter ( is a interface used by http to construct response to the request to write respones back to request )  
        
         ex :   d, _ := ioutil.ReadAll(r.Body)

- http.Request ( the request we got from client and it has many components like body , header ,status etc ) , we have many framework to read and write data to read from body we can use , ioutil.ReadAll() 
          
      ex : fmt.Fprintf(w, "hellow %s\n", d)

- handle err and pass statusCode with response header
        
        if err != nil {
                       http.Error(w,"Opss",http.StatusBadRequest)

                        💡//above line can also  be replace with these 2 lines : 

                         w.WriteHeader(http.StatusBadRequest)
                         w.Write([]byte("Opss"))

                         return   
                }

## #Implimenting Handler using classes as separate package
### -handler package with hello.go file : 

      package handler

      import (
        "fmt"
        "io/ioutil"
        "log"
        "net/http"
          )

          💡// WAy of creating handler thats help in dependencies injection

          type Hello struct {
               l *log.Logger
          }

            func NewHello(l *log.Logger) *Hello { 💡// creat new hello object of class Hellow with log l
                      return &Hello{l} 💡// and return it , that object containg ServeHttp method that we build
                                       💡// for handling
          }


          func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {

                       log.Print("Hellow world")
                        d, err := ioutil.ReadAll(r.Body)

                   if err != nil {
                        http.Error(w, "Opss", http.StatusBadRequest)
                                💡//above line can also  be replace with these 2 lines :
                                // w.WriteHeader(http.StatusBadRequest)
                                // w.Write([]byte("Opss"))
                        return
          }

          fmt.Fprintf(w, "hellow %s\n", d)
        }
        

### -main.go file : 
_here what we do , instead of using default http.Handlrfunc to register our function ,
we create handler in handler package and than define new ServerMUX ( multiplexer ) and
than register our handler on it ._

    
    package main

    import (
        "example/learn0/handler"
        "log"
        "net/http"
        "os"
    )

    func main() {

        l := log.New(os.Stdout, "product-api", log.LstdFlags)
        hl := handler.NewHello(l) 💡// that will creat Hello object with l loger
       💡 //hl is our handler object with servehttp func act as handlerfunc

       💡 //here we define new serverMUX and than register our above handler to it.
        sm := http.NewServeMux()
        sm.Handle("/", hl)

        http.ListenAndServe(":9090", sm)
    }


### # we can creating own server :
- its has serval properties : check on Go doc. [ https://pkg.go.dev/net/http#Server ]

       
       💡//creating our own server , for tuning our application(API) for better performance
       
       s := http.Server{
		 Addr:         ":9090",
	 	Handler:      sm, // our servermux
	 	IdleTimeout:  120 * time.Second,
	 	ReadTimeout:  1 * time.Second,
	 	WriteTimeout: 1 * time.Second,
	  }
       s.ListenAndServe()
       
        💡//default server by http
        //http.ListenAndServe(":9090", sm)
        
        


### #granular shutdown ( first complete all task and than get cuttoff with client )
#### ▶️ use of goroutines ( go - keyword ) : 
_go starts a goroutine, which is managed by golang run-time. read here more :-> https://www.golang-book.com/books/intro/10_
     
           
      
      💡//handle my listenAndServe so that it not gonna block
	    go func() { // calls goroutines
	   	   err := s.ListenAndServe()
   	   	if err != nill {
	       		l.Fatal(err)
	            	}
        	}() //its a function call

-- for gracefull shutdown : 


	sigChan := make(chan os.Signal)

	💡// signal.notify is going to broadcast a message on this channel ( sigchan ) whenever  an operating system
	// kill command or os intrupts is receiveded
	// obviouslly we block here to consume the msg from the  channel
 
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	fmt.Println("Recieved terminate, graceful shutdown", sig)

	💡// shuting down server , by creating  a context with timedelay of 30s wait for all previous task to comlete and than go to shutdown
 
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	
	


## # Building RESTfull services using Go :

REST - stands for Represntational state transfer , so its an Architechtural pattern , one of the most commonly used .

### Lest convert our perious application to REST full : v1.0 -> v2.0 😄

## 🌟 You can see updates on branch v2.0 : go and check 🌟 
	
	




   

   

      
        
        






 


