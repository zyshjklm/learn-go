package main 

import (
	
	"log"
	"os"
	"time"
	"text/template"
	"github.com/jungle85gopy/learn-go/gopl/ch4/github"

)

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------------
Number: {{.Number}}
User: {{.User.Login}}
Title: {{.Title | printf "%.64s"}}
Age: {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	var report = template.Must(
		template.New("issuelist").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
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
