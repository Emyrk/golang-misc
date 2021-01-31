package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Request struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Status      string   `json:"status"`
	Tags        []string `json:"tags"`
}

func main() {
	input := "golang, elixir, python, java"
	tags := strings.Split(input, ",")

	postBody, _ := json.Marshal(Request{
		Name:        "name",
		Description: "desc",
		Status:      "status",
		Tags:        tags,
	})
	fmt.Println(string(postBody))
}
