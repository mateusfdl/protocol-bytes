package protocolbytes

import "unicode/utf8"

type Buffer []byte

// Read a int8 from the buffer
func (w *Buffer) WInt8(i int8) {
	p := w.growSlice(1)

	(*w)[p] = byte(i)
}

// Write a int16 as a byte array
func (w *Buffer) WInt16(i int16) {
	p := w.growSlice(2)

	(*w)[p] = byte(i >> 8)
	(*w)[p+1] = byte(i)
}

// Write a int32 as a byte array
func (w *Buffer) WInt32(i int32) {
	p := w.growSlice(4)

	(*w)[p] = byte(i >> 24)
	(*w)[p+1] = byte(i >> 16)
	(*w)[p+2] = byte(i >> 8)
	(*w)[p+3] = byte(i)
}

// Write a int64 as a byte array
func (w *Buffer) WInt64(i int64) {
	p := w.growSlice(8)

	for n := 0; n < 8; n++ {
		(*w)[p+n] = byte(i >> (56 - 8*n))
	}
}

// Write a uint8 as a byte array
func (w *Buffer) WUInt8(i uint8) {
	p := w.growSlice(1)

	(*w)[p] = byte(i)
}

// Write a uint16 as a byte array
func (w *Buffer) WUInt16(i uint16) {
	p := w.growSlice(2)

	(*w)[p] = byte(i >> 8)
	(*w)[p+1] = byte(i)
}

// Write a uint32 as a byte array
func (w *Buffer) WUInt32(i uint32) {

	p := w.growSlice(4)

	(*w)[p] = byte(i >> 24)
	(*w)[p+1] = byte(i >> 16)
	(*w)[p+2] = byte(i >> 8)
	(*w)[p+3] = byte(i)
}

// Write a uint64 as a byte array
func (w *Buffer) WUInt64(i uint64) {
	p := w.growSlice(8)

	for n := 0; n < 8; n++ {
		(*w)[p+n] = byte(i >> (56 - 8*n))
	}
}

// Write a string as a UTF-8 encoded string prefixed with its length
// the length is the beginning byte of the string
// If the string is not UTF use WString
func (w *Buffer) WUTF(s string) {
	// Each rune can take up to utf8.UTFMax bytes.
	requiredCapacity := len(s)*utf8.UTFMax + 1

	p := w.growSlice(requiredCapacity)
	(*w)[p] = byte(len(s))

  writePos := p + 1
	for _, r := range s {
		var buf [utf8.UTFMax]byte
		n := utf8.EncodeRune(buf[:], r)

		// Grow the buffer further if needed
		if len(*w) < writePos+n {
			w.grow(n)
		}

		copy((*w)[writePos:], buf[:n])
		writePos += n
	}

	// Update the slice length to match the actual written content
	*w = (*w)[:writePos]}

// Write a string as a byte array prefixed with its length
// If the string is UTF-8 use WUTF
func (w *Buffer) WString(s string) {
  p := w.growSlice(len(s) + 1)

  (*w)[p] = byte(len(s))

  for i, r := range s {
    (*w)[p+i+1] = byte(r)
  }
}

// Write a byte array prefixed with its length
func (w *Buffer) WBytes(b []byte) {
  p := w.growSlice(len(b) + 1)

  (*w)[p] = byte(len(b))

  for i, v := range b {
    (*w)[p+i+1] = v
  }
}

// Write a boolean as a byte
func (w *Buffer) WBool(b bool) {
  p := w.growSlice(1)

  if b {
    (*w)[p] = 1
  } else {
    (*w)[p] = 0
  }
}

func (w *Buffer) WVarInt(i int32) {
  p := w.growSlice(5)

  for {
    if (i & ^0x7f) == 0 {
      (*w)[p] = byte(i)
      return
    }

    (*w)[p] = byte(i&0x7f | 0x80)
    i >>= 7 
    p++
  }
}

func (w *Buffer) WVarLong(i int64) {
  p := w.growSlice(10)

  for {
    if (i & ^0x7f) == 0 {
      (*w)[p] = byte(i)
      return
    }

    (*w)[p] = byte(i&0x7f | 0x80)

    i >>= 7 
    p++
  }
}

// Check if the slice is reslicebale, if so return the missing capacity 
// and true, otherwise return 0 and false
func (w *Buffer) checkForReeslice(n int) (int, bool) {
	c := cap((*w))
	l := len((*w))

	if c < l+n {
		return l + n - c, true
	}

	return 0, false
}


// Grow expands the buffer by n bytes and returns the updated buffer.
// If the buffer is nil, it initializes it with a capacity of 64 bytes.
// If more capacity is needed, it doubles the capacity or ensures it can hold len + n bytes.
// Finally, it reslices the buffer to the new length.
func (w *Buffer) grow(n int) Buffer {
	if *w == nil {
		// Preallocate 64 bytes
		*w = make([]byte, 0, 64)
	}

	l := len(*w)

	if missingCap, ok := w.checkForReeslice(n); ok {
		newCap := cap(*w) * 2
		if newCap < l+n {
			newCap = l + missingCap + n
		}

		ns := make([]byte, l, newCap)
		copy(ns, *w)
		*w = ns
	}

	*w = (*w)[:l+n]
	return *w
}

// Grow the slice by n bytes and return the previous length
// as index to start writing
func (w *Buffer) growSlice(n int) int {
	l := len(*w)
	*w = w.grow(n)
	return l
}
