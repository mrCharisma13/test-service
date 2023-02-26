package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r*http.Request){
		log.Print("Hello world")
		d, err :=  ioutil.ReadAll(r.Body)
		fmt.Fprintf(rw, "Hello %s", d)

		if err != nil {
			http.Error(rw, "Ooops", http.StatusBadRequest)
			return
		}
	})

	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request){
		log.Print("Goodbye world")
	})

	http.ListenAndServe(":9090", nil)
}