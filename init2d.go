package histogram

/* init2d.go
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

// Routine that create a 2D histogram using the given 
// values for X and Y ranges
func NewHistogram2d(xrange, yrange []float64) (*Histogram2d, error) {
	nx, ny := len(xrange)-1, len(yrange)-1

	// check arguments 

	if nx == 0 {
		return nil, errors.New("histogram length nx must be positive integer")
	}

	if ny == 0 {
		return nil, errors.New("histogram length ny must be positive integer")
	}

	// check ranges 

	for i := 0; i < nx; i++ {
		if xrange[i] >= xrange[i+1] {
			return nil, errors.New("histogram xrange not in increasing order")
		}
	}

	for j := 0; j < ny; j++ {
		if yrange[j] >= yrange[j+1] {
			return nil, errors.New("histogram yrange not in increasing order")
		}
	}

	// Allocate histogram  

	var h Histogram2d
	h.xrange = make([]float64, nx+1)
	h.yrange = make([]float64, ny+1)
	h.bin = make([]int, nx*ny)

	// init ranges 

	copy(h.xrange, xrange)
	copy(h.yrange, yrange)

	return &h, nil
}
