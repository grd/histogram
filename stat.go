package histogram

/* histogram/stat.go
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
 * File gsl_histogram_stat.c: 
 * Routines for statisticalcomputations on histograms. 
 * Need GSL library and header.
 * Contains the routines:
 * gsl_histogram_mean    compute histogram mean
 * gsl_histogram_sigma   compute histogram sigma
 *
 * Author: S. Piccardi
 * Jan. 2000
 *
 ***************************************************************/

// FIXME: We skip negative values in the histogram h.Bin[i] < 0,
// since those correspond to negative weights (BJG)

import (
	"math"
)

//
// Mean compute the.Bin-weighted arithmetic mean M of a histogram using the
// recurrence relation
//
// M(n) = M(n-1) + (x[n] - M(n-1)) (w(n)/(W(n-1) + w(n))) 
// W(n) = W(n-1) + w(n)
//
func (h *Histogram) Mean() float64 {
	n := h.Len()

	// wmean and W should be "long double" instead of float64 (so Float128 ?)
	var wmean, W float64

	for i := 0; i < n; i++ {
		xi := (h.Range[i+1] + h.Range[i]) / 2
		wi := h.Bin[i]

		if wi > 0 {
			W += wi
			wmean += (xi - wmean) * (wi / W)
		}
	}

	return wmean
}

func (h *Histogram) Sigma() float64 {
	n := h.Len()

	// long double wvariance, wmean and W ... Float128 ?
	var wvariance, wmean, W float64

	// FIXME: should use a single pass formula here, as given in
	// N.J.Higham 'Accuracy and Stability of Numerical Methods', p.12 

	// Compute the mean

	for i := 0; i < n; i++ {
		xi := (h.Range[i+1] + h.Range[i]) / 2
		wi := h.Bin[i]

		if wi > 0 {
			W += wi
			wmean += (xi - wmean) * (wi / W)
		}
	}

	// Compute the variance

	W = 0.0

	for i := 0; i < n; i++ {
		xi := ((h.Range[i+1]) + (h.Range[i])) / 2
		wi := h.Bin[i]

		if wi > 0 {
			// long double delta... Float128 ?
			delta := (xi - wmean)
			W += wi
			wvariance += (delta*delta - wvariance) * (wi / W)
		}
	}

	sigma := math.Sqrt(wvariance)
	return sigma
}

// Sum up all.Bins of histogram
func (h *Histogram) Sum() (res float64) {
	for i := range h.Bin {
		res += h.Bin[i]
	}
	return
}
