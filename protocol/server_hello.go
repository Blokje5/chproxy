package protocol

import "fmt"

type ServerHelloProto struct {
	Name        string
	Major       uint32
	Minor       uint32
	Revision    uint32
	Timezone    string
	DisplayName string
	Patch       uint32
}

func (p *ServerHelloProto) Encode(w *Writer) {
	ServerCodeHello.Encode(w)
	w.PutString(p.Name)
	w.PutUint32(p.Major)
	w.PutUint32(p.Minor)
	w.PutUint32(p.Revision)
	w.PutString(p.Timezone)
	w.PutString(p.DisplayName)
	w.PutUint32(p.Patch)
}

func (p *ServerHelloProto) Decode(r *Reader) error {
	name, err := r.ReadString()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.Name: %w", err)
	}

	p.Name = name

	major, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.Major: %w", err)
	}

	p.Major = major

	minor, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.Minor: %w", err)
	}

	p.Minor = minor

	revision, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.Revision: %w", err)
	}

	p.Revision = revision

	timezone, err := r.ReadString()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.Timezone: %w", err)
	}

	p.Timezone = timezone

	displayName, err := r.ReadString()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.DisplayName: %w", err)
	}

	p.DisplayName = displayName

	patch, err := r.ReadUint32()
	if err != nil {
		return fmt.Errorf("ServerHelloProto.Patch: %w", err)
	}

	p.Patch = patch

	return nil
}
