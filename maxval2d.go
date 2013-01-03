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

// Max finds first value and index of max contents in bins
func (h *Histogram2d) Max() (xmax, ymax, imax int) {
	nx := h.LenX()
	ny := h.LenY()
	imax = h.bin[0]

	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			x := h.bin[i*ny+j]

			if x > imax {
				imax = x
				xmax = i
				ymax = j
			}
		}
	}
	return
}

// Min finds first value and index of min contents in bins
func (h *Histogram2d) Min() (xmin, ymin, imin int) {
	nx := h.LenX()
	ny := h.LenY()
	imin = h.bin[0]

	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			x := h.bin[i*ny+j]

			if x < imin {
				imin = x
				xmin = i
				ymin = j
			}
		}
	}
	return
}
