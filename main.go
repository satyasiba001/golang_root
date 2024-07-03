package main

import "fmt"

func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", HomeHandler)
	r.HandleFunc("/articles", Articlehandler)
	http.Handle("/",r)
}