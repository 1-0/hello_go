package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	var url = reqURL + strconv.Itoa(id)
	fmt.Println("url:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error reading request. ", err)
		log.Fatal("Error reading request. ", err)
	}
	defer resp.Body.Close()
	// defer resp.Close()
	// fmt.Println("url:", url)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response. ", err)
		log.Fatal("Error reading response. ", err)
	}
	fmt.Println(string(body))
	amt := time.Duration(rand.Intn(250))
	time.Sleep(time.Millisecond * amt)
	// time.Sleep(time.Second * 1)
}

func main() {
	const baseURL = "https://jsonplaceholder.typicode.com/"
	url := baseURL + "posts/"
	for i := 1; i <= 100; i++ {
		printPost(url, i)
	}
}
