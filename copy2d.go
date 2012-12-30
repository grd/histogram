package histogram

/* copy2d.go
 *
 * Copyright (C) 2000  Simone Piccardi
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU General Public License as
 * published by the Free Software Foundation; either version 2 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * General Public License for more details.
 *
 * You should have received a copy of the GNU General Public
 * License along with this library; if not, write to the
 * Free Software Foundation, Inc., 59 Temple Place - Suite 330,
 * Boston, MA 02111-1307, USA.
 */

// Copy the contents of histogram src into dest
func (src *Histogram2d) Copy(dest *Histogram2d) error {
	if src.LenX() != dest.LenX() || src.LenY() != dest.LenY() {
		return ErrSize
	}

	copy(dest.xrange, src.xrange)
	copy(dest.yrange, src.yrange)
	copy(dest.bin, src.bin)

	return nil

}

// Clone an histogram creating an identical new one
func (src *Histogram2d) Clone() (clone *Histogram2d, err error) {
	clone, err = NewHistogram2dRange(src.xrange, src.yrange)

	if err != nil {
		return nil, ErrAlloc
	}

	copy(clone.bin, src.bin)

	return
}
