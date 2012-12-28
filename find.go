package histogram

/* find.go
 * 
 * Copyright (C) 1996, 1997, 1998, 1999, 2000 Brian Gough
 * 
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or (at
 * your option) any later version.
 * 
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 * 
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.
 */

import (
	"fmt"
)

var rangeErr = "Value out of range. %v <> [%v..%v]"

func find(Range []float64, x float64) (int, error) {
	var i_linear, lower, upper, mid, i int
	n := len(Range) - 1

	if x < Range[0] || x >= Range[n] {
		return 0, fmt.Errorf(rangeErr, x, Range[0], Range[n])
	}

	// optimize for linear case

	{
		u := (x - Range[0]) / (Range[n] - Range[0])
		i_linear = int(u * float64(n))
	}

	if x >= Range[i_linear] && x < Range[i_linear+1] {
		i = i_linear
		return i, nil
	}

	// perform.binary search

	upper = n
	lower = 0

	for upper-lower > 1 {
		mid = (upper + lower) / 2

		if x >= Range[mid] {
			lower = mid
		} else {
			upper = mid
		}
	}

	i = lower

	// sanity check the result

	if x < Range[lower] || x >= Range[lower+1] {
		return 0, fmt.Errorf(rangeErr, x, Range[lower], Range[lower+1])
	}

	return i, nil
}
