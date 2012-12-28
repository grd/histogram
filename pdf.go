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

	i, err := find(p.Sum, r)

	if err != nil {
		return 0.0, err
	}

	delta := (r - p.Sum[i]) / (p.Sum[i+1] - p.Sum[i])
	x := p.Range[i] + delta*(p.Range[i+1]-p.Range[i])

	return x, nil
}

func NewPdf(n int) (*Pdf, error) {
	var p Pdf

	if n <= 0 {
		return nil, errors.New("histogram pdf length n must be positive integer")
	}

	p.Range = make([]float64, n+1)
	p.Sum = make([]float64, n+1)

	return &p, nil
}

func (p *Pdf) Init(h *Histogram) error {
	if len(p.Sum) != len(h.Range) {
		return errors.New("histogram length must match pdf length")
	}

	for i := range h.Bin {
		if h.Bin[i] < 0 {
			return errors.New("histogram.Bins must be non-negative to compute" +
				"a probability distribution")
		}
	}

	for i := range p.Range {
		p.Range[i] = h.Range[i]
	}

	var mean, Sum float64

	for i := range h.Bin {
		mean += (h.Bin[i] - mean) / float64(i+1)
	}

	p.Sum[0] = 0

	for i := range h.Bin {
		Sum += (h.Bin[i] / mean) / float64(len(h.Bin))
		p.Sum[i+1] = Sum
	}

	return nil
}
