package main

import (
	"fmt"
	"log"
	"net/http"
)

// http 資料來源：https://ithelp.ithome.com.tw/articles/10207214?sc=iThomeR

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
