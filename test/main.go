package main

import (
	"fmt"
	"github.com/ermos/datagouv"
	"log"
)

func main() {
	c, err := datagouv.GetCommune(datagouv.CommuneParameters{
		Name: "la flèche",
		Fields: []string{ "name", "code" , "codesPostaux" },
		Code: "72200",
		Limit: 10,
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
