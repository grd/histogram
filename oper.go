package histogram

/* histogram/oper.go
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
 * File gsl_histogram_oper.c: 
 * Routine to make operation on histograms. 
 * Need GSL library and header.
 * Contains the routines:
 * gsl_histogram_same.Binning check if two histograms have the same.Binning 
 * gsl_histogram_add          add two histograms
 * gsl_histogram_sub          subctract two histograms
 * gsl_histogram_mult         multiply two histograms
 * gsl_histogram_div          divide two histograms
 * gsl_histogram_scale        scale histogram contents
 *
 * Author: S. Piccardi
 * Jan. 2000
 *
 ***************************************************************/

import (
	"log"
)

// EqualBins control if two histograms have the same.Binning
func EqualBins(h1, h2 *Histogram) bool {
	if len(h1.Range) != len(h2.Range) {
		return false
	}

	for i := range h1.Range {
		if h1.Range[i] != h2.Range[i] {
			return false
		}
	}

	return true
}

// Add two histograms
func (h1 *Histogram) Add(h2 *Histogram) {
	if !EqualBins(h1, h2) {
		log.Fatal("histograms have different.Binning")
	}

	for i := range h1.Bin {
		h1.Bin[i] += h2.Bin[i]
	}
}

// Sub subtract two histograms
func (h1 *Histogram) Sub(h2 *Histogram) {
	if !EqualBins(h1, h2) {
		log.Fatal("histograms have different.Binning")
	}

	for i := range h1.Bin {
		h1.Bin[i] -= h2.Bin[i]
	}
}

// Mul multiply two histograms
func (h1 *Histogram) Mul(h2 *Histogram) {
	if !EqualBins(h1, h2) {
		log.Fatal("histograms have different.Binning")
	}

	for i := range h1.Bin {
		h1.Bin[i] *= h2.Bin[i]
	}
}

// Div divide two histograms
func (h1 *Histogram) Div(h2 *Histogram) {
	if !EqualBins(h1, h2) {
		log.Fatal("histograms have different.Binning")
	}

	for i := range h1.Bin {
		h1.Bin[i] /= h2.Bin[i]
	}
}

// Scale a histogram by a numeric factor 
func (h *Histogram) Scale(scale float64) {
	for i := range h.Bin {
		h.Bin[i] *= scale
	}
}

// Shift a histogram by a numeric offset 
func (h *Histogram) Shift(shift float64) {
	for i := range h.Bin {
		h.Bin[i] += shift
	}
}
