## # Building RESTfull services using Go :

REST - stands for Represntational state transfer , so its an Architechtural pattern , one of the most commonly used .

### Lest convert our perious application to REST full : v1.0 -> v2.0 

### ⭐ Project INTRO 🖊️ :
_"kjfkdsfksdjjfsdf"_


### Start refactoring by using standard libraries of Go-lang : 

- now for REST full approach what you'r going to be doing is using HTTP verbs { like PUT , GET , POST , etc. }

▶️ Lets impliment a **POST** to be able to add a new product $  
  
    " now here we have to do opposite of GET request , we have to decode request Body ( actually a io.writter )
      that is in Json format into our Database store format
      for that we are using Decoder() which accept io.writer( request body ) as argument " 

_add new addProducts() to handler  having decoding logic. just below our Product struct_ :
    
    
    //here as a function parameter we get body of post request with our json content 
    //and we decode it to structure of our Product struct

    func (p *Products) FromJson(r io.Reader) error {
      // defining decoder
      e := json.newDecoder(r)

      // return decoded data to structure of Products type
      return e.Decode(p)
    }

_call this function in handler { addProduct() }  of Post request_ :
    
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
    
now we have to add this decoded struct data into our static list of product ( Datastore ) : 

for that 1. creat addProducts() for list and 2. creat a function that generate interger for next ID 
      
        // add new product to our static list 
        func AddProduct(p *Product){
          p.ID = getNextID()
          productList = append(productList ,p)  // appended our new Product structured data to our existing List 
        }

        //fumnction to generate integer for our ID
        func getNextID() int {
          lp := productList[len(productList)-1]  
          return lp.ID + 1  // increment by 1 of Last product Id in the list
        }
        
now send post request to our server with some json data 

- On server side 

<img width="1168" alt="Screenshot 2023-03-24 at 2 04 35 PM" src="https://user-images.githubusercontent.com/87073574/227466913-844a5df7-554e-41c6-9031-f5e2f39e436d.png">

- On client side { sending POST request with body ( -d ) and in verbos mod ( -v ) :

<img width="959" alt="Screenshot 2023-03-24 at 2 06 47 PM" src="https://user-images.githubusercontent.com/87073574/227467426-8d6d4a2b-c49f-4bb6-b9d0-c45a447a8eee.png">

- Check Whether new product is added or not :  curl localhost:9090 | jq 

<img width="616" alt="Screenshot 2023-03-24 at 2 11 50 PM" src="https://user-images.githubusercontent.com/87073574/227468524-1628cc9a-32fe-4f89-9585-a6deba57a766.png">


### For implementing PUT ( update ) method 
-for that we have to  add logic to ServeHttp() where we haveto extract ID from URL from client Request 
  and for extracting we use FindAllStringSubmatch() form  Regexp package : for more about it read :   
  
  https://cs.opensource.google/go/go/+/go1.20.2:src/regexp/regexp.go;l=1197 
  
  



        
  
  
