// server_header echoes the HTTP request

package main 

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// each request calls handler
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "REQ: %s %s %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "  Header[%q] = %q\n", k, v)
	}

	fmt.Fprintf(w, "Host: %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm() ; err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

/*

REQ: GET /hello?q=query&w=jungle HTTP/1.1
  Header["User-Agent"] = ["Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.79 Safari/537.36"]
  Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,* /*;q=0.8"]
  Header["Accept-Encoding"] = ["gzip, deflate, sdch"]
  Header["Accept-Language"] = ["zh-CN,zh;q=0.8,en;q=0.6"]
  Header["Connection"] = ["keep-alive"]
  Header["Upgrade-Insecure-Requests"] = ["1"]
Host: "localhost:8000"
RemoteAddr = "127.0.0.1:54240"
Form["q"] = ["query"]
Form["w"] = ["jungle"]

*/
