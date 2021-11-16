package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street string
	Country string
	PostalCode string
}

type Customer struct {
	FirstName string
	LastName string
	Address string
	Hobbies []string
	Addresses []Address
}

func TestJsonObject(t *testing.T) {
	customer := Customer{
		FirstName: "Sul",
		LastName: "Kifli",
		Address: "Sinjai",
	}
	bytes, _ :=json.Marshal(customer)
	fmt.Println(string(bytes))
}