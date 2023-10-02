package main


import (
	"net/http"
	"fmt"
)

func main(){
	http.HandleFunc("/hello", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "<h1>Hello Bangladesh</h1>")
}