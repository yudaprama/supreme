package main

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

func bisa() {
	var jsonValue = []byte(`{"query":"query {\n  userById(id:100) {\n    id\n  }\n}","variables":null}`)
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

func message(query string, variables string) string {
	values := map[string]string{"query": query, "variables": variables}
	jsonValue, _ := json.Marshal(values)
	fmt.Println("jsonValue", string(jsonValue))
	resp, err := http.Post("https://graphql.tebengan.id/graphql", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}

func main() {
	//bisa()
	fmt.Println(message("query {\n  userById(id:100) {\n    id\n  }\n}","{}"))
	//graph("query {\n  userById(id:100) {\n    id\n  }\n}")
}