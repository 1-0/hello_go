package main

import (
	"fmt"
	"io/ioutil"
	// "log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	// "encoding/json"
)

/* type post struct {
	UserId int     `json:"userId"`
	Id  int           `json:"id"`
	Title string  `json:"title"`
	Body string `json:"body"`
} */

// simpley get and print post
func printPost(reqURL string, id int) {

	var NewURL = reqURL + strconv.Itoa(id)
	fmt.Println("url:", NewURL)
	resp, err1 := http.Get(NewURL)
	if err1 != nil {
		fmt.Println("Error reading request. ", err1)
	}
	defer resp.Body.Close()
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
		go printPost(url, i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
	amt2 := time.Duration(rand.Intn(20))
	time.Sleep(time.Second * amt2*5)
}
