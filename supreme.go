package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
)

func main() {
	var jsonValue = []byte(`{"query":""query {↵  userById(id:100) {↵    id↵  }↵}","variables":{}}`)
	resp, err := http.Post("https://graphql.tebengan.id/graphql", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	
	fmt.Println("response Status:", resp.Status)
  fmt.Println("response Headers:", resp.Header)
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Println("response Body:", string(body))
}