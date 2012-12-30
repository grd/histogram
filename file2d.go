package histogram

/* file2d.go
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
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

// Read function corresponds with io.Reader.
// The data is binary.
func (h *Histogram2d) Read(data []byte) (n int, err error) {

	if len(data) < (len(h.xrange)+len(h.yrange)+len(h.bin))*8 {
		return 0, io.ErrShortBuffer
	}

	for i := range h.xrange {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.xrange[i]))
	}

	for i := range h.yrange {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.yrange[i]))
	}

	for i := range h.bin {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.bin[i]))
	}

	return n, io.EOF
}

// Write function corresponds with io.Writer.
// The data is binary.
func (h *Histogram2d) Write(data []byte) (n int, err error) {
	for i := range h.xrange {
		val, num := binary.Uvarint(data[n:])
		if num <= 0 {
			return n, io.ErrShortWrite
		}
		h.xrange[i] = math.Float64frombits(val)
		n += num
	}

	for i := range h.yrange {
		val, num := binary.Uvarint(data[n:])
		if num <= 0 {
			return n, io.ErrShortWrite
		}
		h.yrange[i] = math.Float64frombits(val)
		n += num
	}

	for i := range h.bin {
		val, num := binary.Uvarint(data[n:])
		if num <= 0 {
			return n, io.ErrShortWrite
		}
		h.bin[i] = math.Float64frombits(val)
		n += num
	}

	return n, io.EOF
}

// FormatString2d is used by the String and Scan functions for data parsing.
// If you want a different output, just modify the variable.
var FormatString2d = "%f %f %f %f %f\n"

// String function corresponds with fmt.Stringer.
// String uses the variabele FormatString for the data parsing
// (which is plain text).
func (h *Histogram2d) String() (res string) {
	for i := 0; i < h.LenX(); i++ {
		for j := 0; j < h.LenY(); j++ {
			str := fmt.Sprintf(FormatString2d, h.xrange[i], h.xrange[i+1],
				h.yrange[j], h.yrange[j+1], h.bin[i*h.LenY()+j])
			res += str
		}
	}
	return
}

// Scan function corresponds with fmt.Scanner. 
// Scan uses the variabele FormatString for the data parsing
// (which is plain text).
func (h *Histogram2d) Scan(s fmt.ScanState, ch rune) (err error) {
	var buf bytes.Buffer

	for i := range h.bin {
		var done bool

		for !done {
			ch, _, err := s.ReadRune()
			if err != nil {
				return io.ErrUnexpectedEOF
			}
			if ch == '\n' {
				done = true
			}
			buf.WriteRune(ch)
		}
		str, _ := buf.ReadString('\n')
		n, _ := fmt.Sscanf(str, FormatString, &h.xrange[i], &h.xrange[i+1],
			&h.yrange[i], &h.yrange[i+1], &h.bin[i])

		if n < 5 {
			return io.ErrUnexpectedEOF
		}
	}

	return
}
