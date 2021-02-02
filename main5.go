package main

import (
	"fmt"
	"sync"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var wg sync.WaitGroup

// simpley write post
func writePost(post []byte, path string, id int) {
	var newFile = path + strconv.Itoa(id)
	err := ioutil.WriteFile(newFile, post, 0644)
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	fmt.Println("newFile - ", newFile)
}

// simpley get and print post
func getPost(reqURL string, id int) {
	var writePath = "./storage/posts/"

	var url = reqURL + strconv.Itoa(id)
	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal("Error reading request. ", err1)
	}
	defer resp.Body.Close()
	defer wg.Done()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal("Error reading response. ", err2)
	}
	writePost(body, writePath, id)
}

func main() {
	const baseURL = "https://jsonplaceholder.typicode.com/"
	url := baseURL + "posts/"
	//_ = os.MkdirAll("./storage/posts/", 0770)
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go getPost(url, i)
	}
	wg.Wait()
}
