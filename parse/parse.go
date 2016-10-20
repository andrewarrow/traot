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
	lines := strings.Split(str, "\n")

	for _, line := range lines {
		if strings.HasPrefix(line, "public class ") {
			fmt.Println(line)
		}
	}
}

func Parse(path string) {
	filepath.Walk(path, visit)
	for i, file := range files {
		readJava(file)
		if i > 5 {
			break
		}
	}
}
