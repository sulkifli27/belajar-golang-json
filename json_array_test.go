package belajargolangjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJsonArray(t *testing.T) {
	customer := Customer{
		FirstName: "Sul",
		LastName: "Kifli",
		Hobbies: []string{"game","coding","sleep"},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Sul","LastName":"Kifli","Address":"","Hobbies":["game","coding","sleep"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Hobbies)
}

func TestJsonArrayComplex(t *testing.T) {
	customer := Customer{
		FirstName:  "Sulkifli",
		Addresses: []Address{
			{
				Street: "Jalan Dusun Salohe",
				Country: "Kabupaten Sinjai",
				PostalCode: "92654",
			},
			{
				Street: "Jalan Dusun Bulo",
				Country: "Kabupaten Sinjai",
				PostalCode: "92654",
			},
		},
	}

	bytes, _ := json.Marshal(customer)
	fmt.Println(string(bytes))
}

func TestJsonArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Sulkifli","LastName":"","Address":"","Hobbies":null,"Addresses":[{"Street":"Jalan Dusun Salohe","Country":"Kabupaten Sinjai","PostalCode":"92654"},{"Street":"Jalan Dusun Bulo","Country":"Kabupaten Sinjai","PostalCode":"92654"}]}`

	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)

	if err != nil {
		panic(err)
	}

	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.Addresses)
}

func TestOnlyJsonArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Jalan Dusun Salohe","Country":"Kabupaten Sinjai","PostalCode":"92654"},{"Street":"Jalan Dusun Bulo","Country":"Kabupaten Sinjai","PostalCode":"92654"}]`

	jsonBytes := []byte(jsonString)

	addresess := &[]Address{}
	err := json.Unmarshal(jsonBytes, addresess)

	if err != nil {
		panic(err)
	}

	fmt.Println(addresess)

}

func TestOnlyJsonArrayComplex(t *testing.T) {
	addresses := []Address{
		{
				Street: "Jalan Dusun Salohe",
				Country: "Kabupaten Sinjai",
				PostalCode: "92654",
		},
			{
				Street: "Jalan Dusun Bulo",
				Country: "Kabupaten Sinjai",
				PostalCode: "92654",
			},
	}

	bytes, _ := json.Marshal(addresses)
	fmt.Println(string(bytes))
}