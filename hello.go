package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	for i := 1; i <= 100; i++ {
		go RequestAndPrint(i)
	}
	var input string
	fmt.Scanln(&input)

}

func RequestAndPrint(i int) {
	res, err := http.Get(fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", i))
	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", robots)
}
