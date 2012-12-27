package histogram

/* histogram/init.go
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

func NewHistogram(n int) (*Histogram, error) {
	var h Histogram

	if n <= 0 {
		return nil, errors.New("histogram length n must be positive integer")
	}

	h.Range = make([]float64, n+1)
	h.Bin = make([]float64, n)

	return &h, nil
}

func make_uniform(Range []float64, xmin, xmax float64) {
	//
	// Simplified calculation. (2012 G.v.d.Schoot)
	//
	incr := (xmax - xmin) / float64(len(Range)-1)
	for i := range Range {
		Range[i] = xmin + float64(i)*incr
	}
}

func NewHistogramUniform(n int, xmin, xmax float64) (*Histogram, error) {
	if xmin >= xmax {
		return nil, errors.New("xmin must be less than xmax")
	}

	h, err := NewHistogram(n)

	if err != nil {
		return h, err
	}

	make_uniform(h.Range, xmin, xmax)

	return h, nil
}

func NewHistogramIncr(n int) (*Histogram, error) {
	h, err := NewHistogram(n)

	if err != nil {
		return h, err
	}

	for i := range h.Range {
		h.Range[i] = float64(i)
	}

	return h, nil
}

//  These initialization functions suggested by Achim Gaedke 

func (h *Histogram) SetRangesUniform(xmin, xmax float64) error {
	if xmin >= xmax {
		return errors.New("xmin must be less than xmax")
	}

	//  initialize Ranges 
	make_uniform(h.Range, xmin, xmax)

	//  clear contents 
	for i := range h.Bin {
		h.Bin[i] = 0
	}

	return nil
}

func (h *Histogram) SetRanges(Range []float64) error {
	if len(h.Range) != len(Range) {
		return errors.New("size of range must match size of histogram")
	}

	//  initialize ranges 
	copy(h.Range, Range)

	//  clear contents 
	for i := range h.Bin {
		h.Bin[i] = 0
	}

	return nil
}

func NewHistogramRange(Range []float64) (*Histogram, error) {
	var h Histogram
	n := len(Range) - 1

	// check arguments 
	if n <= 0 {
		return nil, errors.New("histogram length n must be positive integer")
	}

	// check ranges 
	for i := 0; i < n; i++ {
		if Range[i] >= Range[i+1] {
			return nil, errors.New("Histogram.Range must be in increasing order")
		}
	}

	// Allocate histogram  
	h.Range = make([]float64, n+1)
	h.Bin = make([]float64, n)

	// initialize Ranges 
	copy(h.Range, Range)

	return &h, nil
}
