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

var errRange = "Value out of range. %v <> [%v..%v]"

func find(Range []float64, x float64) (res int, err error) {
	n := len(Range) - 1

	res = -1
	for i := range Range {
		if x >= Range[i] {
			res = i
		}
	}
	if res < 0 || res == n {
		err = fmt.Errorf(errRange, x, Range[0], Range[n])
	}

	return
}

func findInt(Range []int, x int) (res int, err error) {
	n := len(Range) - 1

	res = -1
	for i := range Range {
		if x >= Range[i] {
			res = i
		}
	}
	if res < 0 || res == n {
		err = fmt.Errorf(errRange, x, Range[0], Range[n])
	}

	return
}
