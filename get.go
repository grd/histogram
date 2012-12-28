package histogram

/* get.go
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

var indexErr = "Error: index out of range. %v <> [0..%v]"

func (h *Histogram) Get(i int) float64 {
	if i < 0 || i >= len(h.Bin) {
		log.Fatalf(indexErr, i, len(h.Bin)-1)
	}

	return h.Bin[i]
}

func (h *Histogram) GetRange(i int) (lower, upper float64) {
	if i < 0 || i >= len(h.Bin) {
		log.Fatalf(indexErr, i, len(h.Bin)-1)
	}

	lower = h.Range[i]
	upper = h.Range[i+1]

	return
}

func (h *Histogram) Find(x float64) (int, error) {
	return find(h.Range, x)
}
