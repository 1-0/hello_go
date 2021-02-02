package main

import (
	// "fmt"
	"io/ioutil"
	"sync"
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const baseURL = "https://jsonplaceholder.typicode.com/"
var wg sync.WaitGroup

// simpley save to DB record
var i interface {
    saveToDB()
}

type post struct {
	UserID int   `json:"userId"`
	ID int      `json:"id"`
	Title string `json:"title"`
	Body string  `json:"body"`
}

func (p post) saveToDB() {
	defer wg.Done()
	db, err := sql.Open("mysql",
		"root:2w2w2w2w2w@tcp(127.0.0.1:3306)/hello_go")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO posts(UserID, ID, Title, Body) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(p.UserID, p.ID, p.Title, p.Body)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("insertPost %d ID = %d, affected = %d\n", p.ID, lastId, rowCnt)
}

type comment struct {
	PostID int   `json:"postId"`
	ID int      `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string  `json:"body"`
}

func (c comment) saveToDB() {
	defer wg.Done()
	db, err := sql.Open("mysql",
		"root:2w2w2w2w2w@tcp(127.0.0.1:3306)/hello_go")
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare("INSERT INTO comments(PostID, ID, Name, Email, Body) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(c.PostID, c.ID, c.Name, c.Email, c.Body)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("newComments %d ID = %d, affected = %d\n", c.ID, lastId, rowCnt)
}

// simpley get posts comments and print post
func getPostComments(reqURL string, id int) {
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
	var comments []comment
	json.Unmarshal(body, &comments)
	for c := range comments{
		wg.Add(1)
		go comments[c].saveToDB()
	}
}

// simpley get user posts id and print post
func getUserPosts(reqURL string, id int) {
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
	var posts []post
	json.Unmarshal(body, &posts)
	url2 := baseURL + "comments?postId="
	for i := range posts {
		wg.Add(1)
		go posts[i].saveToDB()
		wg.Add(1)
		go getPostComments(url2, posts[i].ID)
	}	
}

func main() {
	userID := 7
	url := baseURL + "posts?userId="
	wg.Add(1)
	go getUserPosts(url, userID)
	wg.Wait()
}
