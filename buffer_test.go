package pq

import (
	"testing"
	"github.com/bmizerany/assert"
)

func TestBufferSimple(t *testing.T) {
	buf := newBuffer()
	buf.setType('X')
	buf.writeString("testing")
	buf.writeInt16(1)
	buf.writeInt32(2)
	buf.setLength()

	assert.Equal(t, "X\x00\x00\x00\x12testing\x00\x00\x01\x00\x00\x00\x02", string(buf.bytes()))
}
