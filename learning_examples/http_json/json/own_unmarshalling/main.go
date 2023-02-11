package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type structWithPrivateField struct {
	privateField string
}

// UnmarshalJSON redeclare unmarshal method
func (s *structWithPrivateField) UnmarshalJSON(json []byte) error {
	const prefix, suffix = `{"privateField": "`, `"}`

	json = bytes.TrimPrefix(json, []byte(prefix))
	json = bytes.TrimSuffix(json, []byte(suffix))
	json = bytes.TrimSpace(json)

	s.privateField = string(json)

	return nil
}

func main() {
	//for checking contract of our structure. Analog of "instance of". Optional line of code
	var _ json.Unmarshaler = (*structWithPrivateField)(nil)

	data := new(structWithPrivateField)
	err := json.Unmarshal([]byte(`{"privateField": " TEST  "}`), data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data.privateField)
}
