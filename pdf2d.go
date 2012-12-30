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

func NewPdf2d(nx, ny int) (p *Pdf2d, err error) {
	p = &Pdf2d{}
	n := nx * ny

	if n == 0 {
		err = errors.New("histogram2d pdf length n must be positive integer")
		return
	}

	p.xrange = make([]float64, nx+1)
	p.yrange = make([]float64, ny+1)
	p.sum = make([]float64, n+1)
	return
}

func (p *Pdf2d) Init(h *Histogram2d) error {
	nx, ny := len(p.xrange)-1, len(p.yrange)-1
	n := nx * ny

	if nx != h.LenX() || ny != h.LenY() {
		return errors.New("histogram2d size must match pdf size")
	}

	for i := 0; i < n; i++ {
		if h.bin[i] < 0 {
			return errors.New("histogram bins must be non-negative to compute " +
				"a probability distribution")
		}
	}

	copy(p.xrange, h.xrange)
	copy(p.yrange, h.yrange)

	var mean, sum float64

	for i := range h.bin {
		mean += (h.bin[i] - mean) / float64(i+1)
	}

	p.sum[0] = 0

	for i := range h.bin {
		sum += (h.bin[i] / mean) / float64(n)
		p.sum[i+1] = sum
	}

	return nil
}
