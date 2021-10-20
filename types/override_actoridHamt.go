package types

import (
	hamt "github.com/filecoin-project/go-hamt-ipld/v2"
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _Map__ActorID struct {
	m map[_RawAddress]*_ActorID
	t []_Map__ActorID__entry
}
type Map__ActorID = *_Map__ActorID
type _Map__ActorID__entry struct {
	k _RawAddress
	v _ActorID
}

func (n *_Map__ActorID) LookupMaybe(k RawAddress) MaybeActorID {
	v, ok := n.m[*k]
	if !ok {
		return &_Map__ActorID__valueAbsent
	}
	return &_ActorID__Maybe{
		m: schema.Maybe_Value,
		v: *v,
	}
}

var _Map__ActorID__valueAbsent = _ActorID__Maybe{m: schema.Maybe_Absent}

// TODO generate also a plain Lookup method that doesn't box and alloc if this type contains non-nullable values!
type _Map__ActorID__Maybe struct {
	m schema.Maybe
	v Map__ActorID
}
type MaybeMap__ActorID = *_Map__ActorID__Maybe

func (m MaybeMap__ActorID) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMap__ActorID) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMap__ActorID) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMap__ActorID) AsNode() ipld.Node {
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
func (m MaybeMap__ActorID) Must() Map__ActorID {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}

var _ ipld.Node = (Map__ActorID)(&_Map__ActorID{})
var _ schema.TypedNode = (Map__ActorID)(&_Map__ActorID{})

func (Map__ActorID) Kind() ipld.Kind {
	return ipld.Kind_Map
}
func (n Map__ActorID) LookupByString(k string) (ipld.Node, error) {
	var k2 _RawAddress
	if err := (_RawAddress__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	v, exists := n.m[k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k)}
	}
	return v, nil
}
func (n Map__ActorID) LookupByNode(k ipld.Node) (ipld.Node, error) {
	k2, ok := k.(RawAddress)
	if !ok {
		panic("todo invalid key type error")
		// 'ipld.ErrInvalidKey{TypeName:"types.Map__ActorID", Key:&_String{k}}' doesn't quite cut it: need room to explain the type, and it's not guaranteed k can be turned into a string at all
	}
	v, exists := n.m[*k2]
	if !exists {
		return nil, ipld.ErrNotExists{ipld.PathSegmentOfString(k2.String())}
	}
	return v, nil
}
func (Map__ActorID) LookupByIndex(idx int64) (ipld.Node, error) {
	return mixins.Map{"types.Map__ActorID"}.LookupByIndex(0)
}
func (n Map__ActorID) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n Map__ActorID) MapIterator() ipld.MapIterator {
	return &_Map__ActorID__MapItr{n, 0}
}

type _Map__ActorID__MapItr struct {
	n   Map__ActorID
	idx int
}

func (itr *_Map__ActorID__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= len(itr.n.t) {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	x := &itr.n.t[itr.idx]
	k = &x.k
	v = &x.v
	itr.idx++
	return
}
func (itr *_Map__ActorID__MapItr) Done() bool {
	return itr.idx >= len(itr.n.t)
}

func (Map__ActorID) ListIterator() ipld.ListIterator {
	return nil
}
func (n Map__ActorID) Length() int64 {
	return int64(len(n.t))
}
func (Map__ActorID) IsAbsent() bool {
	return false
}
func (Map__ActorID) IsNull() bool {
	return false
}
func (Map__ActorID) AsBool() (bool, error) {
	return mixins.Map{"types.Map__ActorID"}.AsBool()
}
func (Map__ActorID) AsInt() (int64, error) {
	return mixins.Map{"types.Map__ActorID"}.AsInt()
}
func (Map__ActorID) AsFloat() (float64, error) {
	return mixins.Map{"types.Map__ActorID"}.AsFloat()
}
func (Map__ActorID) AsString() (string, error) {
	return mixins.Map{"types.Map__ActorID"}.AsString()
}
func (Map__ActorID) AsBytes() ([]byte, error) {
	return mixins.Map{"types.Map__ActorID"}.AsBytes()
}
func (Map__ActorID) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.Map__ActorID"}.AsLink()
}
func (Map__ActorID) Prototype() ipld.NodePrototype {
	return _Map__ActorID__Prototype{}
}

type _Map__ActorID__Prototype struct{}

func (_Map__ActorID__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _Map__ActorID__Builder
	nb.Reset()
	return &nb
}

type _Map__ActorID__Builder struct {
	_Map__ActorID__Assembler
}

func (nb *_Map__ActorID__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	//HACK: willscott
	if nb.h != nil {
		return nb.h
	}
	return nb.w
}
func (nb *_Map__ActorID__Builder) Reset() {
	var w _Map__ActorID
	var m schema.Maybe
	*nb = _Map__ActorID__Builder{_Map__ActorID__Assembler{w: &w, m: &m}}
}

type _Map__ActorID__Assembler struct {
	w     *_Map__ActorID
	m     *schema.Maybe
	h     *hamt.Node
	state maState

	cm schema.Maybe
	ka _RawAddress__Assembler
	va _ActorID__Assembler
}

func (na *_Map__ActorID__Assembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__ActorID__Assembler) BeginMap(sizeHint int64) (ipld.MapAssembler, error) {
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
		na.w = &_Map__ActorID{}
	}
	na.w.m = make(map[_RawAddress]*_ActorID, sizeHint)
	na.w.t = make([]_Map__ActorID__entry, 0, sizeHint)
	return na, nil
}
func (_Map__ActorID__Assembler) BeginList(sizeHint int64) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.Map__ActorID"}.BeginList(0)
}
func (na *_Map__ActorID__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.Map__ActorID"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__ActorID__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.Map__ActorID"}.AssignBool(false)
}
func (_Map__ActorID__Assembler) AssignInt(int64) error {
	return mixins.MapAssembler{"types.Map__ActorID"}.AssignInt(0)
}
func (_Map__ActorID__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.Map__ActorID"}.AssignFloat(0)
}
func (_Map__ActorID__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.Map__ActorID"}.AssignString("")
}
func (_Map__ActorID__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.Map__ActorID"}.AssignBytes(nil)
}
func (_Map__ActorID__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.Map__ActorID"}.AssignLink(nil)
}
func (na *_Map__ActorID__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map__ActorID); ok {
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
		/// HACK: willscott
	} else if v2, ok := v.(*hamt.Node); ok {
		na.h = v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.Kind() != ipld.Kind_Map {
		return ipld.ErrWrongKind{TypeName: "types.Map__ActorID", MethodName: "AssignNode", AppropriateKind: ipld.KindSet_JustMap, ActualKind: v.Kind()}
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
func (_Map__ActorID__Assembler) Prototype() ipld.NodePrototype {
	return _Map__ActorID__Prototype{}
}
func (ma *_Map__ActorID__Assembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.w = &tz.v
		ma.va.m = &ma.cm
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__ActorID__Assembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.va.w = nil
		ma.cm = schema.Maybe_Absent
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__ActorID__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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

	var k2 _RawAddress
	if err := (_RawAddress__Prototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__ActorID__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.w = &tz.v
	ma.va.m = &ma.cm
	return &ma.va, nil
}
func (ma *_Map__ActorID__Assembler) AssembleKey() ipld.NodeAssembler {
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
	ma.w.t = append(ma.w.t, _Map__ActorID__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__ActorID__Assembler) AssembleValue() ipld.NodeAssembler {
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
func (ma *_Map__ActorID__Assembler) Finish() error {
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
func (ma *_Map__ActorID__Assembler) KeyPrototype() ipld.NodePrototype {
	return _RawAddress__Prototype{}
}
func (ma *_Map__ActorID__Assembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _ActorID__Prototype{}
}
func (Map__ActorID) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n Map__ActorID) Representation() ipld.Node {
	return (*_Map__ActorID__Repr)(n)
}

type _Map__ActorID__Repr _Map__ActorID

var _ ipld.Node = &_Map__ActorID__Repr{}

func (_Map__ActorID__Repr) Kind() ipld.Kind {
	return ipld.Kind_Map
}
func (nr *_Map__ActorID__Repr) LookupByString(k string) (ipld.Node, error) {
	v, err := (Map__ActorID)(nr).LookupByString(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(ActorID).Representation(), nil
}
func (nr *_Map__ActorID__Repr) LookupByNode(k ipld.Node) (ipld.Node, error) {
	v, err := (Map__ActorID)(nr).LookupByNode(k)
	if err != nil || v == ipld.Null {
		return v, err
	}
	return v.(ActorID).Representation(), nil
}
func (_Map__ActorID__Repr) LookupByIndex(idx int64) (ipld.Node, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.LookupByIndex(0)
}
func (n _Map__ActorID__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (nr *_Map__ActorID__Repr) MapIterator() ipld.MapIterator {
	return &_Map__ActorID__ReprMapItr{(Map__ActorID)(nr), 0}
}

type _Map__ActorID__ReprMapItr _Map__ActorID__MapItr

func (itr *_Map__ActorID__ReprMapItr) Next() (k ipld.Node, v ipld.Node, err error) {
	k, v, err = (*_Map__ActorID__MapItr)(itr).Next()
	if err != nil || v == ipld.Null {
		return
	}
	return k, v.(ActorID).Representation(), nil
}
func (itr *_Map__ActorID__ReprMapItr) Done() bool {
	return (*_Map__ActorID__MapItr)(itr).Done()
}

func (_Map__ActorID__Repr) ListIterator() ipld.ListIterator {
	return nil
}
func (rn *_Map__ActorID__Repr) Length() int64 {
	return int64(len(rn.t))
}
func (_Map__ActorID__Repr) IsAbsent() bool {
	return false
}
func (_Map__ActorID__Repr) IsNull() bool {
	return false
}
func (_Map__ActorID__Repr) AsBool() (bool, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.AsBool()
}
func (_Map__ActorID__Repr) AsInt() (int64, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.AsInt()
}
func (_Map__ActorID__Repr) AsFloat() (float64, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.AsFloat()
}
func (_Map__ActorID__Repr) AsString() (string, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.AsString()
}
func (_Map__ActorID__Repr) AsBytes() ([]byte, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.AsBytes()
}
func (_Map__ActorID__Repr) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.Map__ActorID.Repr"}.AsLink()
}
func (_Map__ActorID__Repr) Prototype() ipld.NodePrototype {
	return _Map__ActorID__ReprPrototype{}
}

type _Map__ActorID__ReprPrototype struct{}

func (_Map__ActorID__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _Map__ActorID__ReprBuilder
	nb.Reset()
	return &nb
}

type _Map__ActorID__ReprBuilder struct {
	_Map__ActorID__ReprAssembler
}

func (nb *_Map__ActorID__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	// Hack: willscott
	if nb.h != nil {
		return nb.h
	}
	return nb.w
}
func (nb *_Map__ActorID__ReprBuilder) Reset() {
	var w _Map__ActorID
	var m schema.Maybe
	*nb = _Map__ActorID__ReprBuilder{_Map__ActorID__ReprAssembler{w: &w, m: &m}}
}

type _Map__ActorID__ReprAssembler struct {
	w     *_Map__ActorID
	m     *schema.Maybe
	h     *hamt.Node
	state maState

	cm schema.Maybe
	ka _RawAddress__ReprAssembler
	va _ActorID__ReprAssembler
}

func (na *_Map__ActorID__ReprAssembler) reset() {
	na.state = maState_initial
	na.ka.reset()
	na.va.reset()
}
func (na *_Map__ActorID__ReprAssembler) BeginMap(sizeHint int64) (ipld.MapAssembler, error) {
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
		na.w = &_Map__ActorID{}
	}
	na.w.m = make(map[_RawAddress]*_ActorID, sizeHint)
	na.w.t = make([]_Map__ActorID__entry, 0, sizeHint)
	return na, nil
}
func (_Map__ActorID__ReprAssembler) BeginList(sizeHint int64) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.BeginList(0)
}
func (na *_Map__ActorID__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.Map__ActorID.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_Map__ActorID__ReprAssembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.AssignBool(false)
}
func (_Map__ActorID__ReprAssembler) AssignInt(int64) error {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.AssignInt(0)
}
func (_Map__ActorID__ReprAssembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.AssignFloat(0)
}
func (_Map__ActorID__ReprAssembler) AssignString(string) error {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.AssignString("")
}
func (_Map__ActorID__ReprAssembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.AssignBytes(nil)
}
func (_Map__ActorID__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.Map__ActorID.Repr"}.AssignLink(nil)
}
func (na *_Map__ActorID__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_Map__ActorID); ok {
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
		/// HACK: willscott
	} else if v2, ok := v.(*hamt.Node); ok {
		na.h = v2
		*na.m = schema.Maybe_Value
		return nil
	}
	if v.Kind() != ipld.Kind_Map {
		return ipld.ErrWrongKind{TypeName: "types.Map__ActorID.Repr", MethodName: "AssignNode", AppropriateKind: ipld.KindSet_JustMap, ActualKind: v.Kind()}
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
func (_Map__ActorID__ReprAssembler) Prototype() ipld.NodePrototype {
	return _Map__ActorID__ReprPrototype{}
}
func (ma *_Map__ActorID__ReprAssembler) keyFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.ka.w = nil
		tz := &ma.w.t[len(ma.w.t)-1]
		ma.cm = schema.Maybe_Absent
		ma.state = maState_expectValue
		ma.w.m[tz.k] = &tz.v
		ma.va.w = &tz.v
		ma.va.m = &ma.cm
		ma.ka.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__ActorID__ReprAssembler) valueFinishTidy() bool {
	switch ma.cm {
	case schema.Maybe_Value:
		ma.va.w = nil
		ma.cm = schema.Maybe_Absent
		ma.state = maState_initial
		ma.va.reset()
		return true
	default:
		return false
	}
}
func (ma *_Map__ActorID__ReprAssembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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

	var k2 _RawAddress
	if err := (_RawAddress__ReprPrototype{}).fromString(&k2, k); err != nil {
		return nil, err // TODO wrap in some kind of ErrInvalidKey
	}
	if _, exists := ma.w.m[k2]; exists {
		return nil, ipld.ErrRepeatedMapKey{&k2}
	}
	ma.w.t = append(ma.w.t, _Map__ActorID__entry{k: k2})
	tz := &ma.w.t[len(ma.w.t)-1]
	ma.state = maState_midValue

	ma.w.m[k2] = &tz.v
	ma.va.w = &tz.v
	ma.va.m = &ma.cm
	return &ma.va, nil
}
func (ma *_Map__ActorID__ReprAssembler) AssembleKey() ipld.NodeAssembler {
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
	ma.w.t = append(ma.w.t, _Map__ActorID__entry{})
	ma.state = maState_midKey
	ma.ka.m = &ma.cm
	ma.ka.w = &ma.w.t[len(ma.w.t)-1].k
	return &ma.ka
}
func (ma *_Map__ActorID__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
func (ma *_Map__ActorID__ReprAssembler) Finish() error {
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
func (ma *_Map__ActorID__ReprAssembler) KeyPrototype() ipld.NodePrototype {
	return _RawAddress__ReprPrototype{}
}
func (ma *_Map__ActorID__ReprAssembler) ValuePrototype(_ string) ipld.NodePrototype {
	return _ActorID__ReprPrototype{}
}
