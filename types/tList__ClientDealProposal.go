package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _List__ClientDealProposal struct {
	x []_MarketClientDealProposal
}
type List__ClientDealProposal = *_List__ClientDealProposal
type _List__ClientDealProposal__Maybe struct {
	m schema.Maybe
	v List__ClientDealProposal
}
type MaybeList__ClientDealProposal = *_List__ClientDealProposal__Maybe

func (m MaybeList__ClientDealProposal) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeList__ClientDealProposal) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeList__ClientDealProposal) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeList__ClientDealProposal) AsNode() ipld.Node {
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
func (m MaybeList__ClientDealProposal) Must() List__ClientDealProposal {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (List__ClientDealProposal)(&_List__ClientDealProposal{})
var _ schema.TypedNode = (List__ClientDealProposal)(&_List__ClientDealProposal{})
func (List__ClientDealProposal) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (List__ClientDealProposal) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.List__ClientDealProposal"}.LookupByString("")
}
func (n List__ClientDealProposal) LookupByNode(k ipld.Node) (ipld.Node, error) {
	idx, err := k.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(idx)
}
func (n List__ClientDealProposal) LookupByIndex(idx int) (ipld.Node, error) {
	if n.Length() <= idx {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(idx)}
	}
	v := &n.x[idx]
	return v, nil
}
func (n List__ClientDealProposal) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.List__ClientDealProposal", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (List__ClientDealProposal) MapIterator() ipld.MapIterator {
	return nil
}
func (n List__ClientDealProposal) ListIterator() ipld.ListIterator {
	return &_List__ClientDealProposal__ListItr{n, 0}
}

type _List__ClientDealProposal__ListItr struct {
	n List__ClientDealProposal
	idx  int
}

func (itr *_List__ClientDealProposal__ListItr) Next() (idx int, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.x) {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	x := &itr.n.x[itr.idx]
	v = x
	itr.idx++
	return
}
func (itr *_List__ClientDealProposal__ListItr) Done() bool {
	return itr.idx >= len(itr.n.x)
}

func (n List__ClientDealProposal) Length() int {
	return len(n.x)
}
func (List__ClientDealProposal) IsAbsent() bool {
	return false
}
func (List__ClientDealProposal) IsNull() bool {
	return false
}
func (List__ClientDealProposal) AsBool() (bool, error) {
	return mixins.List{"types.List__ClientDealProposal"}.AsBool()
}
func (List__ClientDealProposal) AsInt() (int, error) {
	return mixins.List{"types.List__ClientDealProposal"}.AsInt()
}
func (List__ClientDealProposal) AsFloat() (float64, error) {
	return mixins.List{"types.List__ClientDealProposal"}.AsFloat()
}
func (List__ClientDealProposal) AsString() (string, error) {
	return mixins.List{"types.List__ClientDealProposal"}.AsString()
}
func (List__ClientDealProposal) AsBytes() ([]byte, error) {
	return mixins.List{"types.List__ClientDealProposal"}.AsBytes()
}
func (List__ClientDealProposal) AsLink() (ipld.Link, error) {
	return mixins.List{"types.List__ClientDealProposal"}.AsLink()
}
func (List__ClientDealProposal) Prototype() ipld.NodePrototype {
	return _List__ClientDealProposal__Prototype{}
}
type _List__ClientDealProposal__Prototype struct{}

func (_List__ClientDealProposal__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _List__ClientDealProposal__Builder
	nb.Reset()
	return &nb
}
type _List__ClientDealProposal__Builder struct {
	_List__ClientDealProposal__Assembler
}
func (nb *_List__ClientDealProposal__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_List__ClientDealProposal__Builder) Reset() {
	var w _List__ClientDealProposal
	var m schema.Maybe
	*nb = _List__ClientDealProposal__Builder{_List__ClientDealProposal__Assembler{w: &w, m: &m}}
}
type _List__ClientDealProposal__Assembler struct {
	w *_List__ClientDealProposal
	m *schema.Maybe
	state laState

	cm schema.Maybe
	va _MarketClientDealProposal__Assembler
}

func (na *_List__ClientDealProposal__Assembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_List__ClientDealProposal__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.BeginMap(0)
}
func (na *_List__ClientDealProposal__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
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
		na.w = &_List__ClientDealProposal{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_MarketClientDealProposal, 0, sizeHint)
	}
	return na, nil
}
func (na *_List__ClientDealProposal__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_List__ClientDealProposal__Assembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignBool(false)
}
func (_List__ClientDealProposal__Assembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignInt(0)
}
func (_List__ClientDealProposal__Assembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignFloat(0)
}
func (_List__ClientDealProposal__Assembler) AssignString(string) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignString("")
}
func (_List__ClientDealProposal__Assembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignBytes(nil)
}
func (_List__ClientDealProposal__Assembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal"}.AssignLink(nil)
}
func (na *_List__ClientDealProposal__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_List__ClientDealProposal); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.List__ClientDealProposal", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_List__ClientDealProposal__Assembler) Prototype() ipld.NodePrototype {
	return _List__ClientDealProposal__Prototype{}
}
func (la *_List__ClientDealProposal__Assembler) valueFinishTidy() bool {
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
func (la *_List__ClientDealProposal__Assembler) AssembleValue() ipld.NodeAssembler {
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
	la.w.x = append(la.w.x, _MarketClientDealProposal{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_List__ClientDealProposal__Assembler) Finish() error {
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
func (la *_List__ClientDealProposal__Assembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _MarketClientDealProposal__Prototype{}
}
func (List__ClientDealProposal) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n List__ClientDealProposal) Representation() ipld.Node {
	return (*_List__ClientDealProposal__Repr)(n)
}
type _List__ClientDealProposal__Repr _List__ClientDealProposal
var _ ipld.Node = &_List__ClientDealProposal__Repr{}
func (_List__ClientDealProposal__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_List__ClientDealProposal__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.LookupByString("")
}
func (nr *_List__ClientDealProposal__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (List__ClientDealProposal)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(MarketClientDealProposal).Representation(), nil
}
func (nr *_List__ClientDealProposal__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	v, err := (List__ClientDealProposal)(nr).LookupByIndex(idx)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(MarketClientDealProposal).Representation(), nil
}
func (n _List__ClientDealProposal__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.List__ClientDealProposal.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_List__ClientDealProposal__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (nr *_List__ClientDealProposal__Repr) ListIterator() ipld.ListIterator {
	return &_List__ClientDealProposal__ReprListItr{(List__ClientDealProposal)(nr), 0}
}

type _List__ClientDealProposal__ReprListItr _List__ClientDealProposal__ListItr

func (itr *_List__ClientDealProposal__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	idx, v, err = (*_List__ClientDealProposal__ListItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return idx, v.(MarketClientDealProposal).Representation(), nil
}
func (itr *_List__ClientDealProposal__ReprListItr) Done() bool {
	return (*_List__ClientDealProposal__ListItr)(itr).Done()
}

func (rn *_List__ClientDealProposal__Repr) Length() int {
	return len(rn.x)
}
func (_List__ClientDealProposal__Repr) IsAbsent() bool {
	return false
}
func (_List__ClientDealProposal__Repr) IsNull() bool {
	return false
}
func (_List__ClientDealProposal__Repr) AsBool() (bool, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.AsBool()
}
func (_List__ClientDealProposal__Repr) AsInt() (int, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.AsInt()
}
func (_List__ClientDealProposal__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.AsFloat()
}
func (_List__ClientDealProposal__Repr) AsString() (string, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.AsString()
}
func (_List__ClientDealProposal__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.AsBytes()
}
func (_List__ClientDealProposal__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.List__ClientDealProposal.Repr"}.AsLink()
}
func (_List__ClientDealProposal__Repr) Prototype() ipld.NodePrototype {
	return _List__ClientDealProposal__ReprPrototype{}
}
type _List__ClientDealProposal__ReprPrototype struct{}

func (_List__ClientDealProposal__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _List__ClientDealProposal__ReprBuilder
	nb.Reset()
	return &nb
}
type _List__ClientDealProposal__ReprBuilder struct {
	_List__ClientDealProposal__ReprAssembler
}
func (nb *_List__ClientDealProposal__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_List__ClientDealProposal__ReprBuilder) Reset() {
	var w _List__ClientDealProposal
	var m schema.Maybe
	*nb = _List__ClientDealProposal__ReprBuilder{_List__ClientDealProposal__ReprAssembler{w: &w, m: &m}}
}
type _List__ClientDealProposal__ReprAssembler struct {
	w *_List__ClientDealProposal
	m *schema.Maybe
	state laState

	cm schema.Maybe
	va _MarketClientDealProposal__ReprAssembler
}

func (na *_List__ClientDealProposal__ReprAssembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_List__ClientDealProposal__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.BeginMap(0)
}
func (na *_List__ClientDealProposal__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
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
		na.w = &_List__ClientDealProposal{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_MarketClientDealProposal, 0, sizeHint)
	}
	return na, nil
}
func (na *_List__ClientDealProposal__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.List__ClientDealProposal.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_List__ClientDealProposal__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.AssignBool(false)
}
func (_List__ClientDealProposal__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.AssignInt(0)
}
func (_List__ClientDealProposal__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.AssignFloat(0)
}
func (_List__ClientDealProposal__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.AssignString("")
}
func (_List__ClientDealProposal__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.AssignBytes(nil)
}
func (_List__ClientDealProposal__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.List__ClientDealProposal.Repr"}.AssignLink(nil)
}
func (na *_List__ClientDealProposal__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_List__ClientDealProposal); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.List__ClientDealProposal.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_List__ClientDealProposal__ReprAssembler) Prototype() ipld.NodePrototype {
	return _List__ClientDealProposal__ReprPrototype{}
}
func (la *_List__ClientDealProposal__ReprAssembler) valueFinishTidy() bool {
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
func (la *_List__ClientDealProposal__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	la.w.x = append(la.w.x, _MarketClientDealProposal{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.w = row
	la.va.m = &la.cm
	return &la.va
}
func (la *_List__ClientDealProposal__ReprAssembler) Finish() error {
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
func (la *_List__ClientDealProposal__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _MarketClientDealProposal__ReprPrototype{}
}