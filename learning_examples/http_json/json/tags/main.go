package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type exampleStruct struct {
	FirstField    int
	SecondField   int   `json:"renamedSecondField"`
	ThirdField    int   `json:"renamedThirdField,omitempty"` //omitempty - for ignoring if zero value
	FourthField   int   `json:",omitempty"`
	FifthField    int   `json:"-"`       //for ignoring field
	SixthField    int   `json:"-,"`      //for rename field
	Int64ToString int64 `json:",string"` //for example - JS can't handle int64 number
}

func main() {
	//marshalling structure with json tag
	in := exampleStruct{
		FirstField:    1,
		SecondField:   2,
		ThirdField:    3,
		FourthField:   4,
		FifthField:    5,
		SixthField:    6,
		Int64ToString: math.MaxInt64,
	}

	jsonBytes, err := json.MarshalIndent(in, "", "\t")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonBytes))

	//unmarshalling structure with json tag
	out := exampleStruct{}

	err = json.Unmarshal(jsonBytes, &out)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", out)
}
