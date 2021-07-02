package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("https://blog.syhlion.tw/sitemap.xml")
	if err != nil {
		log.Fatal(err)
	}
	// res.Body 讀取完，要記得 Close，不然會有 memory leak 等相關問題
	defer res.Body.Close()

	sitemap, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", sitemap)
}
