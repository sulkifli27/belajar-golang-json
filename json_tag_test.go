package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id string `json:"id"`
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func TestJsonTag(t *testing.T) {
	product := Product{
		Id: "B001",
		Name: "Asus Rog 72 juta",
		ImageUrl: "example.com/asus.png",
	}

	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}

func TestJsonTagDecode(t *testing.T) {
	jsonString := `{"Id":"B001","Name":"Asus Rog 72 juta","Image_Url":"example.com/asus.png"}`
	jsonBytes := []byte(jsonString)
	
	product := &Product{}

	json.Unmarshal(jsonBytes, product)
	fmt.Println(product)
}