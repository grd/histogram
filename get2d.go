package histogram

/* get2d.go
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
	"log"
)

func (h *Histogram2d) Get(i, j int) float64 {
	nx := h.LenX()
	ny := h.LenY()

	if i < 0 || i >= nx {
		log.Fatalf(indexErr, i, nx)
	}

	if j < 0 || j >= ny {
		log.Fatalf(indexErr, j, ny)
	}

	return h.bin[i*ny+j]
}

func (h *Histogram2d) GetXrange(i int) (xlower, xupper float64) {
	nx := h.LenX()

	if i < 0 || i >= nx {
		log.Fatalf(indexErr, i, nx)
	}

	xlower = h.xrange[i]
	xupper = h.xrange[i+1]
	return
}

func (h *Histogram2d) GetYrange(j int) (ylower, yupper float64) {
	ny := h.LenY()

	if j < 0 || j >= ny {
		log.Fatalf(indexErr, j, ny)
	}

	ylower = h.yrange[j]
	yupper = h.yrange[j+1]
	return
}

func (h *Histogram2d) Find(x, y float64) (i, j int, err error) {
	if i, err = find(h.xrange, x); err != nil {
		return
	}

	// if statement is not really required...
	if j, err = find(h.yrange, y); err != nil {
		return
	}

	return
}
