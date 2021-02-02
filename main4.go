package main

import (
	"fmt"
	"sync"
	"io/ioutil"
	"net/http"
	"strconv"
)

var wg sync.WaitGroup

// simpley get and print post
func printPost(reqURL string, id int) {

	var NewURL = reqURL + strconv.Itoa(id)
	fmt.Println("url:", NewURL)
	resp, err1 := http.Get(NewURL)
	if err1 != nil {
		fmt.Println("Error reading request. ", err1)
	}
	defer resp.Body.Close()
	defer wg.Done()
	var body, err2 = ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("Error reading response. ", err2)
	}
    fmt.Println(string(body))
}

func main() {
	const baseURL = "https://jsonplaceholder.typicode.com/"
	url := baseURL + "posts/"
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go printPost(url, i)
	}
    wg.Wait()
}
