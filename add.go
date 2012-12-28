package histogram

/* add.go
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

func (h *Histogram) Increment(x float64) error {
	return h.Accumulate(x, 1.0)
}

func (h *Histogram) Accumulate(x, weight float64) error {
	n := len(h.Bin)
	index, err := find(h.Range, x)

	if err != nil {
		return err
	}

	if index >= n {
		return errors.New("index lies outside valid Range of 0 .. n - 1")
	}

	h.Bin[index] += weight

	return nil
}
