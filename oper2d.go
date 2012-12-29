package histogram

/* oper2d.go
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

// EqualRanges control if two histograms have the same ranges
func (h1 *Histogram2d) EqualRanges(h2 *Histogram2d) bool {
	if h1.LenX() != h2.LenX() || h1.LenY() != h2.LenY() {
		return false
	}

	for i := range h1.xrange {
		if h1.xrange[i] != h2.xrange[i] {
			return false
		}
	}
	for i := range h1.yrange {
		if h1.yrange[i] != h2.yrange[i] {
			return false
		}
	}

	return true
}

// Add histogram h1 with h2 
func (h1 *Histogram2d) Add(h2 *Histogram2d) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] += h2.bin[i]
	}

	return nil
}

// Subtract histogram h1 with h2
func (h1 *Histogram2d) Sub(h2 *Histogram2d) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] -= h2.bin[i]
	}

	return nil
}

// Multiply histogram h1 with h2
func (h1 *Histogram2d) Mul(h2 *Histogram2d) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] *= h2.bin[i]
	}

	return nil
}

// Divide histogram h1 with h2
func (h1 *Histogram2d) Div(h2 *Histogram2d) error {
	if !h1.EqualRanges(h2) {
		return difRangeErr
	}

	for i := range h1.bin {
		h1.bin[i] /= h2.bin[i]
	}

	return nil
}

// Scale histogram by a numeric factor
func (h *Histogram2d) Scale(scale float64) {
	for i := range h.bin {
		h.bin[i] *= scale
	}
}

// Shift histogram by a numeric offset
func (h *Histogram2d) Shift(shift float64) {
	for i := range h.bin {
		h.bin[i] += shift
	}
}
