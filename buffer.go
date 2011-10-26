package pq

type buffer struct {
	b []byte
	pos int
}

func newBuffer() *buffer {
	// return a buffer that starts after the header
	return &buffer{
		b: make([]byte, 1024),
		pos:5,
	}
}

func (b *buffer) setType(c byte) {
	b.b[0] = c
}

func (b *buffer) setLength() {
	length := b.pos
	b.pos = 1
	b.writeInt32(int32(length-1)) // don't include Type in length
	b.pos = length
}

func (b *buffer) writeByte(c byte) {
	b.b[b.pos] = c
	b.pos += 1
}

func (b *buffer) writeString(v string) {
	for _, c := range v {
		b.writeByte(byte(c))
	}
	b.writeByte(0)
}

func (b *buffer) writeInt16(v int16) {
	b.writeByte(byte(v >> 8))
	b.writeByte(byte(v))
}

func (b *buffer) writeInt32(v int32) {
	b.writeByte(byte(v >> 24))
	b.writeByte(byte(v >> 16))
	b.writeByte(byte(v >> 8))
	b.writeByte(byte(v))
}

func (b *buffer) bytes() []byte {
	return b.b[:b.pos]
}
