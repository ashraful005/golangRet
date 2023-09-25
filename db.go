package main

import (
	"database/sql"
	"time"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// ...
var db *sql.DB 
var err error
func init(){
db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/youthict?charset=utf8")
if err != nil {
	panic(err)
}
// See "Important settings" section.
db.SetConnMaxLifetime(time.Minute * 3)
db.SetMaxOpenConns(10)
db.SetMaxIdleConns(10)
}

func main(){

fmt.Println(db.Ping())

// rows, err := db.Query("SELECT * FROM company;")
// if err != nil{
// 	fmt.Println(err)
// }
var id int 
var name, email string
prow := db.QueryRow("SELECT id,name,email FROM company;")
prow.Scan(&id,&name,&email)

fmt.Println("id", id)
fmt.Println("name", name)
fmt.Println("email", email)

}