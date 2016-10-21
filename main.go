package main

import "os"
import "fmt"
import "github.com/andrewarrow/traot/parse"

func main() {
	val := os.Args[1]
	fmt.Println(val)
	os.Mkdir(fmt.Sprintf("%s_go", val), os.ModePerm)
	parse.Parse(val)
}
