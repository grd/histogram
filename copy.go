package histogram

/* histogram/copy.go
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
/***************************************************************
 *
 * File gsl_histogram_copy.c: 
 * Routine to copy an histogram. 
 * Need GSL library and headers.
 *
 * Author: S. Piccardi
 * Jan. 2000
 *
 ***************************************************************/

import (
	"errors"
)

// Copy the contents of an Histogram into another
func Copy(dest, src *Histogram) error {
	if len(dest.Bin) != len(src.Bin) {
		return errors.New("histograms have different sizes, cannot copy")
	}

	copy(dest.Range, src.Range)
	copy(dest.Bin, src.Bin)

	return nil
}

// Clone an Histogram creating an identical new one
func Clone(src *Histogram) (*Histogram, error) {
	h, err := NewHistogramRange(src.Range)

	if err != nil {
		return nil, errors.New("failed to allocate space for histogram struct")
	}

	copy(h.Bin, src.Bin)

	return h, nil
}
