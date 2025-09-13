package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler (w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/hello" {
		http.Error(w,"The path appears to be wrong", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w,"Does not support any other Methods other than GET", http.StatusNotFound)
		return
	} 
	fmt.Fprintf(w,"Hello")
}

func formHandler (w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parseform error %v", err)
		return
	} 
	fmt.Fprintf(w, "Parseform success \n")
	name := r.FormValue("name")
	age := r.FormValue("age")
	fmt.Fprintf(w, "name = %s \n", name)
	fmt.Fprintf(w, "age = %s \n", age)
}

func main(){
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileserver)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fmt.Printf("Starting a server on localhost 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	} 
}
