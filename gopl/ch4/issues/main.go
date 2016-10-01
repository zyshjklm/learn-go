// Issues prints a tble of GitHub issues matching the search terms
package main 

import (
	"fmt"
	"log"
	"os"
	"github.com/jungle85gopy/learn-go/gopl/ch4/github"

)

/*
refer:
	https://developer.github.com/v3/search/#search-issues
os.Args like:
	windows label:bug language:python state:open

example:
	go run issues/main.go linux language:go

*/

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)

	for _, item := range result.Items {	
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

// output:
/*
q:  linux+language%3Ago+state%3Aopen
7258 issues:
#10      ianomad Windows/Linux build
#117     mskogly Linux 32
#2        jbenet Linux?
#128   mveytsman Defaults for Amazon Linux
*/
