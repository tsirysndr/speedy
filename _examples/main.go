package main

import (
	"fmt"

	"github.com/tsirysndr/speedy"
)

func main() {
	r, _ := speedy.Start(speedy.OOKLA)
	fmt.Println(r)
}
