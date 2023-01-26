package protocol

const (
	Uint32Size = 4
	Uint64Size = 8
	MaxStringSize = 1024 * 1024 * 10
)

const (
	ClickhouseBinaryVersionV1 = iota
)

type Codec struct {
	version int
}

func NewCodec() Codec {
	return Codec{
		version: ClickhouseBinaryVersionV1,
	}
}

func (c *Codec) NewReader(buf []byte) *Reader {
	switch c.version {
	case ClickhouseBinaryVersionV1: return NewReader(buf)
	}

	return nil
}

func (c *Codec) NewWriter(buf []byte) *Writer {
	switch c.version {
	case ClickhouseBinaryVersionV1: return NewWriter(buf)
	}

	return nil
}