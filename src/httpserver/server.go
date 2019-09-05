package httpserver

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Policy struct {
	ConfigurableHost string   `json:"configurableHost"`
	ApprovedHosts    []string `json:"approvedHosts"`
}

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
<head>
<title>Hello World</title>
</head>
<body>
Hello World!
</body>
</html>`,
	)
}

func policy(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg Policy
	err = json.Unmarshal(b, &msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	output, err := json.Marshal(msg)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(output)
}

func Start() {
	fmt.Println("Starting server...")
	http.HandleFunc("/policy", policy)
	http.ListenAndServe(":9000", nil)
}
