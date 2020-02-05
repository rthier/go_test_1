package main

import "fmt"
import "encoding/json"
import "github.com/google/go-cmp/cmp"

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	fmt.Println("Hello, world.")
	b, _ := json.Marshal(Message{"Alice", "Hello", 1294706395881547000})
	fmt.Println(b)
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
