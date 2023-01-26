package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	const (
		vuint32 = uint32(4)
		vuint64 = uint64(5)
	)

	expectedSize := 1 + 4 + 8
	
	codec := NewCodec()
	
	buffer := make([]byte, expectedSize)
	writer := codec.NewWriter(buffer)

	ClientCodeHello.Encode(writer)
	writer.PutUint32(vuint32)
	writer.PutUint64(vuint64)

	reader := codec.NewReader(buffer)

	code, err := reader.ReadByte()
	assert.NoError(t, err)
	assert.Equal(t, ClientCodeHello, ClientCode(code))

	resuint32, err := reader.ReadUint32()
	assert.NoError(t, err)
	assert.Equal(t, vuint32, resuint32)

	resuint64, err := reader.ReadUint64()
	assert.NoError(t, err)
	assert.Equal(t, vuint64, resuint64)
}