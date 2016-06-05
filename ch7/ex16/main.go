// Copyright © 2016 shoarai

package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

var html = `
<div>
	<input type="text" name="answer" value="0">
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
	<td onclick="setValue('+')">+</td>
</tr>
<tr>
	<td onclick="setValue(0)" colspan="3">0</td>
	<td>=</td>
</tr>
</table>

<script>
	function setValue(val) {
		console.log(val);
		document.myForm.myLine.value = myTotal;
	}
</script>
`

func handler(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte(html))

	// var inputExpr string
	// fmt.Println("Input expression:")
	// fmt.Scan(&inputExpr)
	//
	// expr, err := eval.Parse(inputExpr)
	// if err != nil {
	// 	fmt.Println("Expression parse error")
	// 	return
	// }
	//
	// env := eval.Env{}
	//
	// for {
	// 	fmt.Println("Input key name:")
	//
	// 	var key string
	// 	fmt.Scan(&key)
	// 	if key == "" {
	// 		break
	// 	}
	//
	// 	fmt.Printf("Input value of %s:\n", key)
	// 	var val float64
	// 	fmt.Scan(&val)
	//
	// 	env[eval.Var(key)] = val
	// }
	//
	// ans := expr.Eval(env)
	// fmt.Printf("%s = %f\n", expr, ans)
}
