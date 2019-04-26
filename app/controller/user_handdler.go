package controller

import(
    "github.com/kaarla/proyecto-DEMO/app/model"

)

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetComprasHandler(w http.ResponseWriter, r *http.Request){
    userId := r.Header.Get("userId")
    products := model.GetCompras(userId)
    productJson, err := json.Marshal(products)
    if err != nil{
        fmt.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.Write(productJson)
}
