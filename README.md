		
## Devloping Basic Server :

#### Lets get Start : 

**1. Creat data package with list of products (for  now we are not using data base ) :**

Data : 

 -> products.go : 
  
		type Product struct {
			ID          int
			Name        string
			Description string
			Price       float32
			SKU         string
			CreatedOn   string
			UpdatedOn   string
			DeletedOn   string
		}

		var productList = []*Product{
			&Product{
				ID:          1,
				Name:        "siddharth",
				Description: "strong black milky coffee",
				Price:       3.67,
				SKU:         "abd234",
				CreatedOn:   time.Now().UTC().String(),
				UpdatedOn:   time.Now().UTC().String(),
			},
			&Product{
				ID:          2,
				Name:        "divu",
				Description: "light and  black milky coffee",
				Price:       1.99,
				SKU:         "xyz544",
				CreatedOn:   time.Now().UTC().String(),
				UpdatedOn:   time.Now().UTC().String(),
			},
		}
		

**2. Creat new handler for GET request to send whole product list in response as JSON format :**

here we are firse encode the product list into JSON and tha pass to response of handler -> 

		
	🌟 note : 🌟
	A utility package is supposed to provide some variables to a package who imports it.
	Like export syntax in JavaScript, Go exports a variable if a variable name starts with Uppercase. 
	All other variables not starting with an uppercase letter is private to the package.
		
	Read more about packaging : https://medium.com/rungo/everything-you-need-to-know-about-packages-in-go-b8bac62b74cc


_Creating handler_ : 


		type Product struct {
			l *log.Logger
		}

		func NewProduct(l *log.Logger) *Product {
			return &Product{l}
		}

		func (p *Product) ServeHTTP(w http.ResponseWriter, r *http.Request) {
			//Encoding to json

			lp := data.GetProduct()
			d, err := json.Marshal(lp)
			if err != nil {
				http.Error(w, "Unable to encode to jason", http.StatusInternalServerError)
			}

			// send back with respone
			w.Write(d)
		}


_Registering our handler to our ServerMux in main.go file_ : 
	
	ph := handler.NewProduct(l)    // that will creat Product handler object with l loger
	                               //ph is our handler object with servehttp func act as handlerfunc

	//here we define new serverMUX and than register our above handler to it.
	
	sm := http.NewServeMux()
	sm.Handle("/", ph) 

_run command to get productlist from our running server_ : 
	
	curl localhost:9090 | jq

**output**

<img width="598" alt="Screenshot 2023-03-24 at 12 56 54 AM" src="https://user-images.githubusercontent.com/87073574/227327343-d7ef70d3-b841-47fc-af70-c2395f11b237.png">


_use struct tags to add annotation to your productlist for better output_ :

<img width="494" alt="Screenshot 2023-03-24 at 1 04 38 AM" src="https://user-images.githubusercontent.com/87073574/227330390-b17b4813-151e-4e44-82d8-5fecfe83af12.png">

**change in output :**
	
<img width="432" alt="Screenshot 2023-03-24 at 1 04 00 AM" src="https://user-images.githubusercontent.com/87073574/227330123-c801a92a-e044-447c-a121-f40625482a02.png">

	" as we know we use json.Marshal to encode our jason data but while using this we alocated memory and which make it slower , 
  	so what we do - go has json.encoder() that is very fast and encode the data and write to response directly "

_lets add encoding and writing logic to product in data package_ :
	
	func (p *Products) ToJson(w io.Writer)error {
	//defining encoder
	e := json.NewEncoder(w)

	//pass our product list to encoder to write
	return e.Encode(p)
	}
	
_change GetProducts() signature to return our new custom type Products_ : 

	func GetProduct() Products {
	return productList
	}

_call ToJson() in product handler :
	
	//Encoding to json
	lp := data.GetProduct()
	// d, err := json.Marshal(lp)
	err := lp.ToJson(w)
	

_Now we want to handle defferent http request like GET , POST etc 
so we define logic in serveHTTP() in product handler and add some internal functions_ : 
	
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

_-If your now trying to make request other than GET you get   ->_ 😮

<img width="1215" alt="Screenshot 2023-03-24 at 1 54 41 AM" src="https://user-images.githubusercontent.com/87073574/227344783-cf7c2590-97d9-4efb-8961-e2d5f1297472.png">



	
	




   

   

      
        
        






 


