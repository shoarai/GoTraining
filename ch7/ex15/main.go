// Copyright © 2016 shoarai

package main

import "fmt"
import "./eval"

func main() {
	var inputExpr string
	fmt.Println("Input expression:")
	fmt.Scan(&inputExpr)

	expr, err := eval.Parse(inputExpr)
	if err != nil {
		fmt.Println("Expression parse error")
		return
	}

	env := eval.Env{"x": 0, "y": 0}
	for k, _ := range env {
		fmt.Printf("Input %s:\n", k)
		var f float64
		fmt.Scan(&f)
		env[k] = f
	}

	ans := expr.Eval(env)
	fmt.Printf("%s, %s = %f\n", expr, env, ans)
}
