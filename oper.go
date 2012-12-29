package histogram

/* oper.go
 *
 * Copyright (C) 2000  Simone Piccardi
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

import (
	"errors"
)

var difRangeErr = errors.New("histograms have different ranges")

// EqualRanges control if two histograms have the same ranges
func (h1 *Histogram) EqualRanges(h2 *Histogram) bool {
	if h1.Len() != h2.Len() {
		return false
	}

	for i := range h1.range_ {
		if h1.range_[i] != h2.range_[i] {
			return false
		}
	}

	return true
}

// Add two histograms
func (h1 *Histogram) Add(h2 *Histogram) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] += h2.bin[i]
	}
	return nil
}

// Subtract two histograms
func (h1 *Histogram) Sub(h2 *Histogram) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] -= h2.bin[i]
	}
	return nil
}

// Multiply two histograms
func (h1 *Histogram) Mul(h2 *Histogram) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] *= h2.bin[i]
	}
	return nil
}

// Divide two histograms
func (h1 *Histogram) Div(h2 *Histogram) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] /= h2.bin[i]
	}
	return nil
}

// Scale a histogram by a numeric factor 
func (h *Histogram) Scale(scale float64) {
	for i := range h.bin {
		h.bin[i] *= scale
	}
}

// Shift a histogram by a numeric offset 
func (h *Histogram) Shift(shift float64) {
	for i := range h.bin {
		h.bin[i] += shift
	}
}
