package histogram

/* histogram/file.go
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

func (h *Histogram) Read(data []byte) (n int, err error) {
	if len(data) < (len(h.Range)+len(h.Bin))*8 {
		return 0, io.ErrShortBuffer
	}

	for i := range h.Range {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.Range[i]))
	}

	for i := range h.Bin {
		n += binary.PutUvarint(data[n:], math.Float64bits(h.Bin[i]))
	}

	return n, io.EOF
}

func (h *Histogram) Write(data []byte) (n int, err error) {
	for i := range h.Range {
		val, num := binary.Uvarint(data[n:])
		if num <= 0 {
			return n, io.ErrShortWrite
		}
		h.Range[i] = math.Float64frombits(val)
		n += num
	}

	for i := range h.Bin {
		val, num := binary.Uvarint(data[n:])
		if num <= 0 {
			return n, io.ErrShortWrite
		}
		h.Bin[i] = math.Float64frombits(val)
		n += num
	}

	return n, io.EOF
}

func (h *Histogram) String() (res string) {
	for i := range h.Bin {
		str := fmt.Sprintf("%.19e %.19e %.19e\n", h.Range[i], h.Range[i+1], h.Bin[i])
		// str := fmt.Sprintf("%f %f %f\n", h.Range[i], h.Range[i+1], h.Bin[i])
		res += str
	}
	return
}

func (h *Histogram) Scan(s fmt.ScanState, ch rune) (err error) {
	var buf bytes.Buffer

	for i := range h.Bin {
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
		n, _ := fmt.Sscanf(str, "%f %f %f", &h.Range[i], &h.Range[i+1], &h.Bin[i])

		if n < 3 {
			return io.ErrUnexpectedEOF
		}
	}

	return
}
