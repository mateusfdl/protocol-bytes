package protocolbytes

type Buffer []byte

func (w *Buffer) WInt8(i int8) {
	*w = append((*w), byte(i))
}

func (w *Buffer) WInt16(i int16) {
	*w = append((*w), byte(i>>8), byte(i))
}

func (w *Buffer) WInt32(i int32) {
	*w = append((*w), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func (w *Buffer) WInt64(i int64) {
	for n := 0; n < 8; n++ {
		*w = append((*w), byte(i>>(56-8*n)))
	}
}

func (w *Buffer) WUInt8(i uint8) {
	*w = append((*w), i)
}

func (w *Buffer) WUInt16(i uint16) {
	*w = append((*w), byte(i>>8), byte(i))
}

func (w *Buffer) WUInt32(i uint32) {
	*w = append((*w), byte(i>>24), byte(i>>16), byte(i>>8), byte(i))
}

func (w *Buffer) WUInt64(i uint64) {
	for n := 0; n < 8; n++ {
		*w = append((*w), byte(i>>(56-8*n)))
	}
}

func (w *Buffer) WUTF(s string) {
	w.WVarInt(int32(len(s)))
	*w = append((*w), []byte(s)...)
}

func (w *Buffer) WString(s string) {
	w.WVarInt(int32(len(s)))
	*w = append((*w), []byte(s)...)
}

func (w *Buffer) WBytes(b []byte) {
	w.WVarInt(int32(len(b)))
	*w = append((*w), b...)
}

func (w *Buffer) WBool(b bool) {
	if b {
		w.WUInt8(1)
	} else {
		w.WUInt8(0)
	}
}

func (w *Buffer) WVarInt(i int32) {
	for {
		if (i & ^0x7f) == 0 {
			w.WUInt8(uint8(i))
			return
		}
		w.WUInt8(uint8(i&0x7f | 0x80))
		i >>= 7
	}
}

func (w *Buffer) WVarLong(i int64) {
	for {
		if (i & ^0x7f) == 0 {
			w.WUInt8(uint8(i))
			return
		}
		w.WUInt8(uint8(i&0x7f | 0x80))
		i >>= 7
	}
}
