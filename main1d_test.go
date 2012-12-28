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
	"io"
	"math"
	"os"
	"testing"
)

const N = 397
const NR = 10

func Test_1d(t *testing.T) {
	xr := []float64{0.0, 1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0}

	h, err := NewHistogramIncr(N)
	if err != nil {
		t.Error(err)
	}

	h1, err := NewHistogramIncr(N)
	if err != nil {
		t.Error(err)
	}

	g, err := NewHistogramIncr(N)
	if err != nil {
		t.Error(err)
	}

	gsl_test(h.range_ == nil, "NewHistogramIncr returns valid range pointer")
	gsl_test(h.bin == nil, "NewHistogramIncr returns valid bin pointer")

	hr, err := NewHistogramRange(xr)
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

	{
		err := hr.SetRanges(xr)
		if err != nil {
			t.Error(err)
		}

		for i := range hr.range_ {
			if hr.range_[i] != xr[i] {
				t.Error("Histogram.SetRange sets range")
			}
		}
	}

	for i := 0; i < N; i++ {
		err := h.Accumulate(float64(i), float64(i))
		if err != nil {
			t.Fatal(err)
		}
	}

	{
		for i := 0; i < N; i++ {
			if h.bin[i] != float64(i) {
				t.Fatal("Histogram.Accumulate writes into array")
			}
		}
	}

	{
		for i := 0; i < N; i++ {
			if h.Get(i) != float64(i) {
				t.Fatal("Histogram.Get reads from array")
			}
		}
	}

	for i := 0; i <= N; i++ {
		h1.range_[i] = 100.0 + float64(i)
	}

	err = h.Copy(h1)
	if err != nil {
		t.Fatal(err)
	}

	{
		for i := 0; i <= N; i++ {
			if h1.range_[i] != h.range_[i] {
				t.Fatal("Histogram.Copy copies bin ranges")
			}
		}
	}

	{
		for i := 0; i < N; i++ {
			if h1.Get(i) != h.Get(i) {
				t.Fatal("Histogram.Copy copies bin values")
			}
		}
	}

	// New memory allocation for h1
	h1, err = h.Clone()
	if err != nil {
		t.Fatal(err)
	}

	{
		for i := 0; i <= N; i++ {
			if h1.range_[i] != h.range_[i] {
				t.Fatal("Histogram.Clone copies bin ranges")
			}
		}
	}

	{
		for i := 0; i < N; i++ {
			if h1.Get(i) != h.Get(i) {
				t.Fatal("Histogram.Clone copies bin values")
			}
		}
	}

	h.Reset()

	{
		for i := 0; i < N; i++ {
			if h.bin[i] != 0 {
				t.Fatal("Histogram.Reset zeros array")
			}
		}
	}

	{
		for i := 0; i < N; i++ {
			h.Increment(float64(i))

			for j := 0; j <= i; j++ {
				if h.bin[j] != 1 {
					t.Fatal("Histogram.Increment increases bin value")
				}
			}

			for j := i + 1; j < N; j++ {
				if h.bin[j] != 0 {
					t.Fatal("Histogram.Increment increases bin value")
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
		if h.Max() != N {
			t.Fatal("Histogram.Max returns maximum")
		}
	}

	{
		if h.Min() != 0 {
			t.Fatal("Histogram.Min returns minimum")
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
		max := h.MaxVal()
		gsl_test(max != 123456.0, "Histogram.MaxVal finds maximum value")
	}

	{
		min := h.MinVal()
		gsl_test(min != -654321.0, "Histogram.MinVal finds minimum value")
	}

	{
		imax := h.MaxBin()
		gsl_test(imax != 2, "Histogram.MaxBin finds maximum value bin")
	}

	{
		imin := h.MinBin()
		gsl_test(imin != 4, "Histogram.MinBin find minimum value bin")
	}

	for i := 0; i < N; i++ {
		h.bin[i] = float64(i + 27)
		g.bin[i] = float64((i + 27) * (i + 1))
	}

	{
		Sum := h.Sum()
		gsl_test(Sum != N*27+((N-1)*N)/2, "Histogram.Sum sums all bin values")
	}

	g.Copy(h1)
	h1.Add(h)

	{
		var status bool
		for i := 0; i < N; i++ {
			if h1.bin[i] != g.bin[i]+h.bin[i] {
				status = true
			}
		}
		gsl_test(status, "Histogram.Add histogram addition")
	}

	g.Copy(h1)
	h1.Sub(h)

	{
		var status bool
		for i := 0; i < N; i++ {
			if h1.bin[i] != g.bin[i]-h.bin[i] {
				status = true
			}
		}
		gsl_test(status, "Histogram.Sub histogram subtraction")
	}

	g.Copy(h1)
	h1.Mul(h)

	{
		var status bool
		for i := 0; i < N; i++ {
			if h1.bin[i] != g.bin[i]*h.bin[i] {
				status = true
			}
		}
		gsl_test(status, "Histogram.Mul histogram multiplication")
	}

	g.Copy(h1)
	h1.Div(h)

	{
		var status bool
		for i := 0; i < N; i++ {
			if h1.bin[i] != g.bin[i]/h.bin[i] {
				status = true
			}
		}
		gsl_test(status, "Histogram.Div histogram division")
	}

	g.Copy(h1)
	h1.Scale(0.5)

	{
		var status bool
		for i := 0; i < N; i++ {
			if h1.bin[i] != 0.5*g.bin[i] {
				status = true
			}
		}
		gsl_test(status, "Histogram.Scale histogram scaling")
	}

	g.Copy(h1)
	h1.Shift(0.25)

	{
		var status bool
		for i := 0; i < N; i++ {
			if h1.bin[i] != 0.25+g.bin[i] {
				status = true
			}
		}
		gsl_test(status, "Histogram.Shift histogram shift")
	}

	//  Reallocate h

	h, err = NewHistogramUniform(N, 0.0, 1.0)

	gsl_test(h.range_ == nil,
		"NewHistogramUniform returns valid range pointer")
	gsl_test(h.bin == nil,
		"NewHistogramUniform returns valid bin pointer")

	h.Accumulate(0.0, 1.0)
	h.Accumulate(0.1, 2.0)
	h.Accumulate(0.2, 3.0)
	h.Accumulate(0.3, 4.0)

	{
		var expected float64
		var status bool
		i1, _ := h.Find(0.0)
		i2, _ := h.Find(0.1)
		i3, _ := h.Find(0.2)
		i4, _ := h.Find(0.3)

		for i := 0; i < N; i++ {
			if i == i1 {
				expected = 1.0
			} else if i == i2 {
				expected = 2.0
			} else if i == i3 {
				expected = 3.0
			} else if i == i4 {
				expected = 4.0
			} else {
				expected = 0.0
			}

			if h.bin[i] != expected {
				status = true
			}
		}
		gsl_test(status, "Histogram.Find returns index")
	}

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

		gsl_test(status, "Histogram.Scan and .String")

		f.Close()
	}

	{
		f, _ := os.Create("test.dat")
		_, err := io.Copy(f, h)
		if err != nil {
			t.Fatal(err)
		}
		f.Close()
	}

	{
		f, _ := os.Open("test.dat")
		hh, _ := NewHistogramIncr(N)
		var status bool

		io.Copy(hh, f)

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

		gsl_test(status, "Histogram.Read and .Write")

		f.Close()
	}
}

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
