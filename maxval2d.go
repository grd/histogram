package histogram

/* maxval2d.go
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

// MaxVal returns the maximum contents value of a 2D histogram
func (h *Histogram2d) MaxVal() (max float64) {
	max = h.bin[0]

	for _, val := range h.bin {
		if val > max {
			max = val
		}
	}
	return
}

// MaxBin finds the bin index for maximum value of a 2D histogram
func (h *Histogram2d) MaxBin() (xmax, ymax int) {
	nx := h.LenX()
	ny := h.LenY()
	max := h.bin[0]

	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			x := h.bin[i*ny+j]

			if x > max {
				max = x
				xmax = i
				ymax = j
			}
		}
	}
	return
}

// MinVal returns the minimum contents value of a 2D histogram
func (h *Histogram2d) MinVal() (min float64) {
	min = h.bin[0]

	for _, val := range h.bin {
		if val < min {
			min = val
		}
	}
	return
}

// MinBin finds the bin index for minimum value of a 2D histogram
func (h *Histogram2d) MinBin() (xmin, ymin int) {
	nx := h.LenX()
	ny := h.LenY()
	min := h.bin[0]

	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			x := h.bin[i*ny+j]

			if x < min {
				min = x
				xmin = i
				ymin = j
			}
		}
	}
	return
}
