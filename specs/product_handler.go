package main

import (
	"encoding/json"
	"fmt"
	"net/http"
    "encoding/csv"
    "io"
    "log"
    "os"
    "bufio"
)

type Product struct{
    Sku string `json:"SKU"`
    Name string `json:"Name"`
}

var products []Product


func getProductHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "productss" variable to json
	productsListBytes, err := json.Marshal(products)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
// 	// If all goes well, write the JSON list of products to the response
	w.Write(productsListBytes)
}




func createProductHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance of product
	product := Product{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the product from the form info
	product.Sku = r.Form.Get("sku")
	product.Name = r.Form.Get("name")

	// Append our existing list of products with a new entry
	products = append(products, product)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}

func chargeProductsHandler(w http.ResponseWriter, r *http.Request){
    csvFile, _ := os.Open("specs/proyecto_test/product_detail.csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))
   // var product []Product
    for{
        line, error := reader.Read()
        if error == io.EOF{
            break
        }else if error != nil{
            log.Fatal(error)
        }
        products = append(products, Product{
            Sku: line[0],
            Name: line[1],
        })
        
    }
    http.Redirect(w, r, "/assets/", http.StatusFound)
    

//    productJson, _ := json.Marshal(products)
  //  w.Write(productJson)
}


