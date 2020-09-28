package fcjson

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/polydawn/refmt/shared"
	"github.com/polydawn/refmt/tok"

	ipld "github.com/ipld/go-ipld-prime"
	cidlink "github.com/ipld/go-ipld-prime/linking/cid"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/statediff/types"
)

// This is dagjson with special pretty-print sauce.

func Marshal(n ipld.Node, sink shared.TokenSink) error {
	var tk tok.Token
	switch n.ReprKind() {
	case ipld.ReprKind_Invalid:
		return fmt.Errorf("cannot traverse a node that is absent")
	case ipld.ReprKind_Null:
		tk.Type = tok.TNull
		_, err := sink.Step(&tk)
		return err
	case ipld.ReprKind_Map:
		// Emit start of map.
		tk.Type = tok.TMapOpen
		tk.Length = n.Length()
		if _, err := sink.Step(&tk); err != nil {
			return err
		}
		// Emit map contents (and recurse).
		for itr := n.MapIterator(); !itr.Done(); {
			k, v, err := itr.Next()
			if err != nil {
				return err
			}
			tk.Type = tok.TString
			tk.Str, err = k.AsString()
			if err != nil {
				return err
			}
			if _, ok := k.(types.RawAddress); ok {
				a, err := address.NewFromBytes([]byte(tk.Str))
				if err != nil {
					return err
				}
				tk.Str = a.String()
			}
			if _, err := sink.Step(&tk); err != nil {
				return err
			}
			if err := Marshal(v, sink); err != nil {
				return err
			}
		}
		// Emit map close.
		tk.Type = tok.TMapClose
		_, err := sink.Step(&tk)
		return err
	case ipld.ReprKind_List:
		// Emit start of list.
		tk.Type = tok.TArrOpen
		l := n.Length()
		tk.Length = l
		if _, err := sink.Step(&tk); err != nil {
			return err
		}
		// Emit list contents (and recurse).
		for i := 0; i < l; i++ {
			v, err := n.LookupByIndex(i)
			if err != nil {
				return err
			}
			if err := Marshal(v, sink); err != nil {
				return err
			}
		}
		// Emit list close.
		tk.Type = tok.TArrClose
		_, err := sink.Step(&tk)
		return err
	case ipld.ReprKind_Bool:
		v, err := n.AsBool()
		if err != nil {
			return err
		}
		tk.Type = tok.TBool
		tk.Bool = v
		_, err = sink.Step(&tk)
		return err
	case ipld.ReprKind_Int:
		v, err := n.AsInt()
		if err != nil {
			return err
		}
		tk.Type = tok.TInt
		tk.Int = int64(v)
		_, err = sink.Step(&tk)
		return err
	case ipld.ReprKind_Float:
		v, err := n.AsFloat()
		if err != nil {
			return err
		}
		tk.Type = tok.TFloat64
		tk.Float64 = v
		_, err = sink.Step(&tk)
		return err
	case ipld.ReprKind_String:
		v, err := n.AsString()
		if err != nil {
			return err
		}
		if _, ok := n.(types.RawAddress); ok {
			a, err := address.NewFromBytes([]byte(v))
			if err != nil {
				return err
			}
			tk.Str = a.String()
		} else {
			tk.Str = v
		}
		tk.Type = tok.TString
		_, err = sink.Step(&tk)
		return err
	case ipld.ReprKind_Bytes:
		tk.Type = tok.TString
		v, err := n.AsBytes()
		if err != nil {
			return err
		}

		if _, ok := n.(types.Address); ok {
			a, err := address.NewFromBytes(v)
			if err != nil {
				return err
			}
			tk.Str = a.String()
		} else if _, ok := n.(types.BigInt); ok {
			i := big.NewInt(0)
			i.SetBytes(v)
			tk.Str = i.String()
		} else if _, ok := n.(types.BitField); ok {
			b, err := bitfield.NewFromBytes(v)
			if err != nil {
				if err = b.UnmarshalCBOR(bytes.NewBuffer(v)); err != nil {
					return err
				}
			}
			tk.Type = tok.TMapOpen
			tk.Length = 2
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Type = tok.TString
			tk.Str = "_type"
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Str = "bitfield"
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Str = "bytes"
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			buf := bytes.NewBuffer([]byte{})
			b.MarshalCBOR(buf)
			tk.Str = hex.EncodeToString(buf.Bytes())
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Type = tok.TMapClose
		} else {
			tk.Str = base64.StdEncoding.EncodeToString(v)
		}
		_, err = sink.Step(&tk)
		return err
	case ipld.ReprKind_Link:
		v, err := n.AsLink()
		if err != nil {
			return err
		}
		switch lnk := v.(type) {
		case cidlink.Link:
			// Precisely four tokens to emit:
			tk.Type = tok.TMapOpen
			tk.Length = 1
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Type = tok.TString
			tk.Str = "/"
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Str = lnk.Cid.String()
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			tk.Type = tok.TMapClose
			if _, err = sink.Step(&tk); err != nil {
				return err
			}
			return nil
		default:
			return fmt.Errorf("schemafree link emission only supported by this codec for CID type links")
		}
	default:
		panic("unreachable")
	}
}
