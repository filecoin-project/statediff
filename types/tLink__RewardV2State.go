package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Link__RewardV2State struct{ x ipld.Link }
type Link__RewardV2State = *_Link__RewardV2State
func (n Link__RewardV2State) Link() ipld.Link {
	return n.x
}
func (_Link__RewardV2State__Prototype) FromLink(v ipld.Link) (Link__RewardV2State, error) {
	n := _Link__RewardV2State{v}
	return &n, nil
}
type _Link__RewardV2State__Maybe struct {
	m schema.Maybe
	v Link__RewardV2State
}
type MaybeLink__RewardV2State = *_Link__RewardV2State__Maybe

func (m MaybeLink__RewardV2State) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeLink__RewardV2State) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeLink__RewardV2State) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeLink__RewardV2State) AsNode() ipld.Node {
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
func (m MaybeLink__RewardV2State) Must() Link__RewardV2State {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (Link__RewardV2State)(&_Link__RewardV2State{})
var _ schema.TypedNode = (Link__RewardV2State)(&_Link__RewardV2State{})
func (Link__RewardV2State) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Link
}
func (Link__RewardV2State) LookupByString(string) (ipld.Node, error) {
	return mixins.Link{"types.Link__RewardV2State"}.LookupByString("")
}
func (Link__RewardV2State) LookupByNode(ipld.Node) (ipld.Node, error) {
	return mixins.Link{"types.Link__RewardV2State"}.LookupByNode(nil)
}
func (Link__RewardV2State) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Link{"types.Link__RewardV2State"}.LookupByIndex(0)
}
func (Link__RewardV2State) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return mixins.Link{"types.Link__RewardV2State"}.LookupBySegment(seg)
}
func (Link__RewardV2State) MapIterator() ipld.MapIterator {
	return nil
}
func (Link__RewardV2State) ListIterator() ipld.ListIterator {
	return nil
}
func (Link__RewardV2State) Length() int {
	return -1
}
func (Link__RewardV2State) IsAbsent() bool {
	return false
}
func (Link__RewardV2State) IsNull() bool {
	return false
}
func (Link__RewardV2State) AsBool() (bool, error) {
	return mixins.Link{"types.Link__RewardV2State"}.AsBool()
}
func (Link__RewardV2State) AsInt() (int, error) {
	return mixins.Link{"types.Link__RewardV2State"}.AsInt()
}
func (Link__RewardV2State) AsFloat() (float64, error) {
	return mixins.Link{"types.Link__RewardV2State"}.AsFloat()
}
func (Link__RewardV2State) AsString() (string, error) {
	return mixins.Link{"types.Link__RewardV2State"}.AsString()
}
func (Link__RewardV2State) AsBytes() ([]byte, error) {
	return mixins.Link{"types.Link__RewardV2State"}.AsBytes()
}
func (n Link__RewardV2State) AsLink() (ipld.Link, error) {
	return n.x, nil
}
func (Link__RewardV2State) Prototype() ipld.NodePrototype {
	return _Link__RewardV2State__Prototype{}
}
type _Link__RewardV2State__Prototype struct{}

func (_Link__RewardV2State__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Link__RewardV2State__Builder
	nb.Reset()
	return &nb
}
type _Link__RewardV2State__Builder struct {
	_Link__RewardV2State__Assembler
}
func (nb *_Link__RewardV2State__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Link__RewardV2State__Builder) Reset() {
	var w _Link__RewardV2State
	var m schema.Maybe
	*nb = _Link__RewardV2State__Builder{_Link__RewardV2State__Assembler{w: &w, m: &m}}
}
type _Link__RewardV2State__Assembler struct {
	w *_Link__RewardV2State
	m *schema.Maybe
}

func (na *_Link__RewardV2State__Assembler) reset() {}
func (_Link__RewardV2State__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.BeginMap(0)
}
func (_Link__RewardV2State__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.BeginList(0)
}
func (na *_Link__RewardV2State__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.LinkAssembler{"types.Link__RewardV2State"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	panic("unreachable")
}
func (_Link__RewardV2State__Assembler) AssignBool(bool) error {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.AssignBool(false)
}
func (_Link__RewardV2State__Assembler) AssignInt(int) error {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.AssignInt(0)
}
func (_Link__RewardV2State__Assembler) AssignFloat(float64) error {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.AssignFloat(0)
}
func (_Link__RewardV2State__Assembler) AssignString(string) error {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.AssignString("")
}
func (_Link__RewardV2State__Assembler) AssignBytes([]byte) error {
	return mixins.LinkAssembler{"types.Link__RewardV2State"}.AssignBytes(nil)
}
func (na *_Link__RewardV2State__Assembler) AssignLink(v ipld.Link) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	if na.w == nil {
		na.w = &_Link__RewardV2State{}
	}
	na.w.x = v
	*na.m = schema.Maybe_Value
	return nil
}
func (na *_Link__RewardV2State__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Link__RewardV2State); ok {
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
func (_Link__RewardV2State__Assembler) Prototype() ipld.NodePrototype {
	return _Link__RewardV2State__Prototype{}
}
func (Link__RewardV2State) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Link__RewardV2State) LinkTargetNodePrototype() ipld.NodePrototype {
	return Type.Link__RewardV2State__Repr
}
func (n Link__RewardV2State) Representation() ipld.Node {
	return (*_Link__RewardV2State__Repr)(n)
}
type _Link__RewardV2State__Repr = _Link__RewardV2State
var _ ipld.Node = &_Link__RewardV2State__Repr{}
type _Link__RewardV2State__ReprPrototype = _Link__RewardV2State__Prototype
type _Link__RewardV2State__ReprAssembler = _Link__RewardV2State__Assembler