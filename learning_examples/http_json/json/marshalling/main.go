package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	//simple marshal example
	{
		data := map[string]any{
			"a": 54,
			"b": "text",
			"c": struct{}{},
			"d": 32.2,
		}

		jsonBytes, err := json.Marshal(data)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonBytes))
	}

	//indent marshal example
	{
		data := map[string]int{
			"a": 123,
			"b": 7,
		}

		jsonBytesWithIndent, err := json.MarshalIndent(data, "prefix ", "indent ")

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(jsonBytesWithIndent))
	}
}
