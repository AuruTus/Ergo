package main

import (
	"encoding/json"
	"fmt"
)

type S struct {
	Name string            `json:"name"`
	Data map[string]string `json:"data"`
}

func main() {
	s := new(S)
	raw := `{"name":"text", "data":{"text":" plz "}}`
	fmt.Println(raw)
	if err := json.Unmarshal([]byte(raw), s); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", s)
	fmt.Println(len(s.Name))
}
