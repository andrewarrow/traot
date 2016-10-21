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

func readJava(path, orig string) {
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
		if name == "java" || name == "main" || name == "src" {
			break
		}
		gopackage = append(gopackage, name)
	}
	fmt.Println(gopackage)
	gopackname := strings.Join(gopackage, "_")
	dirname := fmt.Sprintf("%s_go/%s", orig, gopackname)
	os.Mkdir(dirname, os.ModePerm)
	endparts := strings.Split(strings.ToLower(tokens[len(tokens)-1]), ".")
	endpart := endparts[0]
	filename := fmt.Sprintf("%s_go/%s/%s", orig, gopackname, endpart+".go")
	fmt.Println(filename)

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "public ") {
			//fmt.Println(line)
		}
		if strings.HasPrefix(line, "private ") {
			//fmt.Println(line)
		}
		if strings.HasPrefix(line, "protected ") {
			//fmt.Println(line)
		}
	}
}

func Parse(path string) {
	filepath.Walk(path, visit)
	for i, file := range files {
		readJava(file, path)
		if i > 5 {
			break
		}
	}
}
