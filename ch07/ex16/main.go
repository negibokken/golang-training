package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"./eval"
)

type expr struct {
	s string
}

func main() {
	http.HandleFunc("/calc", execute)
	fmt.Println("Listening localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var htmlTemplate = `
<html>
<head>
<meta charset="utf-8"/>
<title>calculator</title>
</head>
<body>
{{.}}
</body>
</html>
`

func execute(w http.ResponseWriter, req *http.Request) {
	exp := req.URL.Query().Get("expr")
	expr, err := eval.Parse(string(exp))
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	}
	str := fmt.Sprintf("%v", expr.Eval(eval.Env{}))
	test := struct {
		expr string
		calc string
	}{
		exp,
		str,
	}
	t := template.Must(template.New("t").Parse(htmlTemplate))
	t.Execute(w, test)
}
