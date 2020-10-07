package statediff

// This class provides a wrapper around bitfield.BitField
// which json Marshal's to include a hint at the type to allow
// clients to provide cleaner rendering of such.

import (
	"bytes"
	"io"
)

// CBORBytes wraps already-serialized bytes as CBOR-marshalable.
type CBORBytes []byte

// MarshalCBOR turns these bytes into cbor-prefixed
func (b CBORBytes) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(b)
	return err
}

// UnmarshalCBOR reads the bytes.
func (b *CBORBytes) UnmarshalCBOR(r io.Reader) error {
	var c bytes.Buffer
	_, err := c.ReadFrom(r)
	*b = c.Bytes()
	return err
}
