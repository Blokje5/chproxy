package protocol

type Proto interface {
	// ExpectedSize after encoding of the Proto struct. Used to initialize a buffer.
	// ExpectedSize() int

	// Encode encodes the Proto to the ClickHouse native protocol.
	Encode(w *Writer)

	// Decode decodes the proto from the ClickHouse native protocol.
	Decode(r *Reader) error
}

// Client Packet Types
// See: https://github.com/ClickHouse/ClickHouse/blob/master/src/Core/Protocol.h
type ClientCode byte

const (
	ClientCodeHello           ClientCode = 0
	ClientCodeQuery           ClientCode = 1
	ClientCodeData            ClientCode = 2
	ClientCodeCancel          ClientCode = 3
	ClientCodePing            ClientCode = 4
	ClientTablesStatusRequest ClientCode = 5
	ClientKeepAlive           ClientCode = 6
	ClientScalar              ClientCode = 7
)

func (c ClientCode) Encode(w *Writer) {
	w.PutByte(byte(c))
}

// Server Packet Types
// See: https://github.com/ClickHouse/ClickHouse/blob/master/src/Core/Protocol.h
type ServerCode byte

const (
	ServerCodeHello       ServerCode = 0
	ServerCodeData        ServerCode = 1
	ServerCodeException   ServerCode = 2
	ServerCodeProgress    ServerCode = 3
	ServerCodePong        ServerCode = 4
	ServerCodeEndOfStream ServerCode = 5
	ServerCodeProfileInfo ServerCode = 6
	ServerCodeTotals      ServerCode = 7
	ServerCodeExtremes    ServerCode = 8
)

func (c ServerCode) Encode(w *Writer) {
	w.PutByte(byte(c))
}
