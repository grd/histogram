package histogram

/* range.go
 *
 * Copyright (C) 2013  G.vd.Schoot
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation; either version 2 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this library; if not, write to the
 * Free Software Foundation, Inc., 59 Temple Place - Suite 330,
 * Boston, MA 02111-1307, USA.
 */

// NaturalRange generates a range of natural numbers for a new histogram.
// The Histogram struct works with []float64 for ranges, so this 
// function generates []float64. Also note that the resulting range contains
// one extra position for the final value. 
// For example:  NaturalRange(0, 5, 10) 
// generates []float64{0.0, 10.0, 20.0, 30.0, 40.0, 50.0}
func NaturalRange(start, size, increment uint) (res []float64) {
	res = make([]float64, size+1)
	for i := range res {
		res[i] = float64(start + uint(i)*increment)
	}
	return
}

// Range generates a range of floating point numbers for a new histogram.
// The resulting range contains one extra position for the final value. 
// For example:  Range(0.0, 5, 0.1)
// generates []float64{0.0, 0.1, 0.2, 0.3, 0.4, 0.5}
func Range(start float64, size uint, increment float64) (res []float64) {
	res = make([]float64, size+1)
	for i := range res {
		res[i] = start + float64(i)*increment
	}
	return
}
