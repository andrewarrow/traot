package parse

import "os"
import "path/filepath"
import "fmt"
import "strings"
import "io/ioutil"

var files []string

func visit(path string, f os.FileInfo, err error) error {
	if strings.Contains(path, ".java") {
		files = append(files, path)
	}
	return nil
}

func readJava(path string) {
	f, _ := os.Open(path)
	data, _ := ioutil.ReadAll(f)
	str := string(data)
	fmt.Println(str)
}

func Parse(path string) {
	filepath.Walk(path, visit)
	fmt.Println(files[0])
	readJava(files[0])
}
