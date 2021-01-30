package main

import (
	"fmt"
	"io/ioutil"
	"log"
	// "math/rand"
	"net/http"
	"strconv"
	// "time"
	// "encoding/json"
)

// type post struct {
// 	UserId int   `json:"userId"`
// 	Id int      `json:"id"`
// 	Title string `json:"title"`
// 	Body string  `json:"body"`
// }

// type comment struct {
// 	PostId int   `json:"postId"`
// 	Id int      `json:"id"`
// 	Name string `json:"name"`
// 	Email string `json:"email"`
// 	Body string  `json:"body"`
// }


// simpley write post
// func writePost(post []byte, path string, id int) {
// 	//message := []byte("Hello, Gophers!")
// 	var newFile = path + strconv.Itoa(id)
// 	err := ioutil.WriteFile(newFile, post, 0644)
// 	if err != nil {
// 		fmt.Println(err)
// 		log.Fatal(err)
// 	}
// 	fmt.Println("newFile - ", newFile)
// }

// simpley get and print post
// func getPosts(reqURL string, id int) {
// 	var writePath = "./storage/posts/"

// 	var url = reqURL + strconv.Itoa(id)
// 	// fmt.Println("url:", url)
// 	resp, err1 := http.Get(url)
// 	if err1 != nil {
// 		log.Fatal("Error reading request. ", err1)
// 	}
// 	defer resp.Body.Close()
// 	body, err2 := ioutil.ReadAll(resp.Body)
// 	if err2 != nil {
// 		log.Fatal("Error reading response. ", err2)
// 	}
// 	writePost(body, writePath, id)
// 	// fmt.Println(string(body))
// }

// simpley get user posts id and print post
func getUserPostsID(reqURL string, id int) string {
	var url = reqURL + strconv.Itoa(id)
	// fmt.Println("url:", url)
	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal("Error reading request. ", err1)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal("Error reading response. ", err2)
	}
	// writePost(body, writePath, id)
	// fmt.Println(string(body))
	res := string(body)
	return res
}

func main() {
	const baseURL = "https://jsonplaceholder.typicode.com/"
	userID := 7
	url := baseURL + "posts?userId="
	// url2 := baseURL + "/comments?postId="
	listIDs := getUserPostsID(url, userID)
	fmt.Println(string(listIDs))
	// for i := 1; i <= 100; i++ {
	// 	go getPost(url, i)
	// 	amt := time.Duration(rand.Intn(250))
	// 	time.Sleep(time.Millisecond * amt)
	// }
	// amt2 := time.Duration(rand.Intn(20))
	// time.Sleep(time.Second * amt2)
}
