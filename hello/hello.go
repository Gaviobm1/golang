package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"github.com/google/go-cmp/cmp"
)

func main() {
	names := []string{"Jackson", "Mateo", "Allen"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
