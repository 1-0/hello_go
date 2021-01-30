package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"encoding/json"
)

type post struct {
	UserId int   `json:"userId"`
	Id int      `json:"id"`
	Title string `json:"title"`
	Body string  `json:"body"`
}

type comment struct {
	PostId int   `json:"postId"`
	Id int      `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string  `json:"body"`
}


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
func getUserPosts(reqURL string, id int) []post {
	var url = reqURL + strconv.Itoa(id)
	resp, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal("Error reading request. ", err1)
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatal("Error reading response. ", err2)
	}
	var posts []post
	json.Unmarshal(body, &posts)
	return posts
}

// simpley get posts comments and print post
func getPostComments(reqURL string, id int) {
//func getPostComments(reqURL string, id int) []comment {
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
	var comments []comment
	json.Unmarshal(body, &comments)
	fmt.Println("comments ", comments)
	// return comments
}

func main() {
	const baseURL = "https://jsonplaceholder.typicode.com/"
	userID := 7
	url := baseURL + "posts?userId="
	url2 := baseURL + "comments?postId="
	listIDs := getUserPosts(url, userID)
	fmt.Println("len(listIDs) ", len(listIDs))
	for i := range listIDs {
		fmt.Println("listIDs[i].id ", listIDs[i].Id)
		go getPostComments(url2, listIDs[i].Id)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
	amt2 := time.Duration(rand.Intn(5))
	time.Sleep(time.Second * amt2)
}
