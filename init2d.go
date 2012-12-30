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

func NewHistogram2d(nx, ny int) (*Histogram2d, error) {
	var h Histogram2d

	if nx <= 0 {
		return nil, errors.New("Error: length nx must be positive integer")
	}

	if ny <= 0 {
		return nil, errors.New("Error: length ny must be positive integer")
	}

	h.xrange = make([]float64, nx+1)
	h.yrange = make([]float64, ny+1)
	h.bin = make([]float64, nx*ny)
	return &h, nil
}

func NewHistogram2dUniform(nx, ny int, xmin, xmax, ymin, ymax float64) (
	*Histogram2d, error) {

	if xmin >= xmax {
		return nil, errors.New("xmin must be less than xmax")
	}

	if ymin >= ymax {
		return nil, errors.New("ymin must be less than ymax")
	}

	h, err := NewHistogram2dNatural(nx, ny)

	if err != nil {
		return nil, err
	}

	make_uniform(h.xrange, xmin, xmax)
	make_uniform(h.yrange, ymin, ymax)

	return h, nil
}

// NewHistogramNatural returns a new 2d histogram with ranges of
// natural numbers, starting from 0, an increment of 1, and sizes of nx and ny.
func NewHistogram2dNatural(nx, ny int) (*Histogram2d, error) {
	if nx == 0 {
		return nil, errors.New("histogram2d length nx must be positive integer")
	}

	if ny == 0 {
		return nil, errors.New("histogram2d length ny must be positive integer")
	}

	var h Histogram2d

	h.xrange = make([]float64, nx+1)
	h.yrange = make([]float64, ny+1)
	h.bin = make([]float64, nx*ny)

	for i := range h.xrange {
		h.xrange[i] = float64(i)
	}

	for i := range h.yrange {
		h.yrange[i] = float64(i)
	}

	return &h, nil
}

func (h *Histogram2d) SetRangesUniform(xmin, xmax, ymin, ymax float64) error {
	if xmin >= xmax {
		return errors.New("xmin must be less than xmax")
	}

	if ymin >= ymax {
		return errors.New("ymin must be less than ymax")
	}

	// initialize ranges 
	make_uniform(h.xrange, xmin, xmax)
	make_uniform(h.yrange, ymin, ymax)

	// clear contents 
	for i := range h.bin {
		h.bin[i] = 0.0
	}

	return nil
}

func (h *Histogram2d) SetRanges(xrange, yrange []float64) error {
	if len(xrange) != len(h.xrange) {
		return errors.New("size of xrange must match size of histogram")
	}

	if len(yrange) != len(h.yrange) {
		return errors.New("size of yrange must match size of histogram")
	}

	// initialize ranges 
	copy(h.xrange, xrange)
	copy(h.yrange, yrange)

	// clear contents 
	for i := range h.bin {
		h.bin[i] = 0
	}

	return nil
}

// Routine that create a 2D histogram using the given 
// values for X and Y ranges
func NewHistogram2dRange(xrange, yrange []float64) (*Histogram2d, error) {
	nx, ny := len(xrange)-1, len(yrange)-1
	// check arguments 

	if nx == 0 {
		return nil, errors.New("histogram length nx must be positive integer")
	}

	if ny == 0 {
		return nil, errors.New("histogram length ny must be positive integer")
	}

	// init ranges 

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
	h.bin = make([]float64, nx*ny)

	// init histogram 

	// init ranges 
	copy(h.xrange, xrange)
	copy(h.yrange, yrange)

	return &h, nil
}
