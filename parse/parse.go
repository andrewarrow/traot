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
	//fmt.Println(path)
	f, _ := os.Open(path)
	data, _ := ioutil.ReadAll(f)
	str := string(data)
	lines := strings.Split(str, "\n")

	tokens := strings.Split(path, "/")
	jpackage := make([]string, 0)
	for _, toke := range tokens {
		jpackage = append(jpackage, strings.ToLower(toke))
	}
	jpackage = jpackage[0 : len(jpackage)-1]
	gopackage := make([]string, 0)
	flip := false
	for _, name := range jpackage {
		if flip {
			gopackage = append(gopackage, name)
		}
		if name == "java" {
			flip = true
		}
	}
	//fmt.Println(gopackage)
	if len(gopackage) > 2 {
		first := strings.Join(gopackage[1:3], "_")
		//fmt.Println(first)
		dirname := fmt.Sprintf("%s_go/%s", orig, first)
		os.Mkdir(dirname, 0777)

		endparts := strings.Split(strings.ToLower(tokens[len(tokens)-1]), ".")
		endpart := endparts[0]

		rest := gopackage[2:len(gopackage)]
		reststr := ""
		if len(rest) > 1 {
			reststr = strings.Join(rest[1:len(rest)], "_")
		}
		filename := fmt.Sprintf("%s_go/%s/%s", orig, first, reststr+"_"+endpart+".go")
		fmt.Println(filename)

		wfile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "public ") {
				wfile.WriteString("//" + line + "\n")
			}
			if strings.HasPrefix(line, "private ") {
				wfile.WriteString("//" + line + "\n")
			}
			if strings.HasPrefix(line, "protected ") {
				wfile.WriteString("//" + line + "\n")
			}
		}

		wfile.Close()
	}
}

func Parse(path string) {
	filepath.Walk(path, visit)
	for _, file := range files {
		readJava(file, path)
	}
}
