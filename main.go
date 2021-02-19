package main

import (
	"fmt"
	"net/http"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Teste 1</h1>")
}

func main(){
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe(":80", nil))
}