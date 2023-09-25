
package main

import (
	"net/url"
	"fmt"
  "log"
  "time"
  "database/sql"
  "github.com/mateors/msql"
_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init(){

	// Connect to database
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/youthict?charset=utf8")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	//defer db.Close()
	log.Println("db connection successful")

}

func insertUpdate(){
	//insert into company
	m := make(url.Values)
	dt := time.Now().Format("2004-01-02 17:04:07")
	m.Set("name","emon")
	m.Set("email","emdash@gmail.com")
	m.Set("create_date",dt)

	m.Set("todo","insert")
	m.Set("id","1")
	m.Set("pkfield","id")
	m.Set("table","company")

	msql.InsertUpdate(m, db)
}


func retrieve(){
	rows, err := msql.GetAllRowsByQuery("select name,email,status from company;",db)
	if err!=nil{
		log.Println(err)
	}
	for _, row := range rows{
		fmt.Println(row)
	}
}

func delete(){
	_, err := db.Query("DELETE FROM company WHERE id=?",1)
	if err!=nil{
		log.Println(err)
	}
}
func main(){

	

	insertUpdate()
	//retrieve
	retrieve()

	//delete()
	
  
}