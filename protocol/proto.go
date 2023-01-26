package protocol

type Proto interface {
	// ExpectedSize after encoding of the Proto struct. Used to initialize a buffer.
	ExpectedSize() int

	// Encode encodes the Proto to the ClickHouse native protocol.
	Encode(w *Writer)

	// Decode decodes the proto from the ClickHouse native protocol.
	Decode(r *Reader) error
}
