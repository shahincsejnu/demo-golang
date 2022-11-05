package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Demo struct {
	Names []string
}

type Inner struct {
	Name string
}

type Outer struct {
	Inner
}

type GenerateGoCode struct {
	Package    string
	StructName string
	Fileds     map[string]string
}

func main() {
	// tmplString := "Hello {{ . }}!\n"

	// tmpl, _ := template.New("demo").Parse("Hello {{ . }}!\n")
	// tmpl2 := template.New("demo2")

	// tmpl.Parse("Hello {{ . }}!\n")
	// tmpl2.Parse(tmplString)

	// tmpl.Execute(os.Stdout, "Sahadat")
	// tmpl.Execute(os.Stdout, 1)

	// tmpl2.Execute(os.Stdout, "Another")

	// tmpl2.ExecuteTemplate(os.Stdout, "demo2", "Nothing")

	// tmpl2 := template.New()
	// tmpl2, err := tmpl2.Parse()

	// template.Must(template.New().Parse())

	// tmpl := template.Must(template.New("t1").Parse(`
	// 	{{ define "t1" }}
	// 	t1
	// 	{{ end }}
	// `))

	// tmpl = template.Must(tmpl.New("t2").Parse(`
	// 	t2
	// 	{{ template "t1" }}
	// `))

	// tmpl.ExecuteTemplate(os.Stdout, "t2", nil)

	// tmplString := "This is: {{ if . }} True {{ else }} False {{ end }}"

	// tmpl := template.Must(template.New("demo").Parse(tmplString))

	// tmpl.Execute(os.Stdout, 1)

	// 	fncMap := map[string]interface{}{
	// 		"toUpper": strings.ToUpper,
	// 	}

	// 	tmpl, _ := template.New("demo").Funcs(fncMap).Parse(`
	// {{- range .Names -}}
	// Hello {{ toUpper . }}, how are you?
	// {{ end -}}
	// 	`)

	// tmpl := template.Must(template.ParseFiles("./temp.tmpl"))

	// names := []string{"first"}

	// demo := Demo{Names: names}

	// tmpl.Execute(os.Stdout, demo)

	// outer := Outer{
	// 	Inner: Inner{
	// 		Name: "Nothing",
	// 	},
	// }

	// tmpl := template.Must(template.New("demo").Parse("Hello {{ .Name }}"))

	// tmpl.Execute(os.Stdout, outer)

	tmpl := template.Must(template.ParseFiles("./temp2.tmpl"))

	buf := new(bytes.Buffer)

	tmpl.Execute(buf, 100)

	fmt.Println(buf.String())
}
