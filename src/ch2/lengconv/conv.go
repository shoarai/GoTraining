// Copyright © 2016 shoarai
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package lengconv

// FToM converts a Feet temperature to Meter.
func FToM(f Feet) Meter { return Meter(f / 0.3048) }

// MToF converts a Meter temperature to Feet.
func MToF(m Meter) Feet { return Feet(m * 0.3048) }
