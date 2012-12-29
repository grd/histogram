package histogram

/* add2d.go
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

func (h *Histogram2d) Increment(x, y float64) error {
	return h.Accumulate(x, y, 1.0)

}

func (h *Histogram2d) Accumulate(x, y, weight float64) error {
	nx := h.LenX()
	ny := h.LenY()

	i, j, err := find2d(h.xrange, h.yrange, x, y)

	if err != nil {
		return err
	}

	if i >= nx {
		return errors.New("index lies outside valid _range of 0 .. nx - 1")
	}

	if j >= ny {
		return errors.New("index lies outside valid _range of 0 .. ny - 1")
	}

	h.bin[i*ny+j] += weight

	return nil
}
