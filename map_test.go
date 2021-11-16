package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapDecode(t *testing.T) {
	jsonString := `{"id":"B001", "name":"Asus Rog GG", "price":72000000}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id" : "B001",
		"name" : "Asus Rog 72 Juta",
		"price" : 2000000,
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}