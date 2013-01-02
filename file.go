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
	"fmt"
)

// FormatString is used by the String and Scan functions for data parsing.
// If you want a different output, just modify the variable.
var FormatString = "%v %v %v\n"

// String uses the variabele FormatString for the data parsing
func (h *Histogram) String() (res string) {
	for i := range h.bin {
		str := fmt.Sprintf(FormatString, h.range_[i], h.range_[i+1], h.bin[i])
		res += str
	}
	return
}

// String uses the variabele FormatString for the data parsing
func (h *HistogramInt) String() (res string) {
	for i := range h.bin {
		str := fmt.Sprintf(FormatString, h.range_[i], h.range_[i+1], h.bin[i])
		res += str
	}
	return
}
