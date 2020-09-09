package statediff

// This class provides a wrapper around bitfield.BitField
// which json Marshal's to include a hint at the type to allow
// clients to provide cleaner rendering of such.

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
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
