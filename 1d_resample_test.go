package histogram

/* histogram/1d_resample_test.go
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
//	"fmt"
	"math"
	"testing"
)

func Test_1d_resample(t *testing.T) {
	h, err := NewHistogramUniform(10, 0.0, 1.0)
	if err != nil {
		t.Error(err)
	}

//	fmt.Println(h)

	h.Increment(0.1)
	h.Increment(0.2)
	h.Increment(0.2)
	h.Increment(0.3)

//	fmt.Println(h)

	p, err := NewPdf(10)
	if err != nil {
		t.Error(err)
	}

	hh, err := NewHistogramUniform(100, 0.0, 1.0)
	if err != nil {
		t.Error(err)
	}

	err = p.Init(h)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 100000; i++ {
		u := urand()
		x, err := p.Sample(u)
		if err != nil {
			t.Fatal(err)
		}
		hh.Increment(x)
	}

	for i := 0; i < 100; i++ {
		y := hh.Get(i) / 2500
		x, _ := hh.GetRange(i)
		k, err := h.Find(x)
		if err != nil {
			t.Error(err)
		}
		ya := h.Get(k)

		if ya == 0 {
			if y != 0 {
				t.Errorf("%d: %g vs %g\n", i, y, ya)
			}
		} else {
			err := 1 / math.Sqrt(hh.Get(i))
			sigma := math.Abs((y - ya) / (ya * err))
			if sigma > 3 {
				t.Errorf("%g vs %g err=%g sigma=%g\n", y, ya, err, sigma)
			}
		}
	}
}
