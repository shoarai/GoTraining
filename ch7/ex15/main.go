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

	env := eval.Env{}

	for {
		fmt.Println("Input key name:")

		var key string
		fmt.Scan(&key)
		if key == "" {
			break
		}

		fmt.Printf("Input value of %s:\n", key)
		var val float64
		fmt.Scan(&val)

		env[eval.Var(key)] = val
	}

	ans := expr.Eval(env)
	fmt.Printf("%s = %f\n", expr, ans)
}
