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
	"fmt"
)

// FormatString2d is used by the String and Scan functions for data parsing.
// If you want a different output, just modify the variable.
var FormatString2d = "%v %v %v %v %v\n"

// String uses the variabele FormatString for the data parsing
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
