package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Professor struct {
	Name string `json:"name"`
	ScienceID int `json:"science_id"`
	IsWorking bool `json:"is_working"`
	University University `json:"university"`
}

type University struct {
	Name string `json:"name"`
	City string `json:"city"`
}


func main() {
	prof1 := Professor{
		Name: "Petr",
		ScienceID: 82734628,
		IsWorking: true,
		University: University{
			Name: "BMSTU",
			City: "Moscow",
		},
	}

	// 1. professor -> bytes sequence
	byteArr, err := json.MarshalIndent(prof1, "", "	") // " " -> "\t"
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(byteArr))

	// 2. save to file
	err = os.WriteFile("output.json", byteArr, 0666) // -rw-rw-rw-
	if err != nil {
		log.Fatal(err)
	}

}