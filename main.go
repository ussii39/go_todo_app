package main

import (
	"fmt"
	"net/http"
	"os"
)

func main(){
err := http.ListenAndServe(
	":8000",
	http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w,r.URL.Path[1:])
	}),
)
if err != nil {
	fmt.Printf("failed")
	os.Exit(1)
}	
}