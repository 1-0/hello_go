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
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type post struct {
	UserID int   `json:"userId"`
	ID int      `json:"id"`
	Title string `json:"title"`
	Body string  `json:"body"`
}

type comment struct {
	PostID int   `json:"postId"`
	ID int      `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Body string  `json:"body"`
}

// simpley write post to DB
func insertPost(newPost post) {
	db, err := sql.Open("mysql",
		"root:2w2w2w2w2w@tcp(127.0.0.1:3306)/hello_go")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("db ", db)
	stmt, err := db.Prepare("INSERT INTO posts(UserID, ID, Title, Body) VALUES(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(newPost.UserID, newPost.ID, newPost.Title, newPost.Body)
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
	log.Printf("insertPost %d ID = %d, affected = %d\n", newPost.ID, lastId, rowCnt)
	// var (
	// 	id int
	// 	name string
	// )
	// rows, err := db.Query("select id from post where id = ?", 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	err := rows.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// }
	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// 	defer db.Close()
	// }
}

// simpley write comment to DB
func insertComment(newComments comment) {
	db, err := sql.Open("mysql",
		"root:2w2w2w2w2w@tcp(127.0.0.1:3306)/hello_go")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("db ", db)
	stmt, err := db.Prepare("INSERT INTO comments(PostID, ID, Name, Email, Body) VALUES(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(newComments.PostID, newComments.ID, newComments.Name, newComments.Email, newComments.Body)
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
	log.Printf("newComments %d ID = %d, affected = %d\n", newComments.ID, lastId, rowCnt)
}

// simpley get posts comments and print post
func getPostComments(reqURL string, id int) {
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
	var comments []comment
	json.Unmarshal(body, &comments)
	for c := range comments{
		// fmt.Println("comments[c].ID ", comments[c].ID)
		go insertComment(comments[c])
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

// simpley get user posts id and print post
func getUserPosts(reqURL string, id int) {
	const baseURL = "https://jsonplaceholder.typicode.com/"
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
	url2 := baseURL + "comments?postId="
	for i := range posts {
		go insertPost(posts[i])
		go getPostComments(url2, posts[i].ID)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}	
}

func main() {
	const baseURL = "https://jsonplaceholder.typicode.com/"
	userID := 7
	url := baseURL + "posts?userId="
	go getUserPosts(url, userID)
	amt2 := time.Duration(5)
	time.Sleep(time.Second * amt2)
	fmt.Println("-----------------------exit--------------------")
}
