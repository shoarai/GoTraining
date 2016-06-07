// Copyright © 2016 shoarai

package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"./eval"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/calculate", calculate)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var html = template.Must(template.New("html").Parse(`
	<div>
		<input type="text" name="expr" value="">
		=<span></span>
	</div>
	<table>
		<tr>
			<td onclick="setValue(7)">7</td>
			<td onclick="setValue(8)">8</td>
			<td onclick="setValue(9)">9</td>
			<td onclick="setValue('*')">*</td>
		</tr>
		<tr>
			<td onclick="setValue(4)">4</td>
			<td onclick="setValue(5)">5</td>
			<td onclick="setValue(6)">6</td>
			<td onclick="setValue('-')">-</td>
		</tr>
		<tr>
			<td onclick="setValue(1)">1</td>
			<td onclick="setValue(2)">2</td>
			<td onclick="setValue(3)">3</td>
			<td onclick="setValue()"></td>
		</tr>
		<tr>
			<td onclick="setValue(0)" colspan="3">0</td>
			<td onclick="equal()">=</td>
		</tr>
	</table>
	<script src="https://code.jquery.com/jquery-2.2.4.min.js"></script>
	<script>
		function setValue(val) {
			$input = $('input');
			$input.val($input.val()+val)
		}

		function equal() {
			$.get({
				url: "calculate?expr=" + $('input').val(),
				success: function(ans) {
					$('span').text(ans);
				}
			});
		}
	</script>
`))

func handler(w http.ResponseWriter, r *http.Request) {
	if err := html.Execute(w, ""); err != nil {
		log.Fatal(err)
	}
}

func calculate(w http.ResponseWriter, r *http.Request) {
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
