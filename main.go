package main

import "os"
import "fmt"
import "github.com/andrewarrow/traot/parse"

func main() {
	val := os.Args[1]
	fmt.Println(val)
	parse.Parse(val)
}
