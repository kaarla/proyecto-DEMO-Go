package controller

import(
    "github.com/kaarla/proyecto-DEMO/app/model"
	"encoding/json"
	"fmt"
	"net/http"
    //"html/template"
    //"encoding/csv"
    //"io"
    //"log"
    //"os"
    //"bufio"
)

var ListProducts model.Products = model.GetProducts()

type ProductsPage struct{
    Title string
    Productos model.Products
}

func GetProductsHandler(w http.ResponseWriter, r *http.Request){
    productJson, err := json.Marshal(ListProducts)
    if err != nil{
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    //p := ProductsPage{Title: "Productos", Productos: ListProducts}
    //t, _ := template.ParseFiles('index.html')
    //t.execute(w, p)
    w.Write(productJson)
}

func GetProductHandler(w http.ResponseWriter, r *http.Request){
    skuprod := r.Header.Get("sku")
    p := model.GetProduct(skuprod)
    productJson, err := json.Marshal(p)
    if err != nil{
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.Write(productJson)
}


