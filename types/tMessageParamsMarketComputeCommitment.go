package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MessageParamsMarketComputeCommitment struct {
	DealIDs _List__DealID
	SectorType _RegisteredSealProof
}
type MessageParamsMarketComputeCommitment = *_MessageParamsMarketComputeCommitment

func (n _MessageParamsMarketComputeCommitment) FieldDealIDs()	List__DealID {
	return &n.DealIDs
}
func (n _MessageParamsMarketComputeCommitment) FieldSectorType()	RegisteredSealProof {
	return &n.SectorType
}
type _MessageParamsMarketComputeCommitment__Maybe struct {
	m schema.Maybe
	v MessageParamsMarketComputeCommitment
}
type MaybeMessageParamsMarketComputeCommitment = *_MessageParamsMarketComputeCommitment__Maybe

func (m MaybeMessageParamsMarketComputeCommitment) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMessageParamsMarketComputeCommitment) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMessageParamsMarketComputeCommitment) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMessageParamsMarketComputeCommitment) AsNode() ipld.Node {
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
func (m MaybeMessageParamsMarketComputeCommitment) Must() MessageParamsMarketComputeCommitment {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MessageParamsMarketComputeCommitment_DealIDs = _String{"DealIDs"}
	fieldName__MessageParamsMarketComputeCommitment_SectorType = _String{"SectorType"}
)
var _ ipld.Node = (MessageParamsMarketComputeCommitment)(&_MessageParamsMarketComputeCommitment{})
var _ schema.TypedNode = (MessageParamsMarketComputeCommitment)(&_MessageParamsMarketComputeCommitment{})
func (MessageParamsMarketComputeCommitment) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MessageParamsMarketComputeCommitment) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "DealIDs":
		return &n.DealIDs, nil
	case "SectorType":
		return &n.SectorType, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n MessageParamsMarketComputeCommitment) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MessageParamsMarketComputeCommitment) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.LookupByIndex(0)
}
func (n MessageParamsMarketComputeCommitment) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MessageParamsMarketComputeCommitment) MapIterator() ipld.MapIterator {
	return &_MessageParamsMarketComputeCommitment__MapItr{n, 0}
}

type _MessageParamsMarketComputeCommitment__MapItr struct {
	n MessageParamsMarketComputeCommitment
	idx  int
}

func (itr *_MessageParamsMarketComputeCommitment__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 2 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MessageParamsMarketComputeCommitment_DealIDs
		v = &itr.n.DealIDs
	case 1:
		k = &fieldName__MessageParamsMarketComputeCommitment_SectorType
		v = &itr.n.SectorType
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMarketComputeCommitment__MapItr) Done() bool {
	return itr.idx >= 2
}

func (MessageParamsMarketComputeCommitment) ListIterator() ipld.ListIterator {
	return nil
}
func (MessageParamsMarketComputeCommitment) Length() int {
	return 2
}
func (MessageParamsMarketComputeCommitment) IsAbsent() bool {
	return false
}
func (MessageParamsMarketComputeCommitment) IsNull() bool {
	return false
}
func (MessageParamsMarketComputeCommitment) AsBool() (bool, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.AsBool()
}
func (MessageParamsMarketComputeCommitment) AsInt() (int, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.AsInt()
}
func (MessageParamsMarketComputeCommitment) AsFloat() (float64, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.AsFloat()
}
func (MessageParamsMarketComputeCommitment) AsString() (string, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.AsString()
}
func (MessageParamsMarketComputeCommitment) AsBytes() ([]byte, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.AsBytes()
}
func (MessageParamsMarketComputeCommitment) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.MessageParamsMarketComputeCommitment"}.AsLink()
}
func (MessageParamsMarketComputeCommitment) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketComputeCommitment__Prototype{}
}
type _MessageParamsMarketComputeCommitment__Prototype struct{}

func (_MessageParamsMarketComputeCommitment__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMarketComputeCommitment__Builder
	nb.Reset()
	return &nb
}
type _MessageParamsMarketComputeCommitment__Builder struct {
	_MessageParamsMarketComputeCommitment__Assembler
}
func (nb *_MessageParamsMarketComputeCommitment__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMarketComputeCommitment__Builder) Reset() {
	var w _MessageParamsMarketComputeCommitment
	var m schema.Maybe
	*nb = _MessageParamsMarketComputeCommitment__Builder{_MessageParamsMarketComputeCommitment__Assembler{w: &w, m: &m}}
}
type _MessageParamsMarketComputeCommitment__Assembler struct {
	w *_MessageParamsMarketComputeCommitment
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_DealIDs _List__DealID__Assembler
	ca_SectorType _RegisteredSealProof__Assembler
	}

func (na *_MessageParamsMarketComputeCommitment__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_DealIDs.reset()
	na.ca_SectorType.reset()
}

var (
	fieldBit__MessageParamsMarketComputeCommitment_DealIDs = 1 << 0
	fieldBit__MessageParamsMarketComputeCommitment_SectorType = 1 << 1
	fieldBits__MessageParamsMarketComputeCommitment_sufficient = 0 + 1 << 0 + 1 << 1
)
func (na *_MessageParamsMarketComputeCommitment__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMarketComputeCommitment{}
	}
	return na, nil
}
func (_MessageParamsMarketComputeCommitment__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.BeginList(0)
}
func (na *_MessageParamsMarketComputeCommitment__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMarketComputeCommitment__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignBool(false)
}
func (_MessageParamsMarketComputeCommitment__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignInt(0)
}
func (_MessageParamsMarketComputeCommitment__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignFloat(0)
}
func (_MessageParamsMarketComputeCommitment__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignString("")
}
func (_MessageParamsMarketComputeCommitment__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignBytes(nil)
}
func (_MessageParamsMarketComputeCommitment__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.MessageParamsMarketComputeCommitment"}.AssignLink(nil)
}
func (na *_MessageParamsMarketComputeCommitment__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMarketComputeCommitment); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMarketComputeCommitment", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MessageParamsMarketComputeCommitment__Assembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketComputeCommitment__Prototype{}
}
func (ma *_MessageParamsMarketComputeCommitment__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_DealIDs.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_SectorType.w = nil
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
func (ma *_MessageParamsMarketComputeCommitment__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "DealIDs":
		if ma.s & fieldBit__MessageParamsMarketComputeCommitment_DealIDs != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketComputeCommitment_DealIDs}
		}
		ma.s += fieldBit__MessageParamsMarketComputeCommitment_DealIDs
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_DealIDs.w = &ma.w.DealIDs
		ma.ca_DealIDs.m = &ma.cm
		return &ma.ca_DealIDs, nil
	case "SectorType":
		if ma.s & fieldBit__MessageParamsMarketComputeCommitment_SectorType != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketComputeCommitment_SectorType}
		}
		ma.s += fieldBit__MessageParamsMarketComputeCommitment_SectorType
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_SectorType.w = &ma.w.SectorType
		ma.ca_SectorType.m = &ma.cm
		return &ma.ca_SectorType, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.MessageParamsMarketComputeCommitment", Key:&_String{k}}
	}
}
func (ma *_MessageParamsMarketComputeCommitment__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MessageParamsMarketComputeCommitment__KeyAssembler)(ma)
}
func (ma *_MessageParamsMarketComputeCommitment__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_DealIDs.w = &ma.w.DealIDs
		ma.ca_DealIDs.m = &ma.cm
		return &ma.ca_DealIDs
	case 1:
		ma.ca_SectorType.w = &ma.w.SectorType
		ma.ca_SectorType.m = &ma.cm
		return &ma.ca_SectorType
	default:
		panic("unreachable")
	}
}
func (ma *_MessageParamsMarketComputeCommitment__Assembler) Finish() error {
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
func (ma *_MessageParamsMarketComputeCommitment__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MessageParamsMarketComputeCommitment__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MessageParamsMarketComputeCommitment__KeyAssembler _MessageParamsMarketComputeCommitment__Assembler
func (_MessageParamsMarketComputeCommitment__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.BeginMap(0)
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.BeginList(0)
}
func (na *_MessageParamsMarketComputeCommitment__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.AssignNull()
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.AssignBool(false)
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.AssignInt(0)
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MessageParamsMarketComputeCommitment__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "DealIDs":
		if ka.s & fieldBit__MessageParamsMarketComputeCommitment_DealIDs != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketComputeCommitment_DealIDs}
		}
		ka.s += fieldBit__MessageParamsMarketComputeCommitment_DealIDs
		ka.state = maState_expectValue
		ka.f = 0
	case "SectorType":
		if ka.s & fieldBit__MessageParamsMarketComputeCommitment_SectorType != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketComputeCommitment_SectorType}
		}
		ka.s += fieldBit__MessageParamsMarketComputeCommitment_SectorType
		ka.state = maState_expectValue
		ka.f = 1
	default:
		return ipld.ErrInvalidKey{TypeName:"types.MessageParamsMarketComputeCommitment", Key:&_String{k}}
	}
	return nil
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.AssignBytes(nil)
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.MessageParamsMarketComputeCommitment.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MessageParamsMarketComputeCommitment__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MessageParamsMarketComputeCommitment__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MessageParamsMarketComputeCommitment) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MessageParamsMarketComputeCommitment) Representation() ipld.Node {
	return (*_MessageParamsMarketComputeCommitment__Repr)(n)
}
type _MessageParamsMarketComputeCommitment__Repr _MessageParamsMarketComputeCommitment
var _ ipld.Node = &_MessageParamsMarketComputeCommitment__Repr{}
func (_MessageParamsMarketComputeCommitment__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_MessageParamsMarketComputeCommitment__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.LookupByString("")
}
func (n *_MessageParamsMarketComputeCommitment__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_MessageParamsMarketComputeCommitment__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.DealIDs.Representation(), nil
	case 1:
		return n.SectorType.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _MessageParamsMarketComputeCommitment__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.MessageParamsMarketComputeCommitment.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_MessageParamsMarketComputeCommitment__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_MessageParamsMarketComputeCommitment__Repr) ListIterator() ipld.ListIterator {
	return &_MessageParamsMarketComputeCommitment__ReprListItr{n, 0}
}

type _MessageParamsMarketComputeCommitment__ReprListItr struct {
	n   *_MessageParamsMarketComputeCommitment__Repr
	idx int
	
}

func (itr *_MessageParamsMarketComputeCommitment__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 2 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.DealIDs.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.SectorType.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMarketComputeCommitment__ReprListItr) Done() bool {
	return itr.idx >= 2
}

func (rn *_MessageParamsMarketComputeCommitment__Repr) Length() int {
	l := 2
	return l
}
func (_MessageParamsMarketComputeCommitment__Repr) IsAbsent() bool {
	return false
}
func (_MessageParamsMarketComputeCommitment__Repr) IsNull() bool {
	return false
}
func (_MessageParamsMarketComputeCommitment__Repr) AsBool() (bool, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.AsBool()
}
func (_MessageParamsMarketComputeCommitment__Repr) AsInt() (int, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.AsInt()
}
func (_MessageParamsMarketComputeCommitment__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.AsFloat()
}
func (_MessageParamsMarketComputeCommitment__Repr) AsString() (string, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.AsString()
}
func (_MessageParamsMarketComputeCommitment__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.AsBytes()
}
func (_MessageParamsMarketComputeCommitment__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.MessageParamsMarketComputeCommitment.Repr"}.AsLink()
}
func (_MessageParamsMarketComputeCommitment__Repr) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketComputeCommitment__ReprPrototype{}
}
type _MessageParamsMarketComputeCommitment__ReprPrototype struct{}

func (_MessageParamsMarketComputeCommitment__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMarketComputeCommitment__ReprBuilder
	nb.Reset()
	return &nb
}
type _MessageParamsMarketComputeCommitment__ReprBuilder struct {
	_MessageParamsMarketComputeCommitment__ReprAssembler
}
func (nb *_MessageParamsMarketComputeCommitment__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMarketComputeCommitment__ReprBuilder) Reset() {
	var w _MessageParamsMarketComputeCommitment
	var m schema.Maybe
	*nb = _MessageParamsMarketComputeCommitment__ReprBuilder{_MessageParamsMarketComputeCommitment__ReprAssembler{w: &w, m: &m}}
}
type _MessageParamsMarketComputeCommitment__ReprAssembler struct {
	w *_MessageParamsMarketComputeCommitment
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_DealIDs _List__DealID__ReprAssembler
	ca_SectorType _RegisteredSealProof__ReprAssembler
	}

func (na *_MessageParamsMarketComputeCommitment__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_DealIDs.reset()
	na.ca_SectorType.reset()
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.BeginMap(0)
}
func (na *_MessageParamsMarketComputeCommitment__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMarketComputeCommitment{}
	}
	return na, nil
}
func (na *_MessageParamsMarketComputeCommitment__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.AssignBool(false)
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.AssignInt(0)
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.AssignFloat(0)
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.AssignString("")
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.AssignBytes(nil)
}
func (_MessageParamsMarketComputeCommitment__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.MessageParamsMarketComputeCommitment.Repr"}.AssignLink(nil)
}
func (na *_MessageParamsMarketComputeCommitment__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMarketComputeCommitment); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMarketComputeCommitment.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_MessageParamsMarketComputeCommitment__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketComputeCommitment__ReprPrototype{}
}
func (la *_MessageParamsMarketComputeCommitment__ReprAssembler) valueFinishTidy() bool {
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
func (la *_MessageParamsMarketComputeCommitment__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
		la.ca_DealIDs.w = &la.w.DealIDs
		la.ca_DealIDs.m = &la.cm
		return &la.ca_DealIDs
	case 1:
		la.ca_SectorType.w = &la.w.SectorType
		la.ca_SectorType.m = &la.cm
		return &la.ca_SectorType
	default:
		panic("unreachable")
	}
}
func (la *_MessageParamsMarketComputeCommitment__ReprAssembler) Finish() error {
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
func (la *_MessageParamsMarketComputeCommitment__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}