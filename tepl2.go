package main

import (
	"html/template"
	"os"
)

type Friend struct {
	Fname string
}

type Person struct {
	UserName string
	Emails   []string
	Friends  []*Friend
}

func main() {
	f1 := Friend{Fname: "Tobin"}
	f2 := Friend{Fname: "Chris"}
	t := template.New("example2")
	t, _ = t.Parse(`hello {{.UserName}}!
		{{range .Emails}}
			an email {{.}}
		{{end}}
		{{with .Friends}}
		{{range .}}
			my friend name is {{.Fname}}
		{{end}}
		{{end}}
		`)
	p := Person{UserName: "Danqi",
		Emails:  []string{"danqi@gmail.com", "dandan@gmail.com"},
		Friends: []*Friend{&f1, &f2}}
	t.Execute(os.Stdout, p)

}
