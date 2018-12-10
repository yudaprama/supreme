package supreme

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

var query = `
query {
  userById(id:100) {
    id
  }
}
`

func message(variables string) string {
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
	fmt.Println(message("{}"))
	//graph("query {\n  userById(id:100) {\n    id\n  }\n}")
}