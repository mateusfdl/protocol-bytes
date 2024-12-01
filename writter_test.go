package protocolbytes_test

import (
	"testing"

	protocolbytes "github.com/mateusfdl/protocol-bytes"
)

func TestWriterInt8(t *testing.T) {
	w := protocolbytes.Buffer{}

	v := int8(0x7F)

	w.WInt8(v)

	if w[0] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w[0])
	}
}

func TestWriterInt16(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 2))

	v := int16(0x7FFF)

	w.WInt16(v)

	if w[0] != byte(v>>8) || w[1] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterInt32(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 4))

	v := int32(0x7FFFFFFF)

	w.WInt32(v)

	if w[0] != byte(v>>24) || w[1] != byte(v>>16) || w[2] != byte(v>>8) || w[3] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterInt64(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 8))

	v := int64(0x7FFFFFFFFFFFFFFF)

	w.WInt64(v)

	for i := 0; i < 8; i++ {
		if w[i] != byte(v>>(56-8*i)) {
			t.Errorf("Expected %v, got %v", v, w)
		}
	}
}

func TestWriterUInt8(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 1))

	v := uint8(0xFF)

	w.WUInt8(v)

	if w[0] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w[0])
	}
}

func TestWriterUInt16(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 2))

	v := uint16(0xFFFF)

	w.WUInt16(v)

	if w[0] != byte(v>>8) || w[1] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterUInt32(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 4))

	v := uint32(0xFFFFFFFF)

	w.WUInt32(v)

	if w[0] != byte(v>>24) || w[1] != byte(v>>16) || w[2] != byte(v>>8) || w[3] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterUInt64(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 8))

	v := uint64(0xFFFFFFFFFFFFFFFF)

	w.WUInt64(v)

	for i := 0; i < 8; i++ {
		if w[i] != byte(v>>(56-8*i)) {
			t.Errorf("Expected %v, got %v", v, w)
		}
	}
}

func TestWriterUTF(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 2))
	v := "こんにちは、世界"

	w.WUTF(v)

	if w[0] != byte(len(v)) || string(w[1:]) != v {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterString(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 2))
	v := "Hello, World!"

	w.WString(v)

	if w[0] != byte(len(v)) || string(w[1:]) != v {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterBytes(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 2))

	v := []byte("Hello, World!")

	w.WBytes(v)

	if w[0] != byte(len(v)) || string(w[1:]) != string(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterBool(t *testing.T) {
	w := protocolbytes.Buffer(make([]byte, 0, 1))

	v := true

	w.WBool(v)

	if w[0] != 1 {
		t.Errorf("Expected %v, got %v", v, w)
	}

	v = false

	w.WBool(v)

	if w[1] != 0 {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterVarInt(t *testing.T) {

	w := protocolbytes.Buffer(make([]byte, 0, 5))

	v := int32(0x7F)

	w.WVarInt(v)

	if w[0] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}

	w = protocolbytes.Buffer(make([]byte, 0, 5))

	v = int32(0x3FFF)

	w.WVarInt(v)

	if w[0] != byte(v>>7|0x80) || w[1] != byte(v&0x7F) {
		t.Errorf("Expected %v, got %v", v, w)
	}

	w = protocolbytes.Buffer(make([]byte, 0, 5))

	v = int32(0x1FFFFF)

	w.WVarInt(v)

	if w[0] != byte(v>>14|0x80) || w[1] != byte(v>>7|0x80) || w[2] != byte(v&0x7F) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func TestWriterVarLong(t *testing.T) {

	w := protocolbytes.Buffer(make([]byte, 0, 10))

	v := int64(0x7F)

	w.WVarLong(v)

	if w[0] != byte(v) {
		t.Errorf("Expected %v, got %v", v, w)
	}

	w = protocolbytes.Buffer(make([]byte, 0, 10))

	v = int64(0x3FFF)

	w.WVarLong(v)

	if w[0] != byte(v>>7|0x80) || w[1] != byte(v&0x7F) {
		t.Errorf("Expected %v, got %v", v, w)
	}

	w = protocolbytes.Buffer(make([]byte, 0, 10))

	v = int64(0x1FFFFF)

	w.WVarLong(v)

	if w[0] != byte(v>>14|0x80) || w[1] != byte(v>>7|0x80) || w[2] != byte(v&0x7F) {
		t.Errorf("Expected %v, got %v", v, w)
	}
}

func BenchmarkWriterInt8(b *testing.B) {
	w := protocolbytes.Buffer{}

	v := uint8(0x7F)

	for i := 0; i < b.N; i++ {
		w.WUInt8(v)
	}

	b.ReportAllocs()
}
func BenchmarkWriteStandardInt8(b *testing.B) {
	w := make([]byte, 1)

	v := uint8(0x7F)

	for i := 0; i < b.N; i++ {
		w = append(w, byte(v))
	}

	b.ReportAllocs()
}
