package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _PowerV2Claim struct {
	SealProofType _Int
	RawBytePower _BigInt
	QualityAdjPower _BigInt
}
type PowerV2Claim = *_PowerV2Claim

func (n _PowerV2Claim) FieldSealProofType()	Int {
	return &n.SealProofType
}
func (n _PowerV2Claim) FieldRawBytePower()	BigInt {
	return &n.RawBytePower
}
func (n _PowerV2Claim) FieldQualityAdjPower()	BigInt {
	return &n.QualityAdjPower
}
type _PowerV2Claim__Maybe struct {
	m schema.Maybe
	v PowerV2Claim
}
type MaybePowerV2Claim = *_PowerV2Claim__Maybe

func (m MaybePowerV2Claim) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybePowerV2Claim) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybePowerV2Claim) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybePowerV2Claim) AsNode() ipld.Node {
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
func (m MaybePowerV2Claim) Must() PowerV2Claim {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__PowerV2Claim_SealProofType = _String{"SealProofType"}
	fieldName__PowerV2Claim_RawBytePower = _String{"RawBytePower"}
	fieldName__PowerV2Claim_QualityAdjPower = _String{"QualityAdjPower"}
)
var _ ipld.Node = (PowerV2Claim)(&_PowerV2Claim{})
var _ schema.TypedNode = (PowerV2Claim)(&_PowerV2Claim{})
func (PowerV2Claim) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n PowerV2Claim) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "SealProofType":
		return &n.SealProofType, nil
	case "RawBytePower":
		return &n.RawBytePower, nil
	case "QualityAdjPower":
		return &n.QualityAdjPower, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n PowerV2Claim) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (PowerV2Claim) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.PowerV2Claim"}.LookupByIndex(0)
}
func (n PowerV2Claim) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n PowerV2Claim) MapIterator() ipld.MapIterator {
	return &_PowerV2Claim__MapItr{n, 0}
}

type _PowerV2Claim__MapItr struct {
	n PowerV2Claim
	idx  int
}

func (itr *_PowerV2Claim__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__PowerV2Claim_SealProofType
		v = &itr.n.SealProofType
	case 1:
		k = &fieldName__PowerV2Claim_RawBytePower
		v = &itr.n.RawBytePower
	case 2:
		k = &fieldName__PowerV2Claim_QualityAdjPower
		v = &itr.n.QualityAdjPower
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_PowerV2Claim__MapItr) Done() bool {
	return itr.idx >= 3
}

func (PowerV2Claim) ListIterator() ipld.ListIterator {
	return nil
}
func (PowerV2Claim) Length() int {
	return 3
}
func (PowerV2Claim) IsAbsent() bool {
	return false
}
func (PowerV2Claim) IsNull() bool {
	return false
}
func (PowerV2Claim) AsBool() (bool, error) {
	return mixins.Map{"types.PowerV2Claim"}.AsBool()
}
func (PowerV2Claim) AsInt() (int, error) {
	return mixins.Map{"types.PowerV2Claim"}.AsInt()
}
func (PowerV2Claim) AsFloat() (float64, error) {
	return mixins.Map{"types.PowerV2Claim"}.AsFloat()
}
func (PowerV2Claim) AsString() (string, error) {
	return mixins.Map{"types.PowerV2Claim"}.AsString()
}
func (PowerV2Claim) AsBytes() ([]byte, error) {
	return mixins.Map{"types.PowerV2Claim"}.AsBytes()
}
func (PowerV2Claim) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.PowerV2Claim"}.AsLink()
}
func (PowerV2Claim) Prototype() ipld.NodePrototype {
	return _PowerV2Claim__Prototype{}
}
type _PowerV2Claim__Prototype struct{}

func (_PowerV2Claim__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _PowerV2Claim__Builder
	nb.Reset()
	return &nb
}
type _PowerV2Claim__Builder struct {
	_PowerV2Claim__Assembler
}
func (nb *_PowerV2Claim__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_PowerV2Claim__Builder) Reset() {
	var w _PowerV2Claim
	var m schema.Maybe
	*nb = _PowerV2Claim__Builder{_PowerV2Claim__Assembler{w: &w, m: &m}}
}
type _PowerV2Claim__Assembler struct {
	w *_PowerV2Claim
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_SealProofType _Int__Assembler
	ca_RawBytePower _BigInt__Assembler
	ca_QualityAdjPower _BigInt__Assembler
	}

func (na *_PowerV2Claim__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_SealProofType.reset()
	na.ca_RawBytePower.reset()
	na.ca_QualityAdjPower.reset()
}

var (
	fieldBit__PowerV2Claim_SealProofType = 1 << 0
	fieldBit__PowerV2Claim_RawBytePower = 1 << 1
	fieldBit__PowerV2Claim_QualityAdjPower = 1 << 2
	fieldBits__PowerV2Claim_sufficient = 0 + 1 << 0 + 1 << 1 + 1 << 2
)
func (na *_PowerV2Claim__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_PowerV2Claim{}
	}
	return na, nil
}
func (_PowerV2Claim__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.PowerV2Claim"}.BeginList(0)
}
func (na *_PowerV2Claim__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.PowerV2Claim"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_PowerV2Claim__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.PowerV2Claim"}.AssignBool(false)
}
func (_PowerV2Claim__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.PowerV2Claim"}.AssignInt(0)
}
func (_PowerV2Claim__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.PowerV2Claim"}.AssignFloat(0)
}
func (_PowerV2Claim__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.PowerV2Claim"}.AssignString("")
}
func (_PowerV2Claim__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.PowerV2Claim"}.AssignBytes(nil)
}
func (_PowerV2Claim__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.PowerV2Claim"}.AssignLink(nil)
}
func (na *_PowerV2Claim__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_PowerV2Claim); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.PowerV2Claim", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_PowerV2Claim__Assembler) Prototype() ipld.NodePrototype {
	return _PowerV2Claim__Prototype{}
}
func (ma *_PowerV2Claim__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_SealProofType.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_RawBytePower.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 2:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_QualityAdjPower.w = nil
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
func (ma *_PowerV2Claim__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "SealProofType":
		if ma.s & fieldBit__PowerV2Claim_SealProofType != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__PowerV2Claim_SealProofType}
		}
		ma.s += fieldBit__PowerV2Claim_SealProofType
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_SealProofType.w = &ma.w.SealProofType
		ma.ca_SealProofType.m = &ma.cm
		return &ma.ca_SealProofType, nil
	case "RawBytePower":
		if ma.s & fieldBit__PowerV2Claim_RawBytePower != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__PowerV2Claim_RawBytePower}
		}
		ma.s += fieldBit__PowerV2Claim_RawBytePower
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_RawBytePower.w = &ma.w.RawBytePower
		ma.ca_RawBytePower.m = &ma.cm
		return &ma.ca_RawBytePower, nil
	case "QualityAdjPower":
		if ma.s & fieldBit__PowerV2Claim_QualityAdjPower != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__PowerV2Claim_QualityAdjPower}
		}
		ma.s += fieldBit__PowerV2Claim_QualityAdjPower
		ma.state = maState_midValue
		ma.f = 2
		ma.ca_QualityAdjPower.w = &ma.w.QualityAdjPower
		ma.ca_QualityAdjPower.m = &ma.cm
		return &ma.ca_QualityAdjPower, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.PowerV2Claim", Key:&_String{k}}
	}
}
func (ma *_PowerV2Claim__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_PowerV2Claim__KeyAssembler)(ma)
}
func (ma *_PowerV2Claim__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_SealProofType.w = &ma.w.SealProofType
		ma.ca_SealProofType.m = &ma.cm
		return &ma.ca_SealProofType
	case 1:
		ma.ca_RawBytePower.w = &ma.w.RawBytePower
		ma.ca_RawBytePower.m = &ma.cm
		return &ma.ca_RawBytePower
	case 2:
		ma.ca_QualityAdjPower.w = &ma.w.QualityAdjPower
		ma.ca_QualityAdjPower.m = &ma.cm
		return &ma.ca_QualityAdjPower
	default:
		panic("unreachable")
	}
}
func (ma *_PowerV2Claim__Assembler) Finish() error {
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
func (ma *_PowerV2Claim__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_PowerV2Claim__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _PowerV2Claim__KeyAssembler _PowerV2Claim__Assembler
func (_PowerV2Claim__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.BeginMap(0)
}
func (_PowerV2Claim__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.BeginList(0)
}
func (na *_PowerV2Claim__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.AssignNull()
}
func (_PowerV2Claim__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.AssignBool(false)
}
func (_PowerV2Claim__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.AssignInt(0)
}
func (_PowerV2Claim__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.AssignFloat(0)
}
func (ka *_PowerV2Claim__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "SealProofType":
		if ka.s & fieldBit__PowerV2Claim_SealProofType != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__PowerV2Claim_SealProofType}
		}
		ka.s += fieldBit__PowerV2Claim_SealProofType
		ka.state = maState_expectValue
		ka.f = 0
	case "RawBytePower":
		if ka.s & fieldBit__PowerV2Claim_RawBytePower != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__PowerV2Claim_RawBytePower}
		}
		ka.s += fieldBit__PowerV2Claim_RawBytePower
		ka.state = maState_expectValue
		ka.f = 1
	case "QualityAdjPower":
		if ka.s & fieldBit__PowerV2Claim_QualityAdjPower != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__PowerV2Claim_QualityAdjPower}
		}
		ka.s += fieldBit__PowerV2Claim_QualityAdjPower
		ka.state = maState_expectValue
		ka.f = 2
	default:
		return ipld.ErrInvalidKey{TypeName:"types.PowerV2Claim", Key:&_String{k}}
	}
	return nil
}
func (_PowerV2Claim__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.AssignBytes(nil)
}
func (_PowerV2Claim__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.PowerV2Claim.KeyAssembler"}.AssignLink(nil)
}
func (ka *_PowerV2Claim__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_PowerV2Claim__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (PowerV2Claim) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n PowerV2Claim) Representation() ipld.Node {
	return (*_PowerV2Claim__Repr)(n)
}
type _PowerV2Claim__Repr _PowerV2Claim
var _ ipld.Node = &_PowerV2Claim__Repr{}
func (_PowerV2Claim__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_PowerV2Claim__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.LookupByString("")
}
func (n *_PowerV2Claim__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_PowerV2Claim__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.SealProofType.Representation(), nil
	case 1:
		return n.RawBytePower.Representation(), nil
	case 2:
		return n.QualityAdjPower.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _PowerV2Claim__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.PowerV2Claim.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_PowerV2Claim__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_PowerV2Claim__Repr) ListIterator() ipld.ListIterator {
	return &_PowerV2Claim__ReprListItr{n, 0}
}

type _PowerV2Claim__ReprListItr struct {
	n   *_PowerV2Claim__Repr
	idx int
	
}

func (itr *_PowerV2Claim__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 3 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.SealProofType.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.RawBytePower.Representation()
	case 2:
		idx = itr.idx
		v = itr.n.QualityAdjPower.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_PowerV2Claim__ReprListItr) Done() bool {
	return itr.idx >= 3
}

func (rn *_PowerV2Claim__Repr) Length() int {
	l := 3
	return l
}
func (_PowerV2Claim__Repr) IsAbsent() bool {
	return false
}
func (_PowerV2Claim__Repr) IsNull() bool {
	return false
}
func (_PowerV2Claim__Repr) AsBool() (bool, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.AsBool()
}
func (_PowerV2Claim__Repr) AsInt() (int, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.AsInt()
}
func (_PowerV2Claim__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.AsFloat()
}
func (_PowerV2Claim__Repr) AsString() (string, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.AsString()
}
func (_PowerV2Claim__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.AsBytes()
}
func (_PowerV2Claim__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.PowerV2Claim.Repr"}.AsLink()
}
func (_PowerV2Claim__Repr) Prototype() ipld.NodePrototype {
	return _PowerV2Claim__ReprPrototype{}
}
type _PowerV2Claim__ReprPrototype struct{}

func (_PowerV2Claim__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _PowerV2Claim__ReprBuilder
	nb.Reset()
	return &nb
}
type _PowerV2Claim__ReprBuilder struct {
	_PowerV2Claim__ReprAssembler
}
func (nb *_PowerV2Claim__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_PowerV2Claim__ReprBuilder) Reset() {
	var w _PowerV2Claim
	var m schema.Maybe
	*nb = _PowerV2Claim__ReprBuilder{_PowerV2Claim__ReprAssembler{w: &w, m: &m}}
}
type _PowerV2Claim__ReprAssembler struct {
	w *_PowerV2Claim
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_SealProofType _Int__ReprAssembler
	ca_RawBytePower _BigInt__ReprAssembler
	ca_QualityAdjPower _BigInt__ReprAssembler
	}

func (na *_PowerV2Claim__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_SealProofType.reset()
	na.ca_RawBytePower.reset()
	na.ca_QualityAdjPower.reset()
}
func (_PowerV2Claim__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.BeginMap(0)
}
func (na *_PowerV2Claim__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_PowerV2Claim{}
	}
	return na, nil
}
func (na *_PowerV2Claim__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.PowerV2Claim.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_PowerV2Claim__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.AssignBool(false)
}
func (_PowerV2Claim__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.AssignInt(0)
}
func (_PowerV2Claim__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.AssignFloat(0)
}
func (_PowerV2Claim__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.AssignString("")
}
func (_PowerV2Claim__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.AssignBytes(nil)
}
func (_PowerV2Claim__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.PowerV2Claim.Repr"}.AssignLink(nil)
}
func (na *_PowerV2Claim__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_PowerV2Claim); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.PowerV2Claim.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_PowerV2Claim__ReprAssembler) Prototype() ipld.NodePrototype {
	return _PowerV2Claim__ReprPrototype{}
}
func (la *_PowerV2Claim__ReprAssembler) valueFinishTidy() bool {
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
	case 2:
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
func (la *_PowerV2Claim__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	if la.f >= 3 {
		return _ErrorThunkAssembler{schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(3)}}
	}
	la.state = laState_midValue
	switch la.f {
	case 0:
		la.ca_SealProofType.w = &la.w.SealProofType
		la.ca_SealProofType.m = &la.cm
		return &la.ca_SealProofType
	case 1:
		la.ca_RawBytePower.w = &la.w.RawBytePower
		la.ca_RawBytePower.m = &la.cm
		return &la.ca_RawBytePower
	case 2:
		la.ca_QualityAdjPower.w = &la.w.QualityAdjPower
		la.ca_QualityAdjPower.m = &la.cm
		return &la.ca_QualityAdjPower
	default:
		panic("unreachable")
	}
}
func (la *_PowerV2Claim__ReprAssembler) Finish() error {
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
func (la *_PowerV2Claim__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}