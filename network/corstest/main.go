package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// go run main.go 8001 https://8002--dogfood-dev--stevenmasley.master.cdr.dev/
	port := os.Args[1]
	link := os.Args[2]
	fmt.Printf("Starting server at port %s with link=%s\n", port, link)
	if err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", port), handler(link)); err != nil {
		log.Fatal(err)
	}
}

func handler(link string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Random", "*")
		fmt.Fprintf(w, `
<!DOCTYPE html>
<html>
<body>

<div id="demo">
<h1>The XMLHttpRequest Object</h1>
<button type="button" onclick="loadDoc()">Change Content</button>
</div>

<script>
function loadDoc() {
	fetch("%[1]s", {
		credentials: "same-origin"
	})
	.then(response => response.json())
  .then(resp => {
    console.log(resp);
  })
.catch(err => console.log(err))
}
</script>

</body>
</html>
	`, link)
	}
}
