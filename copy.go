package histogram

/* copy.go
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

var ErrSize = errors.New("histograms have different sizes, cannot copy")
var ErrAlloc = errors.New("failed to allocate space for histogram struct")

// Copy the contents of histogram src into dest 
func (src *Histogram) Copy(dest *Histogram) error {
	if src.Len() != dest.Len() {
		return ErrSize
	}

	copy(dest.range_, src.range_)
	copy(dest.bin, src.bin)

	return nil
}

// Clone an histogram creating an identical new one
func (src *Histogram) Clone() (clone *Histogram, err error) {
	if clone, err = NewHistogramRange(src.range_); err != nil {
		return nil, ErrAlloc
	}

	copy(clone.bin, src.bin)

	return
}
