package model

import(
    "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
    
)

type Product struct{
    SKU string  `json:"sku"`
    Name string `json:"name"`
}

type Products []Product

func GetProducts() Products{
    ps := Products{}
    c := 0
    fmt.Println("extracting products information from database ./db/demo.db")
    database, e := sql.Open("sqlite3", "./db/demo.db")
    if e == nil {
        fmt.Println("no error opening")
    }
    rows, error := database.Query("SELECT sku, name FROM producto")
    if error != nil {   
        fmt.Println("errROW", error)
    }else{
        var sku string
        var name string
        for rows.Next(){
            rows.Scan(&sku, &name)
            ps = append(ps, Product{
                SKU: sku,
                Name: name,
            })
            c++
            if(c > 10){
                break;
            }
        }
        p := GetProduct("38810171")
        fmt.Println(p.SKU)
        fmt.Println("productos cargados")
    }
    return ps;
}

func GetProduct(sku string) Product{
    p := Product{}
    database, e := sql.Open("sqlite3", "./db/demo.db")
    if e == nil {
        fmt.Println("no error opening")
    }
    q := "SELECT sku, name FROM producto where sku == " + sku
    rows, error := database.Query(q)
    if error != nil {   
        fmt.Println("errROW", error)
    }else{
        var sku string
        var name string
        rows.Next()
        rows.Scan(&sku, &name)
        p = Product{
                SKU: sku,
                Name: name,
            }
        fmt.Println("p.sku", p.SKU, "p.name", p.Name)
    }
    return p
}
