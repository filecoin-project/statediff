package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _List__MinerV0DeadlineLink struct {
	x []_Link__MinerV0Deadline__Maybe
}
type List__MinerV0DeadlineLink = *_List__MinerV0DeadlineLink
type _List__MinerV0DeadlineLink__Maybe struct {
	m schema.Maybe
	v List__MinerV0DeadlineLink
}
type MaybeList__MinerV0DeadlineLink = *_List__MinerV0DeadlineLink__Maybe

func (m MaybeList__MinerV0DeadlineLink) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeList__MinerV0DeadlineLink) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeList__MinerV0DeadlineLink) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeList__MinerV0DeadlineLink) AsNode() ipld.Node {
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
func (m MaybeList__MinerV0DeadlineLink) Must() List__MinerV0DeadlineLink {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (List__MinerV0DeadlineLink)(&_List__MinerV0DeadlineLink{})
var _ schema.TypedNode = (List__MinerV0DeadlineLink)(&_List__MinerV0DeadlineLink{})
func (List__MinerV0DeadlineLink) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (List__MinerV0DeadlineLink) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.LookupByString("")
}
func (n List__MinerV0DeadlineLink) LookupByNode(k ipld.Node) (ipld.Node, error) {
	idx, err := k.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(idx)
}
func (n List__MinerV0DeadlineLink) LookupByIndex(idx int) (ipld.Node, error) {
	if n.Length() <= idx {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfInt(idx)}
	}
	v := &n.x[idx]
	if v.m == schema.Maybe_Null {
		return ipld.Null, nil
	}
	return v.v, nil
}
func (n List__MinerV0DeadlineLink) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.List__MinerV0DeadlineLink", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (List__MinerV0DeadlineLink) MapIterator() ipld.MapIterator {
	return nil
}
func (n List__MinerV0DeadlineLink) ListIterator() ipld.ListIterator {
	return &_List__MinerV0DeadlineLink__ListItr{n, 0}
}

type _List__MinerV0DeadlineLink__ListItr struct {
	n List__MinerV0DeadlineLink
	idx  int
}

func (itr *_List__MinerV0DeadlineLink__ListItr) Next() (idx int, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.x) {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	idx = itr.idx
	x := &itr.n.x[itr.idx]
	switch x.m {
	case schema.Maybe_Null:
		v = ipld.Null
	case schema.Maybe_Value:
		v = x.v
	}
	itr.idx++
	return
}
func (itr *_List__MinerV0DeadlineLink__ListItr) Done() bool {
	return itr.idx >= len(itr.n.x)
}

func (n List__MinerV0DeadlineLink) Length() int {
	return len(n.x)
}
func (List__MinerV0DeadlineLink) IsAbsent() bool {
	return false
}
func (List__MinerV0DeadlineLink) IsNull() bool {
	return false
}
func (List__MinerV0DeadlineLink) AsBool() (bool, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.AsBool()
}
func (List__MinerV0DeadlineLink) AsInt() (int, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.AsInt()
}
func (List__MinerV0DeadlineLink) AsFloat() (float64, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.AsFloat()
}
func (List__MinerV0DeadlineLink) AsString() (string, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.AsString()
}
func (List__MinerV0DeadlineLink) AsBytes() ([]byte, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.AsBytes()
}
func (List__MinerV0DeadlineLink) AsLink() (ipld.Link, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink"}.AsLink()
}
func (List__MinerV0DeadlineLink) Prototype() ipld.NodePrototype {
	return _List__MinerV0DeadlineLink__Prototype{}
}
type _List__MinerV0DeadlineLink__Prototype struct{}

func (_List__MinerV0DeadlineLink__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _List__MinerV0DeadlineLink__Builder
	nb.Reset()
	return &nb
}
type _List__MinerV0DeadlineLink__Builder struct {
	_List__MinerV0DeadlineLink__Assembler
}
func (nb *_List__MinerV0DeadlineLink__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_List__MinerV0DeadlineLink__Builder) Reset() {
	var w _List__MinerV0DeadlineLink
	var m schema.Maybe
	*nb = _List__MinerV0DeadlineLink__Builder{_List__MinerV0DeadlineLink__Assembler{w: &w, m: &m}}
}
type _List__MinerV0DeadlineLink__Assembler struct {
	w *_List__MinerV0DeadlineLink
	m *schema.Maybe
	state laState

	
	va _Link__MinerV0Deadline__Assembler
}

func (na *_List__MinerV0DeadlineLink__Assembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_List__MinerV0DeadlineLink__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.BeginMap(0)
}
func (na *_List__MinerV0DeadlineLink__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
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
		na.w = &_List__MinerV0DeadlineLink{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_Link__MinerV0Deadline__Maybe, 0, sizeHint)
	}
	return na, nil
}
func (na *_List__MinerV0DeadlineLink__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_List__MinerV0DeadlineLink__Assembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignBool(false)
}
func (_List__MinerV0DeadlineLink__Assembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignInt(0)
}
func (_List__MinerV0DeadlineLink__Assembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignFloat(0)
}
func (_List__MinerV0DeadlineLink__Assembler) AssignString(string) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignString("")
}
func (_List__MinerV0DeadlineLink__Assembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignBytes(nil)
}
func (_List__MinerV0DeadlineLink__Assembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink"}.AssignLink(nil)
}
func (na *_List__MinerV0DeadlineLink__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_List__MinerV0DeadlineLink); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.List__MinerV0DeadlineLink", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_List__MinerV0DeadlineLink__Assembler) Prototype() ipld.NodePrototype {
	return _List__MinerV0DeadlineLink__Prototype{}
}
func (la *_List__MinerV0DeadlineLink__Assembler) valueFinishTidy() bool {
	row := &la.w.x[len(la.w.x)-1]
	switch row.m {
	case schema.Maybe_Value:
		row.v = la.va.w
		la.va.w = nil
		fallthrough
	case schema.Maybe_Null:
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_List__MinerV0DeadlineLink__Assembler) AssembleValue() ipld.NodeAssembler {
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
	la.w.x = append(la.w.x, _Link__MinerV0Deadline__Maybe{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.m = &row.m
	row.m = allowNull
	return &la.va
}
func (la *_List__MinerV0DeadlineLink__Assembler) Finish() error {
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
func (la *_List__MinerV0DeadlineLink__Assembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _Link__MinerV0Deadline__Prototype{}
}
func (List__MinerV0DeadlineLink) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n List__MinerV0DeadlineLink) Representation() ipld.Node {
	return (*_List__MinerV0DeadlineLink__Repr)(n)
}
type _List__MinerV0DeadlineLink__Repr _List__MinerV0DeadlineLink
var _ ipld.Node = &_List__MinerV0DeadlineLink__Repr{}
func (_List__MinerV0DeadlineLink__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_List__MinerV0DeadlineLink__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.LookupByString("")
}
func (nr *_List__MinerV0DeadlineLink__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (List__MinerV0DeadlineLink)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Link__MinerV0Deadline).Representation(), nil
}
func (nr *_List__MinerV0DeadlineLink__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	v, err := (List__MinerV0DeadlineLink)(nr).LookupByIndex(idx)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Link__MinerV0Deadline).Representation(), nil
}
func (n _List__MinerV0DeadlineLink__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.List__MinerV0DeadlineLink.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_List__MinerV0DeadlineLink__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (nr *_List__MinerV0DeadlineLink__Repr) ListIterator() ipld.ListIterator {
	return &_List__MinerV0DeadlineLink__ReprListItr{(List__MinerV0DeadlineLink)(nr), 0}
}

type _List__MinerV0DeadlineLink__ReprListItr _List__MinerV0DeadlineLink__ListItr

func (itr *_List__MinerV0DeadlineLink__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	idx, v, err = (*_List__MinerV0DeadlineLink__ListItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return idx, v.(Link__MinerV0Deadline).Representation(), nil
}
func (itr *_List__MinerV0DeadlineLink__ReprListItr) Done() bool {
	return (*_List__MinerV0DeadlineLink__ListItr)(itr).Done()
}

func (rn *_List__MinerV0DeadlineLink__Repr) Length() int {
	return len(rn.x)
}
func (_List__MinerV0DeadlineLink__Repr) IsAbsent() bool {
	return false
}
func (_List__MinerV0DeadlineLink__Repr) IsNull() bool {
	return false
}
func (_List__MinerV0DeadlineLink__Repr) AsBool() (bool, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.AsBool()
}
func (_List__MinerV0DeadlineLink__Repr) AsInt() (int, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.AsInt()
}
func (_List__MinerV0DeadlineLink__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.AsFloat()
}
func (_List__MinerV0DeadlineLink__Repr) AsString() (string, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.AsString()
}
func (_List__MinerV0DeadlineLink__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.AsBytes()
}
func (_List__MinerV0DeadlineLink__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.List__MinerV0DeadlineLink.Repr"}.AsLink()
}
func (_List__MinerV0DeadlineLink__Repr) Prototype() ipld.NodePrototype {
	return _List__MinerV0DeadlineLink__ReprPrototype{}
}
type _List__MinerV0DeadlineLink__ReprPrototype struct{}

func (_List__MinerV0DeadlineLink__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _List__MinerV0DeadlineLink__ReprBuilder
	nb.Reset()
	return &nb
}
type _List__MinerV0DeadlineLink__ReprBuilder struct {
	_List__MinerV0DeadlineLink__ReprAssembler
}
func (nb *_List__MinerV0DeadlineLink__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_List__MinerV0DeadlineLink__ReprBuilder) Reset() {
	var w _List__MinerV0DeadlineLink
	var m schema.Maybe
	*nb = _List__MinerV0DeadlineLink__ReprBuilder{_List__MinerV0DeadlineLink__ReprAssembler{w: &w, m: &m}}
}
type _List__MinerV0DeadlineLink__ReprAssembler struct {
	w *_List__MinerV0DeadlineLink
	m *schema.Maybe
	state laState

	
	va _Link__MinerV0Deadline__ReprAssembler
}

func (na *_List__MinerV0DeadlineLink__ReprAssembler) reset() {
	na.state = laState_initial
	na.va.reset()
}
func (_List__MinerV0DeadlineLink__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.BeginMap(0)
}
func (na *_List__MinerV0DeadlineLink__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
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
		na.w = &_List__MinerV0DeadlineLink{}
	}
	if sizeHint > 0 {
		na.w.x = make([]_Link__MinerV0Deadline__Maybe, 0, sizeHint)
	}
	return na, nil
}
func (na *_List__MinerV0DeadlineLink__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_List__MinerV0DeadlineLink__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.AssignBool(false)
}
func (_List__MinerV0DeadlineLink__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.AssignInt(0)
}
func (_List__MinerV0DeadlineLink__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.AssignFloat(0)
}
func (_List__MinerV0DeadlineLink__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.AssignString("")
}
func (_List__MinerV0DeadlineLink__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.AssignBytes(nil)
}
func (_List__MinerV0DeadlineLink__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.List__MinerV0DeadlineLink.Repr"}.AssignLink(nil)
}
func (na *_List__MinerV0DeadlineLink__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_List__MinerV0DeadlineLink); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.List__MinerV0DeadlineLink.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_List__MinerV0DeadlineLink__ReprAssembler) Prototype() ipld.NodePrototype {
	return _List__MinerV0DeadlineLink__ReprPrototype{}
}
func (la *_List__MinerV0DeadlineLink__ReprAssembler) valueFinishTidy() bool {
	row := &la.w.x[len(la.w.x)-1]
	switch row.m {
	case schema.Maybe_Value:
		row.v = la.va.w
		la.va.w = nil
		fallthrough
	case schema.Maybe_Null:
		la.state = laState_initial
		la.va.reset()
		return true
	default:
		return false
	}
}
func (la *_List__MinerV0DeadlineLink__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	la.w.x = append(la.w.x, _Link__MinerV0Deadline__Maybe{})
	la.state = laState_midValue
	row := &la.w.x[len(la.w.x)-1]
	la.va.m = &row.m
	row.m = allowNull
	return &la.va
}
func (la *_List__MinerV0DeadlineLink__ReprAssembler) Finish() error {
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
func (la *_List__MinerV0DeadlineLink__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	return _Link__MinerV0Deadline__ReprPrototype{}
}