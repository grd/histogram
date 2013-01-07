package histogram

/* init.go
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
	"errors"
)

func NewHistogram(Range []float64) (*Histogram, error) {
	var h Histogram
	n := len(Range) - 1

	// check arguments 
	if n <= 0 {
		return nil, errors.New("histogram length n must be positive integer")
	}

	// check ranges 
	for i := 0; i < n; i++ {
		if Range[i] >= Range[i+1] {
			return nil, errors.New("Histogram range must be in increasing order")
		}
	}

	// Allocate histogram  
	h.range_ = make([]float64, n+1)
	h.bin = make([]int, n)

	// initialize Ranges 
	copy(h.range_, Range)

	return &h, nil
}
