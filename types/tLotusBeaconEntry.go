package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _LotusBeaconEntry struct {
	Round _Int
	Data _Bytes
}
type LotusBeaconEntry = *_LotusBeaconEntry

func (n _LotusBeaconEntry) FieldRound()	Int {
	return &n.Round
}
func (n _LotusBeaconEntry) FieldData()	Bytes {
	return &n.Data
}
type _LotusBeaconEntry__Maybe struct {
	m schema.Maybe
	v LotusBeaconEntry
}
type MaybeLotusBeaconEntry = *_LotusBeaconEntry__Maybe

func (m MaybeLotusBeaconEntry) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeLotusBeaconEntry) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeLotusBeaconEntry) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeLotusBeaconEntry) AsNode() ipld.Node {
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
func (m MaybeLotusBeaconEntry) Must() LotusBeaconEntry {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__LotusBeaconEntry_Round = _String{"Round"}
	fieldName__LotusBeaconEntry_Data = _String{"Data"}
)
var _ ipld.Node = (LotusBeaconEntry)(&_LotusBeaconEntry{})
var _ schema.TypedNode = (LotusBeaconEntry)(&_LotusBeaconEntry{})
func (LotusBeaconEntry) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n LotusBeaconEntry) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "Round":
		return &n.Round, nil
	case "Data":
		return &n.Data, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n LotusBeaconEntry) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (LotusBeaconEntry) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.LookupByIndex(0)
}
func (n LotusBeaconEntry) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n LotusBeaconEntry) MapIterator() ipld.MapIterator {
	return &_LotusBeaconEntry__MapItr{n, 0}
}

type _LotusBeaconEntry__MapItr struct {
	n LotusBeaconEntry
	idx  int
}

func (itr *_LotusBeaconEntry__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__LotusBeaconEntry_Round
		v = &itr.n.Round
	case 1:
		k = &fieldName__LotusBeaconEntry_Data
		v = &itr.n.Data
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_LotusBeaconEntry__MapItr) Done() bool {
	return itr.idx >= 2
}

func (LotusBeaconEntry) ListIterator() ipld.ListIterator {
	return nil
}
func (LotusBeaconEntry) Length() int {
	return 2
}
func (LotusBeaconEntry) IsAbsent() bool {
	return false
}
func (LotusBeaconEntry) IsNull() bool {
	return false
}
func (LotusBeaconEntry) AsBool() (bool, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.AsBool()
}
func (LotusBeaconEntry) AsInt() (int, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.AsInt()
}
func (LotusBeaconEntry) AsFloat() (float64, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.AsFloat()
}
func (LotusBeaconEntry) AsString() (string, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.AsString()
}
func (LotusBeaconEntry) AsBytes() ([]byte, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.AsBytes()
}
func (LotusBeaconEntry) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.LotusBeaconEntry"}.AsLink()
}
func (LotusBeaconEntry) Prototype() ipld.NodePrototype {
	return _LotusBeaconEntry__Prototype{}
}
type _LotusBeaconEntry__Prototype struct{}

func (_LotusBeaconEntry__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _LotusBeaconEntry__Builder
	nb.Reset()
	return &nb
}
type _LotusBeaconEntry__Builder struct {
	_LotusBeaconEntry__Assembler
}
func (nb *_LotusBeaconEntry__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_LotusBeaconEntry__Builder) Reset() {
	var w _LotusBeaconEntry
	var m schema.Maybe
	*nb = _LotusBeaconEntry__Builder{_LotusBeaconEntry__Assembler{w: &w, m: &m}}
}
type _LotusBeaconEntry__Assembler struct {
	w *_LotusBeaconEntry
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_Round _Int__Assembler
	ca_Data _Bytes__Assembler
	}

func (na *_LotusBeaconEntry__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_Round.reset()
	na.ca_Data.reset()
}

var (
	fieldBit__LotusBeaconEntry_Round = 1 << 0
	fieldBit__LotusBeaconEntry_Data = 1 << 1
	fieldBits__LotusBeaconEntry_sufficient = 0 + 1 << 0 + 1 << 1
)
func (na *_LotusBeaconEntry__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_LotusBeaconEntry{}
	}
	return na, nil
}
func (_LotusBeaconEntry__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.BeginList(0)
}
func (na *_LotusBeaconEntry__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_LotusBeaconEntry__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignBool(false)
}
func (_LotusBeaconEntry__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignInt(0)
}
func (_LotusBeaconEntry__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignFloat(0)
}
func (_LotusBeaconEntry__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignString("")
}
func (_LotusBeaconEntry__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignBytes(nil)
}
func (_LotusBeaconEntry__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.LotusBeaconEntry"}.AssignLink(nil)
}
func (na *_LotusBeaconEntry__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_LotusBeaconEntry); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.LotusBeaconEntry", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_LotusBeaconEntry__Assembler) Prototype() ipld.NodePrototype {
	return _LotusBeaconEntry__Prototype{}
}
func (ma *_LotusBeaconEntry__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_Round.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_Data.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (ma *_LotusBeaconEntry__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	switch k {
	case "Round":
		if ma.s & fieldBit__LotusBeaconEntry_Round != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__LotusBeaconEntry_Round}
		}
		ma.s += fieldBit__LotusBeaconEntry_Round
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_Round.w = &ma.w.Round
		ma.ca_Round.m = &ma.cm
		return &ma.ca_Round, nil
	case "Data":
		if ma.s & fieldBit__LotusBeaconEntry_Data != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__LotusBeaconEntry_Data}
		}
		ma.s += fieldBit__LotusBeaconEntry_Data
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_Data.w = &ma.w.Data
		ma.ca_Data.m = &ma.cm
		return &ma.ca_Data, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.LotusBeaconEntry", Key:&_String{k}}
	}
}
func (ma *_LotusBeaconEntry__Assembler) AssembleKey() ipld.NodeAssembler {
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
	ma.state = maState_midKey
	return (*_LotusBeaconEntry__KeyAssembler)(ma)
}
func (ma *_LotusBeaconEntry__Assembler) AssembleValue() ipld.NodeAssembler {
	switch ma.state {
	case maState_initial:
		panic("invalid state: AssembleValue cannot be called when no key is primed")
	case maState_midKey:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling a key")
	case maState_expectValue:
		// carry on
	case maState_midValue:
		panic("invalid state: AssembleValue cannot be called when in the middle of assembling another value")
	case maState_finished:
		panic("invalid state: AssembleValue cannot be called on an assembler that's already finished")
	}
	ma.state = maState_midValue
	switch ma.f {
	case 0:
		ma.ca_Round.w = &ma.w.Round
		ma.ca_Round.m = &ma.cm
		return &ma.ca_Round
	case 1:
		ma.ca_Data.w = &ma.w.Data
		ma.ca_Data.m = &ma.cm
		return &ma.ca_Data
	default:
		panic("unreachable")
	}
}
func (ma *_LotusBeaconEntry__Assembler) Finish() error {
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
	//FIXME check if all required fields are set
	ma.state = maState_finished
	*ma.m = schema.Maybe_Value
	return nil
}
func (ma *_LotusBeaconEntry__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_LotusBeaconEntry__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _LotusBeaconEntry__KeyAssembler _LotusBeaconEntry__Assembler
func (_LotusBeaconEntry__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.BeginMap(0)
}
func (_LotusBeaconEntry__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.BeginList(0)
}
func (na *_LotusBeaconEntry__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.AssignNull()
}
func (_LotusBeaconEntry__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.AssignBool(false)
}
func (_LotusBeaconEntry__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.AssignInt(0)
}
func (_LotusBeaconEntry__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.AssignFloat(0)
}
func (ka *_LotusBeaconEntry__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "Round":
		if ka.s & fieldBit__LotusBeaconEntry_Round != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__LotusBeaconEntry_Round}
		}
		ka.s += fieldBit__LotusBeaconEntry_Round
		ka.state = maState_expectValue
		ka.f = 0
	case "Data":
		if ka.s & fieldBit__LotusBeaconEntry_Data != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__LotusBeaconEntry_Data}
		}
		ka.s += fieldBit__LotusBeaconEntry_Data
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"types.LotusBeaconEntry", Key:&_String{k}}
	}
	return nil
}
func (_LotusBeaconEntry__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.AssignBytes(nil)
}
func (_LotusBeaconEntry__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.LotusBeaconEntry.KeyAssembler"}.AssignLink(nil)
}
func (ka *_LotusBeaconEntry__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_LotusBeaconEntry__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (LotusBeaconEntry) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n LotusBeaconEntry) Representation() ipld.Node {
	return (*_LotusBeaconEntry__Repr)(n)
}
type _LotusBeaconEntry__Repr _LotusBeaconEntry
var _ ipld.Node = &_LotusBeaconEntry__Repr{}
func (_LotusBeaconEntry__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_LotusBeaconEntry__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.LookupByString("")
}
func (n *_LotusBeaconEntry__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_LotusBeaconEntry__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.Round.Representation(), nil
	case 1:
		return n.Data.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _LotusBeaconEntry__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.LotusBeaconEntry.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_LotusBeaconEntry__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_LotusBeaconEntry__Repr) ListIterator() ipld.ListIterator {
	return &_LotusBeaconEntry__ReprListItr{n, 0}
}

type _LotusBeaconEntry__ReprListItr struct {
	n   *_LotusBeaconEntry__Repr
	idx int
	
}

func (itr *_LotusBeaconEntry__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 2 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.Round.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.Data.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_LotusBeaconEntry__ReprListItr) Done() bool {
	return itr.idx >= 2
}

func (rn *_LotusBeaconEntry__Repr) Length() int {
	l := 2
	return l
}
func (_LotusBeaconEntry__Repr) IsAbsent() bool {
	return false
}
func (_LotusBeaconEntry__Repr) IsNull() bool {
	return false
}
func (_LotusBeaconEntry__Repr) AsBool() (bool, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.AsBool()
}
func (_LotusBeaconEntry__Repr) AsInt() (int, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.AsInt()
}
func (_LotusBeaconEntry__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.AsFloat()
}
func (_LotusBeaconEntry__Repr) AsString() (string, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.AsString()
}
func (_LotusBeaconEntry__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.AsBytes()
}
func (_LotusBeaconEntry__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.LotusBeaconEntry.Repr"}.AsLink()
}
func (_LotusBeaconEntry__Repr) Prototype() ipld.NodePrototype {
	return _LotusBeaconEntry__ReprPrototype{}
}
type _LotusBeaconEntry__ReprPrototype struct{}

func (_LotusBeaconEntry__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _LotusBeaconEntry__ReprBuilder
	nb.Reset()
	return &nb
}
type _LotusBeaconEntry__ReprBuilder struct {
	_LotusBeaconEntry__ReprAssembler
}
func (nb *_LotusBeaconEntry__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_LotusBeaconEntry__ReprBuilder) Reset() {
	var w _LotusBeaconEntry
	var m schema.Maybe
	*nb = _LotusBeaconEntry__ReprBuilder{_LotusBeaconEntry__ReprAssembler{w: &w, m: &m}}
}
type _LotusBeaconEntry__ReprAssembler struct {
	w *_LotusBeaconEntry
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_Round _Int__ReprAssembler
	ca_Data _Bytes__ReprAssembler
	}

func (na *_LotusBeaconEntry__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_Round.reset()
	na.ca_Data.reset()
}
func (_LotusBeaconEntry__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.BeginMap(0)
}
func (na *_LotusBeaconEntry__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_LotusBeaconEntry{}
	}
	return na, nil
}
func (na *_LotusBeaconEntry__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.LotusBeaconEntry.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_LotusBeaconEntry__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.AssignBool(false)
}
func (_LotusBeaconEntry__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.AssignInt(0)
}
func (_LotusBeaconEntry__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.AssignFloat(0)
}
func (_LotusBeaconEntry__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.AssignString("")
}
func (_LotusBeaconEntry__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.AssignBytes(nil)
}
func (_LotusBeaconEntry__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.LotusBeaconEntry.Repr"}.AssignLink(nil)
}
func (na *_LotusBeaconEntry__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_LotusBeaconEntry); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.LotusBeaconEntry.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_LotusBeaconEntry__ReprAssembler) Prototype() ipld.NodePrototype {
	return _LotusBeaconEntry__ReprPrototype{}
}
func (la *_LotusBeaconEntry__ReprAssembler) valueFinishTidy() bool {
	switch la.f {
	case 0:
		switch la.cm {
		case schema.Maybe_Value:
			la.cm = schema.Maybe_Absent
			la.state = laState_initial
			la.f++
			return true
		default:
			return false
		}
	case 1:
		switch la.cm {
		case schema.Maybe_Value:
			la.cm = schema.Maybe_Absent
			la.state = laState_initial
			la.f++
			return true
		default:
			return false
		}
	default:
		panic("unreachable")
	}
}
func (la *_LotusBeaconEntry__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	if la.f >= 2 {
		return _ErrorThunkAssembler{schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(2)}}
	}
	la.state = laState_midValue
	switch la.f {
	case 0:
		la.ca_Round.w = &la.w.Round
		la.ca_Round.m = &la.cm
		return &la.ca_Round
	case 1:
		la.ca_Data.w = &la.w.Data
		la.ca_Data.m = &la.cm
		return &la.ca_Data
	default:
		panic("unreachable")
	}
}
func (la *_LotusBeaconEntry__ReprAssembler) Finish() error {
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
func (la *_LotusBeaconEntry__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}