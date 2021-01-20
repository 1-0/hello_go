package main
 
import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// simpley http client
func client_http(req_url, req_type string) string {
	req, err := http.NewRequest(req_type, req_url, nil)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}
 
	req.Header.Set("Cache-Control", "no-cache")
 
	client := &http.Client{Timeout: time.Second * 10}
 
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error reading response. ", err)
	}
	defer resp.Body.Close()
 
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
	}
	return string(body)
}
 
func main() {
	const base_url = "https://jsonplaceholder.typicode.com/"
	var res string
	url := base_url + "posts"
	r_type := "GET"
	res = client_http(url, r_type)

	fmt.Printf("%s\n", res)
}

