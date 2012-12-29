package histogram

/* file.go
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
func (h *Histogram) Read(data []byte) (n int, err error) {
	if len(data) < (len(h.range_)+len(h.bin))*8 {
		return 0, io.ErrShortBuffer
	}

	for i := range h.range_ {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.range_[i]))
	}

	for i := range h.bin {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.bin[i]))
	}

	return n, io.EOF
}

// Write function corresponds with io.Writer.
// The data is binary.
func (h *Histogram) Write(data []byte) (n int, err error) {
	for i := range h.range_ {
		val, num := binary.Uvarint(data[n:])
		if num <= 0 {
			return n, io.ErrShortWrite
		}
		h.range_[i] = math.Float64frombits(val)
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

// FormatString is used by the String and Scan functions for data parsing.
// If you want a different output, just modify the variable.
var FormatString = "%f %f %f\n"

// String function corresponds with fmt.Stringer.
// String uses the variabele FormatString for the data parsing
// (which is plain text).
func (h *Histogram) String() (res string) {
	for i := range h.bin {
		str := fmt.Sprintf(FormatString, h.range_[i], h.range_[i+1], h.bin[i])
		res += str
	}
	return
}

// Scan function corresponds with fmt.Scanner. 
// Scan uses the variabele FormatString for the data parsing
// (which is plain text).
func (h *Histogram) Scan(s fmt.ScanState, ch rune) (err error) {
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
		n, _ := fmt.Sscanf(str, FormatString, &h.range_[i], &h.range_[i+1], &h.bin[i])

		if n < 3 {
			return io.ErrUnexpectedEOF
		}
	}

	return
}
