package statediff

// This class provides a wrapper around bitfield.BitField
// which json Marshal's to include a hint at the type to allow
// clients to provide cleaner rendering of such.

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"io"

	"github.com/filecoin-project/go-bitfield"
)

type JSONBitField struct {
	bitfield.BitField
}

type jsonField struct {
	T string `json:"_type"`
	B string `json:"bytes"`
}

func (j JSONBitField) MarshalJSON() ([]byte, error) {
	b := bytes.NewBuffer([]byte{})
	if err := j.MarshalCBOR(b); err != nil {
		return nil, err
	}
	return json.Marshal(jsonField{
		T: "bitfield",
		B: hex.EncodeToString(b.Bytes()),
	})
}

// Wraps already-serialized bytes as CBOR-marshalable.
type CBORBytes []byte

func (b CBORBytes) MarshalCBOR(w io.Writer) error {
	_, err := w.Write(b)
	return err
}

func (b *CBORBytes) UnmarshalCBOR(r io.Reader) error {
	var c bytes.Buffer
	_, err := c.ReadFrom(r)
	*b = c.Bytes()
	return err
}
