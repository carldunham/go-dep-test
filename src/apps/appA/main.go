package main

import (
	"fmt"

	"libs/coolstuff"

	"github.com/google/uuid"
)

func main() {
	c := &coolstuff.CoolThing{
		uuid.New(),
		"💯",
	}
	fmt.Println(c)
}
