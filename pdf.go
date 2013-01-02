package histogram

/* pdf.go
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

func (p *Pdf) Sample(r float64) (float64, error) {

	// Wrap the exclusive top of the bin down to the inclusive bottom of the bin. 
	// Since this is a single point it should not affect the distribution.

	if r == 1.0 {
		r = 0.0
	}

	i, err := find(p.sum, r)

	if err != nil {
		return 0.0, err
	}

	delta := (r - p.sum[i]) / (p.sum[i+1] - p.sum[i])
	x := p.range_[i] + delta*(p.range_[i+1]-p.range_[i])

	return x, nil
}

func NewPdf(h *Histogram) (*Pdf, error) {
	for i := range h.bin {
		if h.bin[i] < 0 {
			return nil, errors.New("histogram bins must be non-negative " +
				"to compute a probability distribution")
		}
	}

	var p Pdf
	p.range_ = make([]float64, h.Len()+1)
	p.sum = make([]float64, h.Len()+1)

	for i := range p.range_ {
		p.range_[i] = h.range_[i]
	}

	var mean, sum float64

	for i := range h.bin {
		mean += (float64(h.bin[i]) - mean) / float64(i+1)
	}

	for i := range h.bin {
		sum += (float64(h.bin[i]) / mean) / float64(len(h.bin))
		p.sum[i+1] = sum
	}

	return &p, nil
}
