package main

import (
	"demo/OYK2023GoLang/Tests/greet"
	"fmt"
)

func main() {
	fmt.Println(greet.SayHi())       // hi everybody!
	fmt.Println(greet.SayHi("uğur")) // hi uğur!
	fmt.Println(greet.SayHi("uğur", "erhan"))
	// hi uğur!
	// hi erhan!
}
