package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _List__MinerExpirationExtend struct {
	x []_MinerExpirationExtend
}
type List__MinerExpirationExtend = *_List__MinerExpirationExtend
type _List__MinerExpirationExtend__Maybe struct {
	m schema.Maybe
	v List__MinerExpirationExtend
}
type MaybeList__MinerExpirationExtend = *_List__MinerExpirationExtend__Maybe

func (m MaybeList__MinerExpirationExtend) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeList__MinerExpirationExtend) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeList__MinerExpirationExtend) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeList__MinerExpirationExtend) AsNode() ipld.Node {
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
func (m MaybeList__MinerExpirationExtend) Must() List__MinerExpirationExtend {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (List__MinerExpirationExtend)(&_List__MinerExpirationExtend{})
var _ schema.TypedNode = (List__MinerExpirationExtend)(&_List__MinerExpirationExtend{})
func (List__MinerExpirationExtend) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (List__MinerExpirationExtend) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.LookupByString("")
}
func (n List__MinerExpirationExtend) LookupByNode(k ipld.Node) (ipld.Node, error) {
	idx, err := k.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(idx)
}
func (n List__MinerExpirationExtend) LookupByIndex(idx int) (ipld.Node, error) {
	if n.Length() <= idx {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(idx)}
	}
	v := &n.x[idx]
	return v, nil
}
func (n List__MinerExpirationExtend) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.List__MinerExpirationExtend", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (List__MinerExpirationExtend) MapIterator() ipld.MapIterator {
	return nil
}
func (n List__MinerExpirationExtend) ListIterator() ipld.ListIterator {
	return &_List__MinerExpirationExtend__ListItr{n, 0}
}

type _List__MinerExpirationExtend__ListItr struct {
	n List__MinerExpirationExtend
	idx  int
}

func (itr *_List__MinerExpirationExtend__ListItr) Next() (idx int, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.x) {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	x := &itr.n.x[itr.idx]
	v = x
	itr.idx++
	return
}
func (itr *_List__MinerExpirationExtend__ListItr) Done() bool {
	return itr.idx >= len(itr.n.x)
}

func (n List__MinerExpirationExtend) Length() int {
	return len(n.x)
}
func (List__MinerExpirationExtend) IsAbsent() bool {
	return false
}
func (List__MinerExpirationExtend) IsNull() bool {
	return false
}
func (List__MinerExpirationExtend) AsBool() (bool, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.AsBool()
}
func (List__MinerExpirationExtend) AsInt() (int, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.AsInt()
}
func (List__MinerExpirationExtend) AsFloat() (float64, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.AsFloat()
}
func (List__MinerExpirationExtend) AsString() (string, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.AsString()
}
func (List__MinerExpirationExtend) AsBytes() ([]byte, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.AsBytes()
}
func (List__MinerExpirationExtend) AsLink() (ipld.Link, error) {
	return mixins.List{"types.List__MinerExpirationExtend"}.AsLink()
}
func (List__MinerExpirationExtend) Prototype() ipld.NodePrototype {
	return _List__MinerExpirationExtend__Prototype{}
}
type _List__MinerExpirationExtend__Prototype struct{}

func (_List__MinerExpirationExtend__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _List__MinerExpirationExtend__Builder
	nb.Reset()
	return &nb
}
type _List__MinerExpirationExtend__Builder struct {
	_List__MinerExpirationExtend__Assembler
}
func (nb *_List__MinerExpirationExtend__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_List__MinerExpirationExtend__Builder) Reset() {
	var w _List__MinerExpirationExtend
	var m schema.Maybe
	*nb = _List__MinerExpirationExtend__Builder{_List__MinerExpirationExtend__Assembler{w: &w, m: &m}}
}
type _List__MinerExpirationExtend__Assembler struct {
	w *_List__MinerExpirationExtend
	m *schema.Maybe
	state laState

	cm schema.Maybe
	va _MinerExpirationExtend__Assembler
}

func (na *_List__MinerExpirationExtend__Assembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_List__MinerExpirationExtend__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.BeginMap(0)
}
func (na *_List__MinerExpirationExtend__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_List__MinerExpirationExtend{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_MinerExpirationExtend, 0, sizeHint)
	}
	return na, nil
}
func (na *_List__MinerExpirationExtend__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_List__MinerExpirationExtend__Assembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignBool(false)
}
func (_List__MinerExpirationExtend__Assembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignInt(0)
}
func (_List__MinerExpirationExtend__Assembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignFloat(0)
}
func (_List__MinerExpirationExtend__Assembler) AssignString(string) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignString("")
}
func (_List__MinerExpirationExtend__Assembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignBytes(nil)
}
func (_List__MinerExpirationExtend__Assembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend"}.AssignLink(nil)
}
func (na *_List__MinerExpirationExtend__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_List__MinerExpirationExtend); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
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
	if v.ReprKind() != ipld.ReprKind_List {
		return ipld.ErrWrongKind{TypeName: "types.List__MinerExpirationExtend", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
	}
	itr := v.ListIterator()
	for !itr.Done() {
		_, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_List__MinerExpirationExtend__Assembler) Prototype() ipld.NodePrototype {
	return _List__MinerExpirationExtend__Prototype{}
}
func (la *_List__MinerExpirationExtend__Assembler) valueFinishTidy() bool {
	switch la.cm {
	case schema.Maybe_Value:
		la.va.w = nil
		la.cm = schema.Maybe_Absent
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_List__MinerExpirationExtend__Assembler) AssembleValue() ipld.NodeAssembler {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when still in the middle of assembling the previous value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	la.w.x = append(la.w.x, _MinerExpirationExtend{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_List__MinerExpirationExtend__Assembler) Finish() error {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	la.state = laState_finished
	*la.m = schema.Maybe_Value
	return nil
}
func (la *_List__MinerExpirationExtend__Assembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _MinerExpirationExtend__Prototype{}
}
func (List__MinerExpirationExtend) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n List__MinerExpirationExtend) Representation() ipld.Node {
	return (*_List__MinerExpirationExtend__Repr)(n)
}
type _List__MinerExpirationExtend__Repr _List__MinerExpirationExtend
var _ ipld.Node = &_List__MinerExpirationExtend__Repr{}
func (_List__MinerExpirationExtend__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_List__MinerExpirationExtend__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.LookupByString("")
}
func (nr *_List__MinerExpirationExtend__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (List__MinerExpirationExtend)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(MinerExpirationExtend).Representation(), nil
}
func (nr *_List__MinerExpirationExtend__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	v, err := (List__MinerExpirationExtend)(nr).LookupByIndex(idx)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(MinerExpirationExtend).Representation(), nil
}
func (n _List__MinerExpirationExtend__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.List__MinerExpirationExtend.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_List__MinerExpirationExtend__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (nr *_List__MinerExpirationExtend__Repr) ListIterator() ipld.ListIterator {
	return &_List__MinerExpirationExtend__ReprListItr{(List__MinerExpirationExtend)(nr), 0}
}

type _List__MinerExpirationExtend__ReprListItr _List__MinerExpirationExtend__ListItr

func (itr *_List__MinerExpirationExtend__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	idx, v, err = (*_List__MinerExpirationExtend__ListItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return idx, v.(MinerExpirationExtend).Representation(), nil
}
func (itr *_List__MinerExpirationExtend__ReprListItr) Done() bool {
	return (*_List__MinerExpirationExtend__ListItr)(itr).Done()
}

func (rn *_List__MinerExpirationExtend__Repr) Length() int {
	return len(rn.x)
}
func (_List__MinerExpirationExtend__Repr) IsAbsent() bool {
	return false
}
func (_List__MinerExpirationExtend__Repr) IsNull() bool {
	return false
}
func (_List__MinerExpirationExtend__Repr) AsBool() (bool, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.AsBool()
}
func (_List__MinerExpirationExtend__Repr) AsInt() (int, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.AsInt()
}
func (_List__MinerExpirationExtend__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.AsFloat()
}
func (_List__MinerExpirationExtend__Repr) AsString() (string, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.AsString()
}
func (_List__MinerExpirationExtend__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.AsBytes()
}
func (_List__MinerExpirationExtend__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.List__MinerExpirationExtend.Repr"}.AsLink()
}
func (_List__MinerExpirationExtend__Repr) Prototype() ipld.NodePrototype {
	return _List__MinerExpirationExtend__ReprPrototype{}
}
type _List__MinerExpirationExtend__ReprPrototype struct{}

func (_List__MinerExpirationExtend__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _List__MinerExpirationExtend__ReprBuilder
	nb.Reset()
	return &nb
}
type _List__MinerExpirationExtend__ReprBuilder struct {
	_List__MinerExpirationExtend__ReprAssembler
}
func (nb *_List__MinerExpirationExtend__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_List__MinerExpirationExtend__ReprBuilder) Reset() {
	var w _List__MinerExpirationExtend
	var m schema.Maybe
	*nb = _List__MinerExpirationExtend__ReprBuilder{_List__MinerExpirationExtend__ReprAssembler{w: &w, m: &m}}
}
type _List__MinerExpirationExtend__ReprAssembler struct {
	w *_List__MinerExpirationExtend
	m *schema.Maybe
	state laState

	cm schema.Maybe
	va _MinerExpirationExtend__ReprAssembler
}

func (na *_List__MinerExpirationExtend__ReprAssembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_List__MinerExpirationExtend__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.BeginMap(0)
}
func (na *_List__MinerExpirationExtend__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if sizeHint < 0 {
		sizeHint = 0
	}
	if na.w == nil {
		na.w = &_List__MinerExpirationExtend{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_MinerExpirationExtend, 0, sizeHint)
	}
	return na, nil
}
func (na *_List__MinerExpirationExtend__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_List__MinerExpirationExtend__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.AssignBool(false)
}
func (_List__MinerExpirationExtend__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.AssignInt(0)
}
func (_List__MinerExpirationExtend__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.AssignFloat(0)
}
func (_List__MinerExpirationExtend__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.AssignString("")
}
func (_List__MinerExpirationExtend__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.AssignBytes(nil)
}
func (_List__MinerExpirationExtend__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.List__MinerExpirationExtend.Repr"}.AssignLink(nil)
}
func (na *_List__MinerExpirationExtend__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_List__MinerExpirationExtend); ok {
		switch *na.m {
		case schema.Maybe_Value, schema.Maybe_Null:
			panic("invalid state: cannot assign into assembler that's already finished")
		case midvalue:
			panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
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
	if v.ReprKind() != ipld.ReprKind_List {
		return ipld.ErrWrongKind{TypeName: "types.List__MinerExpirationExtend.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
	}
	itr := v.ListIterator()
	for !itr.Done() {
		_, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_List__MinerExpirationExtend__ReprAssembler) Prototype() ipld.NodePrototype {
	return _List__MinerExpirationExtend__ReprPrototype{}
}
func (la *_List__MinerExpirationExtend__ReprAssembler) valueFinishTidy() bool {
	switch la.cm {
	case schema.Maybe_Value:
		la.va.w = nil
		la.cm = schema.Maybe_Absent
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_List__MinerExpirationExtend__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when still in the middle of assembling the previous value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	la.w.x = append(la.w.x, _MinerExpirationExtend{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_List__MinerExpirationExtend__ReprAssembler) Finish() error {
	switch la.state {
	case laState_initial:
		// carry on
	case laState_midValue:
		if !la.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case laState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	la.state = laState_finished
	*la.m = schema.Maybe_Value
	return nil
}
func (la *_List__MinerExpirationExtend__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _MinerExpirationExtend__ReprPrototype{}
}