package histogram

/* pdf2d.go
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

func (p *Pdf2d) Sample(r1, r2 float64) (x, y float64, err error) {

	// Wrap the exclusive top of the bin down to the inclusive bottom of the bin. 
	// Since this is a single point it should not affect the distribution.

	if r2 == 1.0 {
		r2 = 0.0
	}
	if r1 == 1.0 {
		r1 = 0.0
	}

	var k int
	if k, err = find(p.sum, r1); err != nil {
		err = errors.New("cannot find r1 in cumulative pdf")
		return
	}

	ny := len(p.yrange) - 1
	i := k / ny
	j := k - (i * ny)
	delta := (r1 - p.sum[k]) / (p.sum[k+1] - p.sum[k])
	x = p.xrange[i] + delta*(p.xrange[i+1]-p.xrange[i])
	y = p.yrange[j] + r2*(p.yrange[j+1]-p.yrange[j])
	return
}

var errNegPdf = "histogram bins must be non-negative to compute " +
	"a probability distribution"

func NewPdf2d(h *Histogram2d) (p *Pdf2d, err error) {
	for i := range h.bin {
		if h.bin[i] < 0 {
			err = errors.New(errNegPdf)
			return
		}
	}

	p = &Pdf2d{}
	p.xrange = make([]float64, len(h.xrange))
	p.yrange = make([]float64, len(h.yrange))
	p.sum = make([]float64, len(h.bin)+1)

	copy(p.xrange, h.xrange)
	copy(p.yrange, h.yrange)

	var mean, sum float64

	for i := range h.bin {
		mean += (float64(h.bin[i]) - mean) / float64(i+1)
	}

	for i := range h.bin {
		sum += (float64(h.bin[i]) / mean) / float64(len(h.bin))
		p.sum[i+1] = sum
	}

	return
}
