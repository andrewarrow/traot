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
		//fmt.Println(filename)

		wfile, _ := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0777)
		wfile.WriteString("package foo\n\n")

		for _, line := range lines {
			line = strings.TrimSpace(line)
			handleLine(wfile, line)
		}

		wfile.Close()
	}
}
func handleLine(wfile *os.File, line string) {
	hasLevel := false
	hasEqual := false
	hasParen := false
	if strings.HasPrefix(line, "public ") || strings.HasPrefix(line, "private ") || strings.HasPrefix(line, "protected ") {
		hasLevel = true
	}

	if strings.Contains(line, "=") {
		hasEqual = true
	}
	if strings.Contains(line, "(") {
		hasParen = true
	}

	if hasLevel && !hasEqual && hasParen {
		name := ""
		tokens := strings.Split(line, " ")
		for _, t := range tokens {
			if strings.Contains(t, "(") {
				name = t
				break
			}
		}
		flip := false
		params := make([]string, 0)
		for _, t := range tokens {
			if strings.Contains(t, "(") {
				flip = true
			}
			if flip {
				//[CopyNoChildren()]
				//[filterVideoContentOnList(List<VideoContentVO> contentList, List<Long> userIdList)]
				//[filterSeasonContentOnList(List<SeriesContentVO> contentList, List<Long> userIdList)]
				entry := t
				if strings.Contains(t, "(") {
					inside := strings.Split(t, "(")
					entry = inside[1]
				}
				if strings.Contains(entry, ")") {
					inside := strings.Split(entry, ")")
					entry = inside[0]
				}

				params = append(params, entry)
			}
			if strings.Contains(t, ")") {
				break
			}
		}
		fmt.Println(params)
		params = strings.Split(strings.Join(params, " "), ",")
		plist := make([]string, 0)
		if len(params) == 0 {
			for _, t := range params {
				t = strings.TrimSpace(t)
				inside := strings.Split(t, " ")
				plist = append(plist, inside[1]+" "+inside[0])
			}
		}
		tokens = strings.Split(name, "(")
		name = tokens[0]

		wfile.WriteString("func " + name + "(")
		wfile.WriteString(strings.Join(plist, ", "))
		wfile.WriteString(") {\n")
		wfile.WriteString("}\n")
	}
}

func Parse(path string) {
	filepath.Walk(path, visit)
	for _, file := range files {
		readJava(file, path)
	}
}
