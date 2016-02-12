// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Main converts its numeric argument to temperature, length and weight.
package main

import (
	"fmt"
	"os"
	"strconv"

	"ch2/tempconv"
	"ch2/lengconv"
	"ch2/weigconv"
)

func main() {
	input := os.Args[1:]
	if len(input) == 0 {
		var str string
		fmt.Scan(&str)
		input = make([]string, 1)
		input[0] = str
	}

	for _, arg := range input {
		val, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		tf := tempconv.Fahrenheit(val)
		tc := tempconv.Celsius(val)
		fmt.Printf("%s = %s, %s = %s\n",
			tf, tempconv.FToC(tf), tc, tempconv.CToF(tc))

		lf := lengconv.Feet(val)
		lm := lengconv.Meter(val)
		fmt.Printf("%s = %s, %s = %s\n",
			lf, lengconv.FToM(lf), lm, lengconv.MToF(lm))

		lp := weigconv.Pound(val)
		lk := weigconv.Kilogram(val)
		fmt.Printf("%s = %s, %s = %s\n",
			lp, weigconv.PToKG(lp), lk, weigconv.KGToP(lk))
	}
}