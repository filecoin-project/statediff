package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Multimap__PowerV0CronEvent struct {
	m map[_String]MaybeMap__PowerV0CronEvent
	t []_Multimap__PowerV0CronEvent__entry
}
type Multimap__PowerV0CronEvent = *_Multimap__PowerV0CronEvent
type _Multimap__PowerV0CronEvent__entry struct {
	k _String
	v _Map__PowerV0CronEvent__Maybe
}
func (n *_Multimap__PowerV0CronEvent) LookupMaybe(k String) MaybeMap__PowerV0CronEvent {
	v, ok := n.m[*k]
	if !ok {
		return &_Multimap__PowerV0CronEvent__valueAbsent
	}
	return v
}

var _Multimap__PowerV0CronEvent__valueAbsent = _Map__PowerV0CronEvent__Maybe{m:schema.Maybe_Absent}

// TODO generate also a plain Lookup method that doesn't box and alloc if this type contains non-nullable values!
type _Multimap__PowerV0CronEvent__Maybe struct {
	m schema.Maybe
	v Multimap__PowerV0CronEvent
}
type MaybeMultimap__PowerV0CronEvent = *_Multimap__PowerV0CronEvent__Maybe

func (m MaybeMultimap__PowerV0CronEvent) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMultimap__PowerV0CronEvent) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMultimap__PowerV0CronEvent) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMultimap__PowerV0CronEvent) AsNode() ipld.Node {
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
func (m MaybeMultimap__PowerV0CronEvent) Must() Multimap__PowerV0CronEvent {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var _ ipld.Node = (Multimap__PowerV0CronEvent)(&_Multimap__PowerV0CronEvent{})
var _ schema.TypedNode = (Multimap__PowerV0CronEvent)(&_Multimap__PowerV0CronEvent{})
func (Multimap__PowerV0CronEvent) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n Multimap__PowerV0CronEvent) LookupByString(k string) (ipld.Node, error) {
	var k2 _String
	if err := (_String__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	v, exists := n.m[k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k)}
	}
	if v.m == schema.Maybe_Null {
		return ipld.Null, nil
	}
	return v.v, nil
}
func (n Multimap__PowerV0CronEvent) LookupByNode(k ipld.Node) (ipld.Node, error) {
	k2, ok := k.(String)
	if !ok {
		panic("todo invalid key type error")
		// 'ipld.ErrInvalidKey{TypeName:"types.Multimap__PowerV0CronEvent", Key:&_String{k}}' doesn't quite cut it: need room to explain the type, and it's not guaranteed k can be turned into a string at all
	}
	v, exists := n.m[*k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k2.String())}
	}
	if v.m == schema.Maybe_Null {
		return ipld.Null, nil
	}
	return v.v, nil
}
func (Multimap__PowerV0CronEvent) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.LookupByIndex(0)
}
func (n Multimap__PowerV0CronEvent) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n Multimap__PowerV0CronEvent) MapIterator() ipld.MapIterator {
	return &_Multimap__PowerV0CronEvent__MapItr{n, 0}
}

type _Multimap__PowerV0CronEvent__MapItr struct {
	n Multimap__PowerV0CronEvent
	idx  int
}

func (itr *_Multimap__PowerV0CronEvent__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.t) {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	x := &itr.n.t[itr.idx]
	k = &x.k
	switch x.v.m {
	case schema.Maybe_Null:
		v = ipld.Null
	case schema.Maybe_Value:
		v = x.v.v
	}
	itr.idx++
	return
}
func (itr *_Multimap__PowerV0CronEvent__MapItr) Done() bool {
	return itr.idx >= len(itr.n.t)
}

func (Multimap__PowerV0CronEvent) ListIterator() ipld.ListIterator {
	return nil
}
func (n Multimap__PowerV0CronEvent) Length() int {
	return len(n.t)
}
func (Multimap__PowerV0CronEvent) IsAbsent() bool {
	return false
}
func (Multimap__PowerV0CronEvent) IsNull() bool {
	return false
}
func (Multimap__PowerV0CronEvent) AsBool() (bool, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.AsBool()
}
func (Multimap__PowerV0CronEvent) AsInt() (int, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.AsInt()
}
func (Multimap__PowerV0CronEvent) AsFloat() (float64, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.AsFloat()
}
func (Multimap__PowerV0CronEvent) AsString() (string, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.AsString()
}
func (Multimap__PowerV0CronEvent) AsBytes() ([]byte, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.AsBytes()
}
func (Multimap__PowerV0CronEvent) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent"}.AsLink()
}
func (Multimap__PowerV0CronEvent) Prototype() ipld.NodePrototype {
	return _Multimap__PowerV0CronEvent__Prototype{}
}
type _Multimap__PowerV0CronEvent__Prototype struct{}

func (_Multimap__PowerV0CronEvent__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Multimap__PowerV0CronEvent__Builder
	nb.Reset()
	return &nb
}
type _Multimap__PowerV0CronEvent__Builder struct {
	_Multimap__PowerV0CronEvent__Assembler
}
func (nb *_Multimap__PowerV0CronEvent__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Multimap__PowerV0CronEvent__Builder) Reset() {
	var w _Multimap__PowerV0CronEvent
	var m schema.Maybe
	*nb = _Multimap__PowerV0CronEvent__Builder{_Multimap__PowerV0CronEvent__Assembler{w: &w, m: &m}}
}
type _Multimap__PowerV0CronEvent__Assembler struct {
	w *_Multimap__PowerV0CronEvent
	m *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__Assembler
	va _Map__PowerV0CronEvent__Assembler
}

func (na *_Multimap__PowerV0CronEvent__Assembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Multimap__PowerV0CronEvent__Assembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
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
		na.w = &_Multimap__PowerV0CronEvent{}
	}
	na.w.m = make(map[_String]MaybeMap__PowerV0CronEvent, sizeHint)
	na.w.t = make([]_Multimap__PowerV0CronEvent__entry, 0, sizeHint)
	return na, nil
}
func (_Multimap__PowerV0CronEvent__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.BeginList(0)
}
func (na *_Multimap__PowerV0CronEvent__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Multimap__PowerV0CronEvent__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignBool(false)
}
func (_Multimap__PowerV0CronEvent__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignInt(0)
}
func (_Multimap__PowerV0CronEvent__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignFloat(0)
}
func (_Multimap__PowerV0CronEvent__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignString("")
}
func (_Multimap__PowerV0CronEvent__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignBytes(nil)
}
func (_Multimap__PowerV0CronEvent__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent"}.AssignLink(nil)
}
func (na *_Multimap__PowerV0CronEvent__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Multimap__PowerV0CronEvent); ok {
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
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "types.Multimap__PowerV0CronEvent", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Multimap__PowerV0CronEvent__Assembler) Prototype() ipld.NodePrototype {
	return _Multimap__PowerV0CronEvent__Prototype{}
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.m = &tz.v.m
		tz.v.m = allowNull
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) valueFinishTidy() bool {
	tz := &ma.w.t[len(ma.w.t)-1]
	switch tz.v.m {
	case schema.Maybe_Null:
		ma.state = maState_initial
		ma.va.reset()
		return true
	case schema.Maybe_Value:
		tz.v.v = ma.va.w
		ma.va.w = nil
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}

	var k2 _String
	if err := (_String__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Multimap__PowerV0CronEvent__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.m = &tz.v.m
	tz.v.m = allowNull
	return &ma.va, nil
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.w.t = append(ma.w.t, _Multimap__PowerV0CronEvent__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_Multimap__PowerV0CronEvent__Assembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _Map__PowerV0CronEvent__Prototype{}
}
func (Multimap__PowerV0CronEvent) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Multimap__PowerV0CronEvent) Representation() ipld.Node {
	return (*_Multimap__PowerV0CronEvent__Repr)(n)
}
type _Multimap__PowerV0CronEvent__Repr _Multimap__PowerV0CronEvent
var _ ipld.Node = &_Multimap__PowerV0CronEvent__Repr{}
func (_Multimap__PowerV0CronEvent__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (nr *_Multimap__PowerV0CronEvent__Repr) LookupByString(k string) (ipld.Node, error) {
	v, err := (Multimap__PowerV0CronEvent)(nr).LookupByString(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Map__PowerV0CronEvent).Representation(), nil
}
func (nr *_Multimap__PowerV0CronEvent__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (Multimap__PowerV0CronEvent)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(Map__PowerV0CronEvent).Representation(), nil
}
func (_Multimap__PowerV0CronEvent__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.LookupByIndex(0)
}
func (n _Multimap__PowerV0CronEvent__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (nr *_Multimap__PowerV0CronEvent__Repr) MapIterator() ipld.MapIterator {
	return &_Multimap__PowerV0CronEvent__ReprMapItr{(Multimap__PowerV0CronEvent)(nr), 0}
}

type _Multimap__PowerV0CronEvent__ReprMapItr _Multimap__PowerV0CronEvent__MapItr

func (itr *_Multimap__PowerV0CronEvent__ReprMapItr) Next() (k ipld.Node, v ipld.Node, err error) {
	k, v, err = (*_Multimap__PowerV0CronEvent__MapItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return k, v.(Map__PowerV0CronEvent).Representation(), nil
}
func (itr *_Multimap__PowerV0CronEvent__ReprMapItr) Done() bool {
	return (*_Multimap__PowerV0CronEvent__MapItr)(itr).Done()
}

func (_Multimap__PowerV0CronEvent__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_Multimap__PowerV0CronEvent__Repr) Length() int {
	return len(rn.t)
}
func (_Multimap__PowerV0CronEvent__Repr) IsAbsent() bool {
	return false
}
func (_Multimap__PowerV0CronEvent__Repr) IsNull() bool {
	return false
}
func (_Multimap__PowerV0CronEvent__Repr) AsBool() (bool, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.AsBool()
}
func (_Multimap__PowerV0CronEvent__Repr) AsInt() (int, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.AsInt()
}
func (_Multimap__PowerV0CronEvent__Repr) AsFloat() (float64, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.AsFloat()
}
func (_Multimap__PowerV0CronEvent__Repr) AsString() (string, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.AsString()
}
func (_Multimap__PowerV0CronEvent__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.AsBytes()
}
func (_Multimap__PowerV0CronEvent__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.Multimap__PowerV0CronEvent.Repr"}.AsLink()
}
func (_Multimap__PowerV0CronEvent__Repr) Prototype() ipld.NodePrototype {
	return _Multimap__PowerV0CronEvent__ReprPrototype{}
}
type _Multimap__PowerV0CronEvent__ReprPrototype struct{}

func (_Multimap__PowerV0CronEvent__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Multimap__PowerV0CronEvent__ReprBuilder
	nb.Reset()
	return &nb
}
type _Multimap__PowerV0CronEvent__ReprBuilder struct {
	_Multimap__PowerV0CronEvent__ReprAssembler
}
func (nb *_Multimap__PowerV0CronEvent__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_Multimap__PowerV0CronEvent__ReprBuilder) Reset() {
	var w _Multimap__PowerV0CronEvent
	var m schema.Maybe
	*nb = _Multimap__PowerV0CronEvent__ReprBuilder{_Multimap__PowerV0CronEvent__ReprAssembler{w: &w, m: &m}}
}
type _Multimap__PowerV0CronEvent__ReprAssembler struct {
	w *_Multimap__PowerV0CronEvent
	m *schema.Maybe
	state maState

	cm schema.Maybe
	ka _String__ReprAssembler
	va _Map__PowerV0CronEvent__ReprAssembler
}

func (na *_Multimap__PowerV0CronEvent__ReprAssembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Multimap__PowerV0CronEvent__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
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
		na.w = &_Multimap__PowerV0CronEvent{}
	}
	na.w.m = make(map[_String]MaybeMap__PowerV0CronEvent, sizeHint)
	na.w.t = make([]_Multimap__PowerV0CronEvent__entry, 0, sizeHint)
	return na, nil
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.BeginList(0)
}
func (na *_Multimap__PowerV0CronEvent__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.AssignBool(false)
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.AssignInt(0)
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.AssignFloat(0)
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.AssignString("")
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.AssignBytes(nil)
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.Multimap__PowerV0CronEvent.Repr"}.AssignLink(nil)
}
func (na *_Multimap__PowerV0CronEvent__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Multimap__PowerV0CronEvent); ok {
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
	if v.ReprKind() != ipld.ReprKind_Map {
		return ipld.ErrWrongKind{TypeName: "types.Multimap__PowerV0CronEvent.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
	}
	itr := v.MapIterator()
	for !itr.Done() {
		k, v, err := itr.Next()
		if err != nil {
			return err
		}
		if err := na.AssembleKey().AssignNode(k); err != nil {
			return err
		}
		if err := na.AssembleValue().AssignNode(v); err != nil {
			return err
		}
	}
	return na.Finish()
}
func (_Multimap__PowerV0CronEvent__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Multimap__PowerV0CronEvent__ReprPrototype{}
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.m = &tz.v.m
		tz.v.m = allowNull
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) valueFinishTidy() bool {
	tz := &ma.w.t[len(ma.w.t)-1]
	switch tz.v.m {
	case schema.Maybe_Null:
		ma.state = maState_initial
		ma.va.reset()
		return true
	case schema.Maybe_Value:
		tz.v.v = ma.va.w
		ma.va.w = nil
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleEntry cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleEntry cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleEntry cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleEntry cannot be called on an assembler that's already finished")
	}

	var k2 _String
	if err := (_String__ReprPrototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Multimap__PowerV0CronEvent__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.m = &tz.v.m
	tz.v.m = allowNull
	return &ma.va, nil
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) AssembleKey() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: AssembleKey cannot be called when in the middle of assembling another key")
	case maState_expectValue:
		panic("invalid state: AssembleKey cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: AssembleKey cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: AssembleKey cannot be called on an assembler that's already finished")
	}
	ma.w.t = append(ma.w.t, _Multimap__PowerV0CronEvent__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		if !ma.keyFinishTidy() {
			panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
		} // if tidy success: carry on
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	return &ma.va
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) Finish() error {
	switch ma.state {
	case maState_initial:
		// carry on
	case maState_midKey:
		panic("invalid state: Finish cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		panic("invalid state: Finish cannot be called when expecting start of value assembly")
	case maState_midValue:
		if !ma.valueFinishTidy() {
			panic("invalid state: Finish cannot be called when in the middle of assembling a value")
		} // if tidy success: carry on
	case maState_finished:
		panic("invalid state: Finish cannot be called on an assembler that's already finished")
	}
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _String__ReprPrototype{}
}
func (ma *_Multimap__PowerV0CronEvent__ReprAssembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _Map__PowerV0CronEvent__ReprPrototype{}
}