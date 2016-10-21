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

	tokens := strings.Split(path, "/")
	jpackage := make([]string, 0)
	for _, toke := range tokens {
		jpackage = append([]string{toke}, jpackage...)
	}
	jpackage = jpackage[1 : len(jpackage)-1]
	gopackage := make([]string, 0)
	for _, name := range jpackage {
		fmt.Println(name)
		if name == "java" || name == "main" || name == "src" {
			break
		}
		gopackage = append(gopackage, name)
	}
	fmt.Println(gopackage)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "public ") {
			fmt.Println(line)
		}
		if strings.HasPrefix(line, "private ") {
			fmt.Println(line)
		}
		if strings.HasPrefix(line, "protected ") {
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
