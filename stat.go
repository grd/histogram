package histogram

/* stat.go
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

// FIXME: We skip negative values in the histogram h.bin[i] < 0,
// since those correspond to negative weights (BJG)

import (
	"math"
)

// Mean compute the bin-weighted arithmetic mean of the histogram
func (h *Histogram) Mean() float64 {
	n := h.Len()

	var wmean, W float64

	for i := 0; i < n; i++ {
		xi := (h.range_[i+1] + h.range_[i]) / 2
		wi := h.bin[i]

		if wi > 0 {
			W += wi
			wmean += (xi - wmean) * (wi / W)
		}
	}

	return wmean
}

// Sigma compute the bin-weighted sigma of the histogram
func (h *Histogram) Sigma() float64 {
	n := h.Len()

	var wvariance, wmean, W float64

	// FIXME: should use a single pass formula here, as given in
	// N.J.Higham 'Accuracy and Stability of Numerical Methods', p.12 

	// Compute the mean

	for i := 0; i < n; i++ {
		xi := (h.range_[i+1] + h.range_[i]) / 2
		wi := h.bin[i]

		if wi > 0 {
			W += wi
			wmean += (xi - wmean) * (wi / W)
		}
	}

	// Compute the variance

	W = 0.0

	for i := 0; i < n; i++ {
		xi := ((h.range_[i+1]) + (h.range_[i])) / 2
		wi := h.bin[i]

		if wi > 0 {
			delta := (xi - wmean)
			W += wi
			wvariance += (delta*delta - wvariance) * (wi / W)
		}
	}

	sigma := math.Sqrt(wvariance)
	return sigma
}

// Sum up all bins of the histogram
func (h *Histogram) Sum() (res float64) {
	for _, val := range h.bin {
		res += val
	}
	return
}
