## # Building RESTfull services using Go :

REST - stands for Represntational state transfer , so its an Architechtural pattern , one of the most commonly used .

### Lest convert our perious application to REST full : v1.0 -> v2.0 

### ⭐ Project INTRO 🖊️ :
_"kjfkdsfksdjjfsdf"_


#### Lets get Start : 

1. Creat data package with list of products (for  now we are not using data base ) :

Data : 

 -> products.go : 
  
  		
		package data

		import "time"

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

	
	




   

   

      
        
        






 


