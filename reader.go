package protocolbytes

func BindBuffer(b []byte) *Buffer {
	return (*Buffer)(&b)
}

func (w *Buffer) RInt8() int8 {
	if w.checkIfHitOutOfBounds(1) {
		return 0
	}

	v := int8((*w)[0])
	*w = (*w)[1:]

	return v
}

func (w *Buffer) RInt16() int16 {
	if w.checkIfHitOutOfBounds(2) {
		return 0
	}

	v := int16((*w)[0])<<8 | int16((*w)[1])
	*w = (*w)[2:]

	return v
}

func (w *Buffer) RInt32() int32 {
	if w.checkIfHitOutOfBounds(4) {
		return 0
	}

	v := int32((*w)[0])<<24 | int32((*w)[1])<<16 | int32((*w)[2])<<8 | int32((*w)[3])
	*w = (*w)[4:]

	return v
}

func (w *Buffer) RInt64() int64 {
	if w.checkIfHitOutOfBounds(8) {
		return 0
	}
	var v int64
	for i := 0; i < 8; i++ {
		v |= int64((*w)[i]) << (56 - 8*i)
	}
	*w = (*w)[8:]

	return v
}

func (w *Buffer) RUInt8() uint8 {
	if w.checkIfHitOutOfBounds(1) {
		return 0
	}
	v := uint8((*w)[0])
	*w = (*w)[1:]

	return v
}

func (w *Buffer) RUInt16() uint16 {
	if w.checkIfHitOutOfBounds(2) {
		return 0
	}
	v := uint16((*w)[0])<<8 | uint16((*w)[1])
	*w = (*w)[2:]

	return v
}

func (w *Buffer) RUInt32() uint32 {
	if w.checkIfHitOutOfBounds(4) {
		return 0
	}
	v := uint32((*w)[0])<<24 | uint32((*w)[1])<<16 | uint32((*w)[2])<<8 | uint32((*w)[3])
	*w = (*w)[4:]

	return v
}

func (w *Buffer) RUInt64() uint64 {
	if w.checkIfHitOutOfBounds(8) {
		return 0
	}
	var v uint64
	for i := 0; i < 8; i++ {
		v |= uint64((*w)[i]) << (56 - 8*i)
	}
	*w = (*w)[8:]

	return v
}

func (w *Buffer) RUTF() string {
	if w.checkIfHitOutOfBounds(2) {
		return ""
	}
	l := w.RVarInt()
	s := string((*w)[:l])
	*w = (*w)[l:]

	return s
}

func (w *Buffer) RString() string {
	if w.checkIfHitOutOfBounds(2) {
		return ""
	}
	l := w.RVarInt()
	s := string((*w)[:l])
	*w = (*w)[l:]

	return s
}

func (w *Buffer) RBytes() []byte {
	if w.checkIfHitOutOfBounds(2) {
		return nil
	}

	l := w.RVarInt()
	b := make([]byte, l)
	copy(b, (*w)[:l])
	*w = (*w)[l:]

	return b
}

func (w *Buffer) RBool() bool {
	if w.checkIfHitOutOfBounds(1) {
		return false
	}
	return w.RUInt8() == 1
}

func (w *Buffer) RVarInt() int32 {
	if w.checkIfHitOutOfBounds(1) {
		return 0
	}
	var v int32
	var s uint
	for {
		b := w.RUInt8()
		v |= int32(b&0x7f) << s
		s += 7
		if b&0x80 == 0 {
			break
		}
	}

	return v
}

func (w *Buffer) checkIfHitOutOfBounds(n int) bool {
	return len(*w) < n
}
