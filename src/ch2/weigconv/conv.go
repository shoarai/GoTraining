// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package weigconv

// PToKG converts a Pound temperature to Kilogram.
func PToKG(p Pound) Kilogram { return Kilogram(p * 0.45359237) }

// KGToP converts a Kilogram temperature to Pound.
func KGToP(k Kilogram) Pound { return Pound(k / 0.45359237) }