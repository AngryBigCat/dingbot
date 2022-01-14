package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestPost(t *testing.T) {
	client := new(HttpClient)

	data, err := client.Post("http://api.91tool.com/tb/accounts", map[string]string{
		"apikey":  "test",
		"account": "test",
		"cache":   "1",
	})
	if err != nil {
		log.Println(err.Error())
	}

	fmt.Println(data["msg"])
}
