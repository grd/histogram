package histogram

/* maxval.go
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

// MaxVal find max content values
func (h *Histogram) MaxVal() float64 {
	max := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] > max {
			max = h.Bin[i]
		}
	}
	return max
}

// MaxBin find index of max contents in bins
func (h *Histogram) MaxBin() int {
	var imax int
	max := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] > max {
			max = h.Bin[i]
			imax = i
		}
	}
	return imax
}

// MinVal find min content values
func (h *Histogram) MinVal() float64 {
	min := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] < min {
			min = h.Bin[i]
		}
	}
	return min
}

// MinBin find index of min contents in bins
func (h *Histogram) MinBin() int {
	var imin int
	min := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] < min {
			min = h.Bin[i]
			imin = i
		}
	}
	return imin
}
