package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MinerV0PowerPair struct {
	Raw _BigInt
	QA _BigInt
}
type MinerV0PowerPair = *_MinerV0PowerPair

func (n _MinerV0PowerPair) FieldRaw()	BigInt {
	return &n.Raw
}
func (n _MinerV0PowerPair) FieldQA()	BigInt {
	return &n.QA
}
type _MinerV0PowerPair__Maybe struct {
	m schema.Maybe
	v MinerV0PowerPair
}
type MaybeMinerV0PowerPair = *_MinerV0PowerPair__Maybe

func (m MaybeMinerV0PowerPair) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMinerV0PowerPair) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMinerV0PowerPair) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMinerV0PowerPair) AsNode() ipld.Node {
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
func (m MaybeMinerV0PowerPair) Must() MinerV0PowerPair {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MinerV0PowerPair_Raw = _String{"Raw"}
	fieldName__MinerV0PowerPair_QA = _String{"QA"}
)
var _ ipld.Node = (MinerV0PowerPair)(&_MinerV0PowerPair{})
var _ schema.TypedNode = (MinerV0PowerPair)(&_MinerV0PowerPair{})
func (MinerV0PowerPair) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MinerV0PowerPair) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "Raw":
		return &n.Raw, nil
	case "QA":
		return &n.QA, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n MinerV0PowerPair) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MinerV0PowerPair) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.LookupByIndex(0)
}
func (n MinerV0PowerPair) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MinerV0PowerPair) MapIterator() ipld.MapIterator {
	return &_MinerV0PowerPair__MapItr{n, 0}
}

type _MinerV0PowerPair__MapItr struct {
	n MinerV0PowerPair
	idx  int
}

func (itr *_MinerV0PowerPair__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MinerV0PowerPair_Raw
		v = &itr.n.Raw
	case 1:
		k = &fieldName__MinerV0PowerPair_QA
		v = &itr.n.QA
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MinerV0PowerPair__MapItr) Done() bool {
	return itr.idx >= 2
}

func (MinerV0PowerPair) ListIterator() ipld.ListIterator {
	return nil
}
func (MinerV0PowerPair) Length() int {
	return 2
}
func (MinerV0PowerPair) IsAbsent() bool {
	return false
}
func (MinerV0PowerPair) IsNull() bool {
	return false
}
func (MinerV0PowerPair) AsBool() (bool, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.AsBool()
}
func (MinerV0PowerPair) AsInt() (int, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.AsInt()
}
func (MinerV0PowerPair) AsFloat() (float64, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.AsFloat()
}
func (MinerV0PowerPair) AsString() (string, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.AsString()
}
func (MinerV0PowerPair) AsBytes() ([]byte, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.AsBytes()
}
func (MinerV0PowerPair) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.MinerV0PowerPair"}.AsLink()
}
func (MinerV0PowerPair) Prototype() ipld.NodePrototype {
	return _MinerV0PowerPair__Prototype{}
}
type _MinerV0PowerPair__Prototype struct{}

func (_MinerV0PowerPair__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MinerV0PowerPair__Builder
	nb.Reset()
	return &nb
}
type _MinerV0PowerPair__Builder struct {
	_MinerV0PowerPair__Assembler
}
func (nb *_MinerV0PowerPair__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MinerV0PowerPair__Builder) Reset() {
	var w _MinerV0PowerPair
	var m schema.Maybe
	*nb = _MinerV0PowerPair__Builder{_MinerV0PowerPair__Assembler{w: &w, m: &m}}
}
type _MinerV0PowerPair__Assembler struct {
	w *_MinerV0PowerPair
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_Raw _BigInt__Assembler
	ca_QA _BigInt__Assembler
	}

func (na *_MinerV0PowerPair__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_Raw.reset()
	na.ca_QA.reset()
}

var (
	fieldBit__MinerV0PowerPair_Raw = 1 << 0
	fieldBit__MinerV0PowerPair_QA = 1 << 1
	fieldBits__MinerV0PowerPair_sufficient = 0 + 1 << 0 + 1 << 1
)
func (na *_MinerV0PowerPair__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MinerV0PowerPair{}
	}
	return na, nil
}
func (_MinerV0PowerPair__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.BeginList(0)
}
func (na *_MinerV0PowerPair__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MinerV0PowerPair__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignBool(false)
}
func (_MinerV0PowerPair__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignInt(0)
}
func (_MinerV0PowerPair__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignFloat(0)
}
func (_MinerV0PowerPair__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignString("")
}
func (_MinerV0PowerPair__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignBytes(nil)
}
func (_MinerV0PowerPair__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.MinerV0PowerPair"}.AssignLink(nil)
}
func (na *_MinerV0PowerPair__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MinerV0PowerPair); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MinerV0PowerPair", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MinerV0PowerPair__Assembler) Prototype() ipld.NodePrototype {
	return _MinerV0PowerPair__Prototype{}
}
func (ma *_MinerV0PowerPair__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_Raw.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_QA.w = nil
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
func (ma *_MinerV0PowerPair__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "Raw":
		if ma.s & fieldBit__MinerV0PowerPair_Raw != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MinerV0PowerPair_Raw}
		}
		ma.s += fieldBit__MinerV0PowerPair_Raw
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_Raw.w = &ma.w.Raw
		ma.ca_Raw.m = &ma.cm
		return &ma.ca_Raw, nil
	case "QA":
		if ma.s & fieldBit__MinerV0PowerPair_QA != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MinerV0PowerPair_QA}
		}
		ma.s += fieldBit__MinerV0PowerPair_QA
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_QA.w = &ma.w.QA
		ma.ca_QA.m = &ma.cm
		return &ma.ca_QA, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.MinerV0PowerPair", Key:&_String{k}}
	}
}
func (ma *_MinerV0PowerPair__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MinerV0PowerPair__KeyAssembler)(ma)
}
func (ma *_MinerV0PowerPair__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_Raw.w = &ma.w.Raw
		ma.ca_Raw.m = &ma.cm
		return &ma.ca_Raw
	case 1:
		ma.ca_QA.w = &ma.w.QA
		ma.ca_QA.m = &ma.cm
		return &ma.ca_QA
	default:
		panic("unreachable")
	}
}
func (ma *_MinerV0PowerPair__Assembler) Finish() error {
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
func (ma *_MinerV0PowerPair__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MinerV0PowerPair__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MinerV0PowerPair__KeyAssembler _MinerV0PowerPair__Assembler
func (_MinerV0PowerPair__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.BeginMap(0)
}
func (_MinerV0PowerPair__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.BeginList(0)
}
func (na *_MinerV0PowerPair__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.AssignNull()
}
func (_MinerV0PowerPair__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.AssignBool(false)
}
func (_MinerV0PowerPair__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.AssignInt(0)
}
func (_MinerV0PowerPair__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MinerV0PowerPair__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "Raw":
		if ka.s & fieldBit__MinerV0PowerPair_Raw != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MinerV0PowerPair_Raw}
		}
		ka.s += fieldBit__MinerV0PowerPair_Raw
		ka.state = maState_expectValue
		ka.f = 0
	case "QA":
		if ka.s & fieldBit__MinerV0PowerPair_QA != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MinerV0PowerPair_QA}
		}
		ka.s += fieldBit__MinerV0PowerPair_QA
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"types.MinerV0PowerPair", Key:&_String{k}}
	}
	return nil
}
func (_MinerV0PowerPair__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.AssignBytes(nil)
}
func (_MinerV0PowerPair__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.MinerV0PowerPair.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MinerV0PowerPair__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MinerV0PowerPair__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MinerV0PowerPair) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MinerV0PowerPair) Representation() ipld.Node {
	return (*_MinerV0PowerPair__Repr)(n)
}
type _MinerV0PowerPair__Repr _MinerV0PowerPair
var _ ipld.Node = &_MinerV0PowerPair__Repr{}
func (_MinerV0PowerPair__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_MinerV0PowerPair__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.LookupByString("")
}
func (n *_MinerV0PowerPair__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_MinerV0PowerPair__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.Raw.Representation(), nil
	case 1:
		return n.QA.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _MinerV0PowerPair__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.MinerV0PowerPair.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_MinerV0PowerPair__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_MinerV0PowerPair__Repr) ListIterator() ipld.ListIterator {
	return &_MinerV0PowerPair__ReprListItr{n, 0}
}

type _MinerV0PowerPair__ReprListItr struct {
	n   *_MinerV0PowerPair__Repr
	idx int
	
}

func (itr *_MinerV0PowerPair__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 2 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.Raw.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.QA.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MinerV0PowerPair__ReprListItr) Done() bool {
	return itr.idx >= 2
}

func (rn *_MinerV0PowerPair__Repr) Length() int {
	l := 2
	return l
}
func (_MinerV0PowerPair__Repr) IsAbsent() bool {
	return false
}
func (_MinerV0PowerPair__Repr) IsNull() bool {
	return false
}
func (_MinerV0PowerPair__Repr) AsBool() (bool, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.AsBool()
}
func (_MinerV0PowerPair__Repr) AsInt() (int, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.AsInt()
}
func (_MinerV0PowerPair__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.AsFloat()
}
func (_MinerV0PowerPair__Repr) AsString() (string, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.AsString()
}
func (_MinerV0PowerPair__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.AsBytes()
}
func (_MinerV0PowerPair__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.MinerV0PowerPair.Repr"}.AsLink()
}
func (_MinerV0PowerPair__Repr) Prototype() ipld.NodePrototype {
	return _MinerV0PowerPair__ReprPrototype{}
}
type _MinerV0PowerPair__ReprPrototype struct{}

func (_MinerV0PowerPair__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MinerV0PowerPair__ReprBuilder
	nb.Reset()
	return &nb
}
type _MinerV0PowerPair__ReprBuilder struct {
	_MinerV0PowerPair__ReprAssembler
}
func (nb *_MinerV0PowerPair__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MinerV0PowerPair__ReprBuilder) Reset() {
	var w _MinerV0PowerPair
	var m schema.Maybe
	*nb = _MinerV0PowerPair__ReprBuilder{_MinerV0PowerPair__ReprAssembler{w: &w, m: &m}}
}
type _MinerV0PowerPair__ReprAssembler struct {
	w *_MinerV0PowerPair
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_Raw _BigInt__ReprAssembler
	ca_QA _BigInt__ReprAssembler
	}

func (na *_MinerV0PowerPair__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_Raw.reset()
	na.ca_QA.reset()
}
func (_MinerV0PowerPair__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.BeginMap(0)
}
func (na *_MinerV0PowerPair__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MinerV0PowerPair{}
	}
	return na, nil
}
func (na *_MinerV0PowerPair__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.MinerV0PowerPair.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MinerV0PowerPair__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.AssignBool(false)
}
func (_MinerV0PowerPair__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.AssignInt(0)
}
func (_MinerV0PowerPair__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.AssignFloat(0)
}
func (_MinerV0PowerPair__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.AssignString("")
}
func (_MinerV0PowerPair__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.AssignBytes(nil)
}
func (_MinerV0PowerPair__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.MinerV0PowerPair.Repr"}.AssignLink(nil)
}
func (na *_MinerV0PowerPair__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MinerV0PowerPair); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MinerV0PowerPair.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_MinerV0PowerPair__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MinerV0PowerPair__ReprPrototype{}
}
func (la *_MinerV0PowerPair__ReprAssembler) valueFinishTidy() bool {
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
func (la *_MinerV0PowerPair__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
		la.ca_Raw.w = &la.w.Raw
		la.ca_Raw.m = &la.cm
		return &la.ca_Raw
	case 1:
		la.ca_QA.w = &la.w.QA
		la.ca_QA.m = &la.cm
		return &la.ca_QA
	default:
		panic("unreachable")
	}
}
func (la *_MinerV0PowerPair__ReprAssembler) Finish() error {
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
func (la *_MinerV0PowerPair__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}