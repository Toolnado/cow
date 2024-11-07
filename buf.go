package cow

import (
	"unsafe"
)

type COWBuffer struct {
	data []byte
	refs *int
}

func NewCOWBuffer(data []byte) COWBuffer {
	return COWBuffer{
		data: data,
		refs: new(int),
	}
}

func (b *COWBuffer) Clone() COWBuffer {
	*b.refs++
	return *b
}

func (b *COWBuffer) Close() {
	*b.refs--
}

func (b *COWBuffer) Update(index int, value byte) bool {
	if index < 0 || len(b.data) <= index {
		return false
	}

	if *b.refs != 0 {
		*b.refs--
		buf := make([]byte, len(b.data))
		copy(buf, b.data)
		b.data = buf
		b.refs = new(int)
	}

	b.data[index] = value
	return true
}

func (b *COWBuffer) String() string {
	if len(b.data) == 0 {
		return ""
	}
	str := unsafe.String(unsafe.SliceData(b.data), len(b.data))
	return str
}
