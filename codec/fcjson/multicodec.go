package fcjson

import (
	"io"

	"github.com/polydawn/refmt/json"

	ipld "github.com/ipld/go-ipld-prime"
)

func Encode(n ipld.Node, w io.Writer) error {
	return Marshal(n, json.NewEncoder(w, json.EncodeOptions{
		Line:   []byte{'\n'},
		Indent: []byte{'\t'},
	}))
}

func (d *DagMarshaler) Encode(n ipld.Node, w io.Writer) error {
	return d.MarshalRecursive(n, json.NewEncoder(w, json.EncodeOptions{
		Line:   []byte{'\n'},
		Indent: []byte{'\t'},
	}))
}
