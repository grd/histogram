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

// Copy the contents of Histogram src into dest 
func (src *Histogram) Copy(dest *Histogram) error {
	if len(dest.bin) != len(src.bin) {
		return errors.New("histograms have different sizes, cannot copy")
	}

	copy(dest.range_, src.range_)
	copy(dest.bin, src.bin)

	return nil
}

// Clone an Histogram creating an identical new one
func (src *Histogram) Clone() (clone *Histogram, err error) {
	clone, err = NewHistogramRange(src.range_)

	if err != nil {
		return nil, errors.New("failed to allocate space for histogram struct")
	}

	copy(clone.bin, src.bin)

	return
}
