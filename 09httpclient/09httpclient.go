package main

import (
	"fmt"
	"log"
	"time"

	"github.com/syhlion/greq"
	requestwork "github.com/syhlion/requestwork.v2"
)

/*
package github.com/syhlion/greq
這是Scott因應業務需求，需要有大量可以存取各式 rest api 並且又要控制好併發數量，而開發出來的套件，使用方式如下
*/

func main() {

	//need import https://github.com/syhlion/requestwork.v2
	// 這是最高同時併發的控制，是 greq 的核心套件
	worker := requestwork.New(50)

	client := greq.New(worker, 10*time.Second, false)

	//GET
	data, httpstatus, err := client.Get("http://localhost:8080/api/query/steve", nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("data:%s ,httpstatus:%d", data, httpstatus)
}
