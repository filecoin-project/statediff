package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Link__MarketV0RawDealProposal struct{ x ipld.Link }
type Link__MarketV0RawDealProposal = *_Link__MarketV0RawDealProposal
func (n Link__MarketV0RawDealProposal) Link() ipld.Link {
	return n.x
}
func (_Link__MarketV0RawDealProposal__Prototype) FromLink(v ipld.Link) (Link__MarketV0RawDealProposal, error) {
	n := _Link__MarketV0RawDealProposal{v}
	return &n, nil
}
type _Link__MarketV0RawDealProposal__Maybe struct {
	m schema.Maybe
	v Link__MarketV0RawDealProposal
}
type MaybeLink__MarketV0RawDealProposal = *_Link__MarketV0RawDealProposal__Maybe

func (m MaybeLink__MarketV0RawDealProposal) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeLink__MarketV0RawDealProposal) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeLink__MarketV0RawDealProposal) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeLink__MarketV0RawDealProposal) AsNode() ipld.Node {
	switch m.m {
		case schema.Maybe_Absent:
			return ipld.Absent
		case schema.Maybe_Null:
			return ipld.Null
		case schema.Maybe_Value:
			return m.v
		default:
			panic("unreachable")
	}
}
func (m MaybeLink__MarketV0RawDealProposal) Must() Link__MarketV0RawDealProposal {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (Link__MarketV0RawDealProposal)(&_Link__MarketV0RawDealProposal{})
var _ schema.TypedNode = (Link__MarketV0RawDealProposal)(&_Link__MarketV0RawDealProposal{})
func (Link__MarketV0RawDealProposal) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Link
}
func (Link__MarketV0RawDealProposal) LookupByString(string) (ipld.Node, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.LookupByString("")
}
func (Link__MarketV0RawDealProposal) LookupByNode(ipld.Node) (ipld.Node, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.LookupByNode(nil)
}
func (Link__MarketV0RawDealProposal) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.LookupByIndex(0)
}
func (Link__MarketV0RawDealProposal) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.LookupBySegment(seg)
}
func (Link__MarketV0RawDealProposal) MapIterator() ipld.MapIterator {
	return nil
}
func (Link__MarketV0RawDealProposal) ListIterator() ipld.ListIterator {
	return nil
}
func (Link__MarketV0RawDealProposal) Length() int {
	return -1
}
func (Link__MarketV0RawDealProposal) IsAbsent() bool {
	return false
}
func (Link__MarketV0RawDealProposal) IsNull() bool {
	return false
}
func (Link__MarketV0RawDealProposal) AsBool() (bool, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.AsBool()
}
func (Link__MarketV0RawDealProposal) AsInt() (int, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.AsInt()
}
func (Link__MarketV0RawDealProposal) AsFloat() (float64, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.AsFloat()
}
func (Link__MarketV0RawDealProposal) AsString() (string, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.AsString()
}
func (Link__MarketV0RawDealProposal) AsBytes() ([]byte, error) {
	return mixins.Link{"types.Link__MarketV0RawDealProposal"}.AsBytes()
}
func (n Link__MarketV0RawDealProposal) AsLink() (ipld.Link, error) {
	return n.x, nil
}
func (Link__MarketV0RawDealProposal) Prototype() ipld.NodePrototype {
	return _Link__MarketV0RawDealProposal__Prototype{}
}
type _Link__MarketV0RawDealProposal__Prototype struct{}

func (_Link__MarketV0RawDealProposal__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Link__MarketV0RawDealProposal__Builder
	nb.Reset()
	return &nb
}
type _Link__MarketV0RawDealProposal__Builder struct {
	_Link__MarketV0RawDealProposal__Assembler
}
func (nb *_Link__MarketV0RawDealProposal__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Link__MarketV0RawDealProposal__Builder) Reset() {
	var w _Link__MarketV0RawDealProposal
	var m schema.Maybe
	*nb = _Link__MarketV0RawDealProposal__Builder{_Link__MarketV0RawDealProposal__Assembler{w: &w, m: &m}}
}
type _Link__MarketV0RawDealProposal__Assembler struct {
	w *_Link__MarketV0RawDealProposal
	m *schema.Maybe
}

func (na *_Link__MarketV0RawDealProposal__Assembler) reset() {}
func (_Link__MarketV0RawDealProposal__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.BeginMap(0)
}
func (_Link__MarketV0RawDealProposal__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.BeginList(0)
}
func (na *_Link__MarketV0RawDealProposal__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	panic("unreachable")
}
func (_Link__MarketV0RawDealProposal__Assembler) AssignBool(bool) error {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.AssignBool(false)
}
func (_Link__MarketV0RawDealProposal__Assembler) AssignInt(int) error {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.AssignInt(0)
}
func (_Link__MarketV0RawDealProposal__Assembler) AssignFloat(float64) error {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.AssignFloat(0)
}
func (_Link__MarketV0RawDealProposal__Assembler) AssignString(string) error {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.AssignString("")
}
func (_Link__MarketV0RawDealProposal__Assembler) AssignBytes([]byte) error {
	return mixins.LinkAssembler{"types.Link__MarketV0RawDealProposal"}.AssignBytes(nil)
}
func (na *_Link__MarketV0RawDealProposal__Assembler) AssignLink(v ipld.Link) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	if na.w == nil {
		na.w = &_Link__MarketV0RawDealProposal{}
	}
	na.w.x = v
	*na.m = schema.Maybe_Value
	return nil
}
func (na *_Link__MarketV0RawDealProposal__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Link__MarketV0RawDealProposal); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		}
		if na.w == nil {
			na.w = v2
			*na.m = schema.Maybe_Value
			return nil
		}
		*na.w = *v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v2, err := v.AsLink(); err != nil {
		return err
	} else {
		return na.AssignLink(v2)
	}
}
func (_Link__MarketV0RawDealProposal__Assembler) Prototype() ipld.NodePrototype {
	return _Link__MarketV0RawDealProposal__Prototype{}
}
func (Link__MarketV0RawDealProposal) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Link__MarketV0RawDealProposal) LinkTargetNodePrototype() ipld.NodePrototype {
	return Type.Link__MarketV0RawDealProposal__Repr
}
func (n Link__MarketV0RawDealProposal) Representation() ipld.Node {
	return (*_Link__MarketV0RawDealProposal__Repr)(n)
}
type _Link__MarketV0RawDealProposal__Repr = _Link__MarketV0RawDealProposal
var _ ipld.Node = &_Link__MarketV0RawDealProposal__Repr{}
type _Link__MarketV0RawDealProposal__ReprPrototype = _Link__MarketV0RawDealProposal__Prototype
type _Link__MarketV0RawDealProposal__ReprAssembler = _Link__MarketV0RawDealProposal__Assembler