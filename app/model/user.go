package model

import(
     "fmt"
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

type User struct{
    Id string `json:"UserID"`
    TransactionId string `json:"id"`
}

func GetCompras(userId string) Products{
    ps := Products{}
    c := 0
    fmt.Println("extracting products information from database ./db/demo.db")
    database, e := sql.Open("sqlite3", "./db/demo.db")
    if e == nil {
        fmt.Println("no error opening")
    }
    q := "select p.name, p.sku from productTransaction as pt inner join (select * from transactionIDuser where userid == " + userId + ") as t on t.id = pt.id inner join (select * from producto) as p on p.sku = pt.sku"
    rows, error := database.Query(q)
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
    }
    return ps;
}

