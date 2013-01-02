package histogram

/* main1d_test.go
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
	"fmt"
	"math"
	"testing"
)

const N = 397
const NR = 10

func Test_1d(t *testing.T) {
	xr := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}
	nr := make([]float64, N+1)
	for i := range nr {
		nr[i] = float64(i)
	}

	h, err := NewHistogram(nr)
	if err != nil {
		t.Error(err)
	}

	g, err := NewHistogram(nr)
	if err != nil {
		t.Error(err)
	}

	gsl_test(h.range_ == nil, "NewHistogramIncr returns valid range pointer")
	gsl_test(h.bin == nil, "NewHistogramIncr returns valid bin pointer")

	hr, err := NewHistogram(xr)
	if err != nil {
		t.Error(err)
	}

	gsl_test(hr.range_ == nil, "NewHistogramIncr returns valid range pointer")
	gsl_test(hr.bin == nil, "NewHistogramIncr returns valid bin pointer")

	{
		for i := 0; i <= NR; i++ {
			if hr.range_[i] != xr[i] {
				t.Error("NewHistogramRange creates range")

			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < i; j++ {
			err = h.Add(float64(i))
		}
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		for i := 0; i < N; i++ {
			if h.bin[i] != i {
				t.Fatal("Histogram.Add writes into array")
			}
		}
	}

	{
		for i := 0; i < N; i++ {
			if h.Get(i) != i {
				t.Fatal("Histogram.Get reads from array")
			}
		}
	}

	{
		// reset h.bin
		for i := range h.bin {
			h.bin[i] = 0
		}

		for i := 0; i < N; i++ {
			h.Add(float64(i))

			for j := 0; j <= i; j++ {
				if h.bin[j] != 1 {
					t.Fatal("Histogram.Add increases bin value")
				}
			}

			for j := i + 1; j < N; j++ {
				if h.bin[j] != 0 {
					t.Fatal("Histogram.Add increases bin value")
				}
			}
		}
	}

	{
		for i := 0; i < N; i++ {
			x0, x1 := h.GetRange(i)

			if x0 != float64(i) || x1 != float64(i+1) {
				t.Fatal("Histogram.GetRange returns range")
			}
		}
	}

	{
		if h.Len() != N {
			t.Fatal("Histogram.Len returns number of bins")
		}
	}

	h.bin[2] = 123456.0
	h.bin[4] = -654321

	{
		max, imax := h.Max()
		gsl_test(max != 123456.0, "Histogram.Max finds maximum value")
		gsl_test(imax != 2, "Histogram.Max finds maximum value bin")
	}

	{
		min, imin := h.Min()
		gsl_test(min != -654321.0, "Histogram.Min finds minimum value")
		gsl_test(imin != 4, "Histogram.Min find minimum value bin")
	}

	for i := 0; i < N; i++ {
		h.bin[i] = i + 27
		g.bin[i] = (i + 27) * (i + 1)
	}

	{
		sum := h.Sum()
		gsl_test(sum != N*27+((N-1)*N)/2, "Histogram.Sum sums all bin values")
	}

	//  Reallocate h

	uniRange := make([]float64, N+1)
	uniIncr := 1.0 / N
	for i := range uniRange {
		uniRange[i] = uniIncr * float64(i)
	}
	uniRange[N] = 1.0

	h, err = NewHistogram(uniRange)

	gsl_test(h.range_ == nil,
		"NewHistogramUniform returns valid range pointer")
	gsl_test(h.bin == nil,
		"NewHistogramUniform returns valid bin pointer")

	h.Add(0.0)
	for i := 0; i < 2; i++ {
		h.Add(0.1)
	}
	for i := 0; i < 3; i++ {
		h.Add(0.2)
	}
	for i := 0; i < 4; i++ {
		h.Add(0.3)
	}

	{
		var expected int
		var status bool
		i1, _ := h.Find(0.0)
		i2, _ := h.Find(0.1)
		i3, _ := h.Find(0.2)
		i4, _ := h.Find(0.3)

		for i := 0; i < N; i++ {
			if i == i1 {
				expected = 1
			} else if i == i2 {
				expected = 2
			} else if i == i3 {
				expected = 3
			} else if i == i4 {
				expected = 4
			} else {
				expected = 0
			}

			if h.bin[i] != expected {
				status = true
			}
		}
		gsl_test(status, "Histogram.Find returns index")
	}
	/*
		{
			f, _ := os.Create("test.txt")
			_, err = fmt.Fprint(f, h)
			f.Close()
		}

		{
			f, _ := os.Open("test.txt")
			hh, _ := NewHistogramIncr(N)
			var status bool

			fmt.Fscan(f, hh)

			for i := 0; i < N; i++ {
				if h.range_[i] != hh.range_[i] {
					status = true
				}
				if h.bin[i] != hh.bin[i] {
					status = true
				}
			}
			if h.range_[N] != hh.range_[N] {
				status = true
			}

			gsl_test(status, "Histogram.String")

			f.Close()
		}
	*/
}

func Test_1d_resample(t *testing.T) {
	hs := make([]float64, 11)
	for i := range hs {
		hs[i] = float64(i) * 0.1
	}

	h, err := NewHistogram(hs)
	if err != nil {
		t.Error(err)
	}

	h.Add(0.1)
	h.Add(0.2)
	h.Add(0.2)
	h.Add(0.31)

	fmt.Println(h)

	hhs := make([]float64, 101)
	for i := range hhs {
		hhs[i] = float64(i) * 0.01
	}

	hh, err := NewHistogram(hhs)
	if err != nil {
		t.Error(err)
	}

	p, err := NewPdf(h)
	if err != nil {
		t.Error(err)
	}

	for i := 0; i < 100000; i++ {
		u := urand()
		x, err := p.Sample(u)
		if err != nil {
			t.Fatal(err)
		}
		hh.Add(x)
	}

	for i := 0; i < 100; i++ {
		y := float64(hh.Get(i)) / 2500
		x, _ := hh.GetRange(i)
		k, err := h.Find(x)
		if err != nil {
			t.Error(err)
		}
		ya := float64(h.Get(k))

		if ya == 0 {
			if y != 0 {
				t.Errorf("%d: %g vs %g\n", i, y, ya)
			}
		} else {
			err := 1 / math.Sqrt(float64(hh.Get(i)))
			sigma := math.Abs((y - ya) / (ya * err))
			if sigma > 3 {
				t.Errorf("%g vs %g err=%g sigma=%g\n", y, ya, err, sigma)
			}
		}
	}
}
