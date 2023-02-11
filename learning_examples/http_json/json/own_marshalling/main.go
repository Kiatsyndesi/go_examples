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

// MarshalJSON redeclare marshal method
func (s *structWithPrivateField) MarshalJSON() ([]byte, error) {
	const prefix, suffix = `{"privateField": "`, `"}`

	buf := new(bytes.Buffer)
	buf.Grow(len(prefix) + len(suffix) + len(s.privateField))

	buf.WriteString(prefix)
	buf.WriteString(s.privateField)
	buf.WriteString(suffix)

	return buf.Bytes(), nil
}

func main() {
	//for checking contract of our structure. Analog of "instance of". Optional line of code
	var _ json.Marshaler = (*structWithPrivateField)(nil)

	data := &structWithPrivateField{privateField: "hello gopher"}

	jsonBytes, err := json.Marshal(data)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))
}
