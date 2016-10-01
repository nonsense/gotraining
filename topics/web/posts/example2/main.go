package main

import (
	"fmt"
	"log"
	"net/http"
)

func router(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getHandler(res, req)
	case "POST":
		postHandler(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html")
	res.Write([]byte(`
<form action="/" method="POST">
<p>
	<input type="text" name="FirstName" placeholder="First Name" />
</p>
<p>
	<input type="text" name="LastName" placeholder="Last Name" />
</p>
<p>
	<input type="submit" value="CLICK ME!!" />
</p>
</form>
	`))
}

func postHandler(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(res, err)
		return
	}
	form := req.Form
	fmt.Fprintf(res, "First Name: %s\nLast Name %s", form.Get("FirstName"), form.Get("LastName"))
}

func main() {
	http.HandleFunc("/", router)
	log.Fatal(http.ListenAndServe(":3000", nil))
}