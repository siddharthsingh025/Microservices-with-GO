# # Rectoring and building API using GorilaMUX framework :
## Lets Get Start.....

" what we do now , we first delete ServeHTTP() method from product handler and make all functions for get , add , update request public 
  and in main.go we creat Router using Gorilamux and creat many subroutes specific to our request methods and call all function to them "
  
  1. go.main file :
    
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
    
   2.product.go (handler) file : 
      

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
                  pdt := &data.Product{}
                  err := pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
                  if err != nil {
                    http.Error(w, "Unable to unmarshall json", http.StatusBadRequest)
                  }

                  p.l.Printf("product : %#v", pdt) // it will print decoded data into nice format in logWindow

                  data.AddProduct(pdt) // call func to add decoded data in list
                }

                // for PUT req
                func (p Product) UpdateProducts(w http.ResponseWriter, r *http.Request) {
                  // mux extract id from URL using Vars()
                  vars := mux.Vars(r)
                  id, err := strconv.Atoi(vars["id"])

                  if err != nil {
                    http.Error(w, "enable to convert id", http.StatusBadRequest)
                  }

                  p.l.Println("Handle PUT Products", id)
                  pdt := &data.Product{}

                  err = pdt.FromJson(r.Body) // we call FromJson func of Product and pass body of post request we got
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


### Refactoring done using gorilla mux  
    




