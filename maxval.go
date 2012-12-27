package histogram

/* histogram/maxval.go
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
 * File gsl_histogram_maxval.c: 
 * Routine to find maximum and minumum content of a hisogram. 
 * Need GSL library and header.
 * Contains the routines:
 * gsl_histogram_max_val find max content values
 * gsl_histogram_min_val find min content values
 * gsl_histogram.Bin_max find coordinates of max contents.Bin
 * gsl_histogram.Bin_min find coordinates of min contents.Bin
 *
 * Author: S. Piccardi
 * Jan. 2000
 *
 ***************************************************************/

func (h *Histogram) MaxVal() float64 {
	max := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] > max {
			max = h.Bin[i]
		}
	}
	return max
}

func (h *Histogram) MaxBin() int {
	var imax int
	max := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] > max {
			max = h.Bin[i]
			imax = i
		}
	}
	return imax
}

func (h *Histogram) MinVal() float64 {
	min := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] < min {
			min = h.Bin[i]
		}
	}
	return min
}

func (h *Histogram) MinBin() int {
	var imin int
	min := h.Bin[0]
	for i := range h.Bin {
		if h.Bin[i] < min {
			min = h.Bin[i]
			imin = i
		}
	}
	return imin
}
