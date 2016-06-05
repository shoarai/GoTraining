// Copyright © 2016 shoarai

package main

import (
	"log"
	"net/http"
	"strconv"

	"./eval"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	inputExpr := r.URL.Query().Get("expr")

	expr, err := eval.Parse(inputExpr)
	if err != nil {
		w.Write([]byte("Expression parse error"))
		return
	}

	env := eval.Env{}
	for k, v := range r.URL.Query() {
		if k != "expr" {
			val, err := strconv.ParseFloat(v[0], 64)
			if err == nil {
				env[eval.Var(k)] = val
			}
		}
	}

	ans := expr.Eval(env)
	ansStr := strconv.FormatFloat(ans, 'f', 4, 64)

	w.Write([]byte(ansStr))
}
