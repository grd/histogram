package histogram

/* stat2d.go
 *
 * Copyright (C) 2002  Achim Gaedke
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
	"math"
)

// Sum up all bins of histogram2d
func (h *Histogram2d) Sum() (sum int) {
	for _, val := range h.bin {
		sum += val
	}
	return
}

// Xmean returns the histogram arithmetic mean
func (h *Histogram2d) Xmean() (wmean float64) {
	nx, ny := h.LenX(), h.LenY()
	var W float64

	// Compute the mean 

	for i := 0; i < nx; i++ {
		xi := (h.xrange[i+1] + h.xrange[i]) / 2.0
		var wi float64

		for j := 0; j < ny; j++ {
			wij := float64(h.bin[i*ny+j])
			if wij > 0 {
				wi += wij
			}
		}
		if wi > 0 {
			W += wi
			wmean += (xi - wmean) * (wi / W)
		}
	}

	return
}

// Ymean returns the histogram arithmetic mean
func (h *Histogram2d) Ymean() (wmean float64) {
	nx, ny := h.LenX(), h.LenY()

	// Compute the mean 

	var W float64

	for j := 0; j < ny; j++ {
		yj := (h.yrange[j+1] + h.yrange[j]) / 2.0
		var wj float64

		for i := 0; i < nx; i++ {
			wij := float64(h.bin[i*ny+j])
			if wij > 0 {
				wj += wij
			}
		}

		if wj > 0 {
			W += wj
			wmean += (yj - wmean) * (wj / W)
		}
	}

	return
}

func (h *Histogram2d) Xsigma() (xsigma float64) {
	xmean := h.Xmean()
	nx, ny := h.LenX(), h.LenY()

	// Compute the variance 

	var wvariance, W float64

	for i := 0; i < nx; i++ {
		xi := (h.xrange[i+1]+h.xrange[i])/2 - xmean
		var wi float64

		for j := 0; j < ny; j++ {
			wij := float64(h.bin[i*ny+j])
			if wij > 0 {
				wi += wij
			}
		}

		if wi > 0 {
			W += wi
			wvariance += ((xi * xi) - wvariance) * (wi / W)
		}
	}

	xsigma = math.Sqrt(wvariance)
	return
}

func (h *Histogram2d) Ysigma() (ysigma float64) {
	ymean := h.Ymean()
	nx, ny := h.LenX(), h.LenY()

	// Compute the variance

	var wvariance, W float64

	for j := 0; j < ny; j++ {
		yj := (h.yrange[j+1]+h.yrange[j])/2.0 - ymean
		var wj float64

		for i := 0; i < nx; i++ {
			wij := float64(h.bin[i*ny+j])
			if wij > 0 {
				wj += wij
			}
		}
		if wj > 0 {
			W += wj
			wvariance += ((yj * yj) - wvariance) * (wj / W)
		}
	}

	ysigma = math.Sqrt(wvariance)
	return
}

func (h *Histogram2d) Covariance() float64 {
	xmean := h.Xmean()
	ymean := h.Ymean()
	nx, ny := h.LenX(), h.LenY()

	// Compute the covariance

	var wcovariance, W float64

	for j := 0; j < ny; j++ {
		for i := 0; i < nx; i++ {
			xi := (h.xrange[i+1]+h.xrange[i])/2.0 - xmean
			yj := (h.yrange[j+1]+h.yrange[j])/2.0 - ymean
			wij := float64(h.bin[i*ny+j])

			if wij > 0 {
				W += wij
				wcovariance += ((xi * yj) - wcovariance) * (wij / W)
			}
		}
	}

	return wcovariance
}
