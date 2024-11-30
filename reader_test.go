package protocolbytes_test

import (
	"testing"

	protocolbytes "github.com/mateusfdl/protocol-bytes"
)

func TestReaderInt8(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x7F})

	if len(r) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(r))
	}

	v := r.RInt8()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != int8(0x7F) {
		t.Errorf("Expected %v, got %v", int8(0x7F), v)
	}
}

func TestReaderInt16(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x7F, 0xFF})

	if len(r) != 2 {
		t.Errorf("Expected %v, got %v", 2, len(r))
	}

	v := r.RInt16()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != int16(0x7FFF) {
		t.Errorf("Expected %v, got %v", int16(0x7FFF), v)
	}
}

func TestReaderInt32(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x7F, 0xFF, 0xFF, 0xFF})

	if len(r) != 4 {
		t.Errorf("Expected %v, got %v", 4, len(r))
	}

	v := r.RInt32()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != int32(0x7FFFFFFF) {
		t.Errorf("Expected %v, got %v", int32(0x7FFFFFFF), v)
	}
}

func TestReaderInt64(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x7F, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	if len(r) != 8 {
		t.Errorf("Expected %v, got %v", 8, len(r))
	}

	v := r.RInt64()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != int64(0x7FFFFFFFFFFFFFFF) {
		t.Errorf("Expected %v, got %v", int64(0x7FFFFFFFFFFFFFFF), v)
	}
}

func TestReaderUInt8(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0xFF})

	if len(r) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(r))
	}

	v := r.RUInt8()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != uint8(0xFF) {
		t.Errorf("Expected %v, got %v", uint8(0xFF), v)
	}
}

func TestReaderUInt16(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0xFF, 0xFF})

	if len(r) != 2 {
		t.Errorf("Expected %v, got %v", 2, len(r))
	}

	v := r.RUInt16()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != uint16(0xFFFF) {
		t.Errorf("Expected %v, got %v", uint16(0xFFFF), v)
	}
}

func TestReaderUInt32(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0xFF, 0xFF, 0xFF, 0xFF})

	if len(r) != 4 {
		t.Errorf("Expected %v, got %v", 4, len(r))
	}

	v := r.RUInt32()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != uint32(0xFFFFFFFF) {
		t.Errorf("Expected %v, got %v", uint32(0xFFFFFFFF), v)
	}
}

func TestReaderUInt64(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0xFF})

	if len(r) != 8 {
		t.Errorf("Expected %v, got %v", 8, len(r))
	}

	v := r.RUInt64()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != uint64(0xFFFFFFFFFFFFFFFF) {
		t.Errorf("Expected %v, got %v", uint64(0xFFFFFFFFFFFFFFFF), v)
	}
}

func TestReaderUTF(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x02, 0x41, 0x42})

	if len(r) != 3 {
		t.Errorf("Expected %v, got %v", 3, len(r))
	}

	v := r.RUTF()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != "AB" {
		t.Errorf("Expected %v, got %v", "AB", v)
	}
}

func TestReaderString(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x02, 0x41, 0x42})

	if len(r) != 3 {
		t.Errorf("Expected %v, got %v", 3, len(r))
	}

	v := r.RString()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != "AB" {
		t.Errorf("Expected %v, got %v", "AB", v)
	}
}

func TestReaderBytes(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x02, 0x41, 0x42})

	if len(r) != 3 {
		t.Errorf("Expected %v, got %v", 3, len(r))
	}

	v := r.RBytes()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if string(v) != "AB" {
		t.Errorf("Expected %v, got %v", "AB", v)
	}
}

func TestReaderBool(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x01})

	if len(r) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(r))
	}

	v := r.RBool()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if !v {
		t.Errorf("Expected %v, got %v", true, v)
	}
}

func TestReaderVarInt(t *testing.T) {
	r := protocolbytes.Buffer([]byte{0x01})

	if len(r) != 1 {
		t.Errorf("Expected %v, got %v", 1, len(r))
	}

	v := r.RVarInt()

	if len(r) != 0 {
		t.Errorf("Expected %v, got %v", 0, len(r))
	}
	if v != 1 {
		t.Errorf("Expected %v, got %v", 1, v)
	}
}

func BenchmarkReaderInt8(b *testing.B) {
	r := protocolbytes.Buffer([]byte{})

	for i := 0; i < b.N; i++ {
		r.RInt8()
	}

	b.ReportAllocs()
}

func BenckmarkReaderInt8Standard(b *testing.B) {

	m := make([]byte, 1)

	for i := 0; i < b.N; i++ {
		m[i] = 0x7F
		_ = int8(m[0])
	}

	b.ReportAllocs()
}
