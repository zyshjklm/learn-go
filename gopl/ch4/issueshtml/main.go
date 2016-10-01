package main 

import (
	
	"log"
	"os"
	"html/template"
	"github.com/jungle85gopy/learn-go/gopl/ch4/github"

)

const templ = `
<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
</tr>
{{range .Items}}
<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>`


func main() {
	var report = template.Must(
		template.New("issuelist").
		Parse(templ))

	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}

/*
// output:
go run issuesreport/main.go repo:golang/go is:open json decoder

16 issues:
--------------------------------
Number: 11046
User: kurin
Title: encoding/json: Decoder internally buffers full input
Age: 464 days
--------------------------------
Number: 15314
User: okdave
Title: proposal: some way to reject unknown fields in encoding/json.Dec
Age: 147 days
*/
