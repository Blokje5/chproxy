package protocol

import "fmt"

var _ Proto = &ClientHelloProto{}

type ClientHelloProto struct {
	Name string

	Major uint32
	Minor uint32

	ProtocolVersion uint32

	Database string
	User     string
	Password string
}

func (c *ClientHelloProto) Encode(w *Writer) {
	ClientCodeHello.Encode(w)
	w.PutString(c.Name)
	w.PutUint32(c.Major)
	w.PutUint32(c.Minor)
	w.PutUint32(c.ProtocolVersion)
	w.PutString(c.Database)
	w.PutString(c.User)
	w.PutString(c.Password)
}

func (c *ClientHelloProto) Decode(r *Reader) error {
	name, err := r.ReadString()
	if err != nil {
		return fmt.Errorf("ClientHelloProto.Name: %w", err)
	}

	c.Name = name

	major, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ClientHelloProto.Major: %w", err)
	}

	c.Major = major

	minor, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ClientHelloProto.Minor: %w", err)
	}

	c.Minor = minor

	protocolVersion, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ClientHelloProto.ProtocolVersion: %w", err)
	}

	c.ProtocolVersion = protocolVersion

	user, err := r.ReadString()
	if err != nil {
		return fmt.Errorf("ClientHelloProto.User: %w", err)
	}

	c.User = user

	password, err := r.ReadString()
	if err != nil {
		return fmt.Errorf("ClientHelloProto.Password: %w", err)
	}

	c.Password = password

	return nil
}
