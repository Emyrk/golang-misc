package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// go run main.go 8002 https://8001--dogfood-dev--stevenmasley.master.cdr.dev/
	port := os.Args[1]
	link := os.Args[2]
	fmt.Printf("Starting server at port %s with link=%s\n", port, link)
	mux := http.NewServeMux()
	mux.Handle("/", handler(link))
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, `{"ok":true}`)
	})
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "https://8001--dogfood-dev--stevenmasley.master.cdr.dev/")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// w.Header().Set("Access-Control-Allow-Headers", "*")
		// w.Header().Set("Vary", "Origin")
		o := r.Header.Get("Origin")
		if o != "" {
			w.Header().Set("Access-Control-Allow-Origin", o)
		}

		mux.ServeHTTP(w, r)
	})); err != nil {
		log.Fatal(err)
	}
}

func handler(link string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<body>

<div id="demo">
<h1>Fetch %[1]s</h1>
<div id="content"></div>
<button type="button" onclick="loadDoc()">Click me</button>
</div>

<script>
function loadDoc() {
	fetch("%[1]s", {
		credentials: "include"
	})
  .then(resp => {
    console.log(resp);
		document.getElementById("content").innerHTML = "It worked! Make sure the devurls are not public for testing."
  })
.catch(err => {
	console.log(err);
	document.getElementById("content").innerHTML = "It failed, check the console logs"
})
}
</script>

</body>
</html>
	`, link)
	}
}
