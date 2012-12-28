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
	max := h.bin[0]
	for i := range h.bin {
		if h.bin[i] > max {
			max = h.bin[i]
		}
	}
	return max
}

// MaxBin find index of max contents in bins
func (h *Histogram) MaxBin() int {
	var imax int
	max := h.bin[0]
	for i := range h.bin {
		if h.bin[i] > max {
			max = h.bin[i]
			imax = i
		}
	}
	return imax
}

// MinVal find min content values
func (h *Histogram) MinVal() float64 {
	min := h.bin[0]
	for i := range h.bin {
		if h.bin[i] < min {
			min = h.bin[i]
		}
	}
	return min
}

// MinBin find index of min contents in bins
func (h *Histogram) MinBin() int {
	var imin int
	min := h.bin[0]
	for i := range h.bin {
		if h.bin[i] < min {
			min = h.bin[i]
			imin = i
		}
	}
	return imin
}
