package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Link__PowerV0CronEvent struct{ x ipld.Link }
type Link__PowerV0CronEvent = *_Link__PowerV0CronEvent
func (n Link__PowerV0CronEvent) Link() ipld.Link {
	return n.x
}
func (_Link__PowerV0CronEvent__Prototype) FromLink(v ipld.Link) (Link__PowerV0CronEvent, error) {
	n := _Link__PowerV0CronEvent{v}
	return &n, nil
}
type _Link__PowerV0CronEvent__Maybe struct {
	m schema.Maybe
	v Link__PowerV0CronEvent
}
type MaybeLink__PowerV0CronEvent = *_Link__PowerV0CronEvent__Maybe

func (m MaybeLink__PowerV0CronEvent) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeLink__PowerV0CronEvent) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeLink__PowerV0CronEvent) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeLink__PowerV0CronEvent) AsNode() ipld.Node {
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
func (m MaybeLink__PowerV0CronEvent) Must() Link__PowerV0CronEvent {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (Link__PowerV0CronEvent)(&_Link__PowerV0CronEvent{})
var _ schema.TypedNode = (Link__PowerV0CronEvent)(&_Link__PowerV0CronEvent{})
func (Link__PowerV0CronEvent) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Link
}
func (Link__PowerV0CronEvent) LookupByString(string) (ipld.Node, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.LookupByString("")
}
func (Link__PowerV0CronEvent) LookupByNode(ipld.Node) (ipld.Node, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.LookupByNode(nil)
}
func (Link__PowerV0CronEvent) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.LookupByIndex(0)
}
func (Link__PowerV0CronEvent) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.LookupBySegment(seg)
}
func (Link__PowerV0CronEvent) MapIterator() ipld.MapIterator {
	return nil
}
func (Link__PowerV0CronEvent) ListIterator() ipld.ListIterator {
	return nil
}
func (Link__PowerV0CronEvent) Length() int {
	return -1
}
func (Link__PowerV0CronEvent) IsAbsent() bool {
	return false
}
func (Link__PowerV0CronEvent) IsNull() bool {
	return false
}
func (Link__PowerV0CronEvent) AsBool() (bool, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.AsBool()
}
func (Link__PowerV0CronEvent) AsInt() (int, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.AsInt()
}
func (Link__PowerV0CronEvent) AsFloat() (float64, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.AsFloat()
}
func (Link__PowerV0CronEvent) AsString() (string, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.AsString()
}
func (Link__PowerV0CronEvent) AsBytes() ([]byte, error) {
	return mixins.Link{"types.Link__PowerV0CronEvent"}.AsBytes()
}
func (n Link__PowerV0CronEvent) AsLink() (ipld.Link, error) {
	return n.x, nil
}
func (Link__PowerV0CronEvent) Prototype() ipld.NodePrototype {
	return _Link__PowerV0CronEvent__Prototype{}
}
type _Link__PowerV0CronEvent__Prototype struct{}

func (_Link__PowerV0CronEvent__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Link__PowerV0CronEvent__Builder
	nb.Reset()
	return &nb
}
type _Link__PowerV0CronEvent__Builder struct {
	_Link__PowerV0CronEvent__Assembler
}
func (nb *_Link__PowerV0CronEvent__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Link__PowerV0CronEvent__Builder) Reset() {
	var w _Link__PowerV0CronEvent
	var m schema.Maybe
	*nb = _Link__PowerV0CronEvent__Builder{_Link__PowerV0CronEvent__Assembler{w: &w, m: &m}}
}
type _Link__PowerV0CronEvent__Assembler struct {
	w *_Link__PowerV0CronEvent
	m *schema.Maybe
}

func (na *_Link__PowerV0CronEvent__Assembler) reset() {}
func (_Link__PowerV0CronEvent__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.BeginMap(0)
}
func (_Link__PowerV0CronEvent__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.BeginList(0)
}
func (na *_Link__PowerV0CronEvent__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	panic("unreachable")
}
func (_Link__PowerV0CronEvent__Assembler) AssignBool(bool) error {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.AssignBool(false)
}
func (_Link__PowerV0CronEvent__Assembler) AssignInt(int) error {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.AssignInt(0)
}
func (_Link__PowerV0CronEvent__Assembler) AssignFloat(float64) error {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.AssignFloat(0)
}
func (_Link__PowerV0CronEvent__Assembler) AssignString(string) error {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.AssignString("")
}
func (_Link__PowerV0CronEvent__Assembler) AssignBytes([]byte) error {
	return mixins.LinkAssembler{"types.Link__PowerV0CronEvent"}.AssignBytes(nil)
}
func (na *_Link__PowerV0CronEvent__Assembler) AssignLink(v ipld.Link) error {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	}
	if na.w == nil {
		na.w = &_Link__PowerV0CronEvent{}
	}
	na.w.x = v
	*na.m = schema.Maybe_Value
	return nil
}
func (na *_Link__PowerV0CronEvent__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Link__PowerV0CronEvent); ok {
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
func (_Link__PowerV0CronEvent__Assembler) Prototype() ipld.NodePrototype {
	return _Link__PowerV0CronEvent__Prototype{}
}
func (Link__PowerV0CronEvent) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (Link__PowerV0CronEvent) LinkTargetNodePrototype() ipld.NodePrototype {
	return Type.Link__PowerV0CronEvent__Repr
}
func (n Link__PowerV0CronEvent) Representation() ipld.Node {
	return (*_Link__PowerV0CronEvent__Repr)(n)
}
type _Link__PowerV0CronEvent__Repr = _Link__PowerV0CronEvent
var _ ipld.Node = &_Link__PowerV0CronEvent__Repr{}
type _Link__PowerV0CronEvent__ReprPrototype = _Link__PowerV0CronEvent__Prototype
type _Link__PowerV0CronEvent__ReprAssembler = _Link__PowerV0CronEvent__Assembler