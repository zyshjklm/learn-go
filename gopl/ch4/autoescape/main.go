package main 

import (
	"log"
	"os"
	"html/template"
)

func main() {
	const templ = `<p>A: {{.A}}</p><p>B: {{.B}}</p>`
	var data struct {
		A string 			// untrusted plain text
		B template.HTML 	// trusted HTML
	}

	t := template.Must(template.New("escape").
		Parse(templ))

	data.A = "<b>Hello!</b>"
	data.B = "<b>Hello!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

/*

output as text:
<p>A: &lt;b&gt;Hello!&lt;/b&gt;</p><p>B: <b>Hello!</b></p>%

as it appears in a browser:

A: <b>Hello!</b>
B: Hello!

*/