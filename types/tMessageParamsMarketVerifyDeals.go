package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MessageParamsMarketVerifyDeals struct {
	DealIDs _List__DealID
	SectorExpiry _ChainEpoch
	SectorStart _ChainEpoch
}
type MessageParamsMarketVerifyDeals = *_MessageParamsMarketVerifyDeals

func (n _MessageParamsMarketVerifyDeals) FieldDealIDs()	List__DealID {
	return &n.DealIDs
}
func (n _MessageParamsMarketVerifyDeals) FieldSectorExpiry()	ChainEpoch {
	return &n.SectorExpiry
}
func (n _MessageParamsMarketVerifyDeals) FieldSectorStart()	ChainEpoch {
	return &n.SectorStart
}
type _MessageParamsMarketVerifyDeals__Maybe struct {
	m schema.Maybe
	v MessageParamsMarketVerifyDeals
}
type MaybeMessageParamsMarketVerifyDeals = *_MessageParamsMarketVerifyDeals__Maybe

func (m MaybeMessageParamsMarketVerifyDeals) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMessageParamsMarketVerifyDeals) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMessageParamsMarketVerifyDeals) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMessageParamsMarketVerifyDeals) AsNode() ipld.Node {
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
func (m MaybeMessageParamsMarketVerifyDeals) Must() MessageParamsMarketVerifyDeals {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MessageParamsMarketVerifyDeals_DealIDs = _String{"DealIDs"}
	fieldName__MessageParamsMarketVerifyDeals_SectorExpiry = _String{"SectorExpiry"}
	fieldName__MessageParamsMarketVerifyDeals_SectorStart = _String{"SectorStart"}
)
var _ ipld.Node = (MessageParamsMarketVerifyDeals)(&_MessageParamsMarketVerifyDeals{})
var _ schema.TypedNode = (MessageParamsMarketVerifyDeals)(&_MessageParamsMarketVerifyDeals{})
func (MessageParamsMarketVerifyDeals) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MessageParamsMarketVerifyDeals) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "DealIDs":
		return &n.DealIDs, nil
	case "SectorExpiry":
		return &n.SectorExpiry, nil
	case "SectorStart":
		return &n.SectorStart, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n MessageParamsMarketVerifyDeals) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MessageParamsMarketVerifyDeals) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.LookupByIndex(0)
}
func (n MessageParamsMarketVerifyDeals) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MessageParamsMarketVerifyDeals) MapIterator() ipld.MapIterator {
	return &_MessageParamsMarketVerifyDeals__MapItr{n, 0}
}

type _MessageParamsMarketVerifyDeals__MapItr struct {
	n MessageParamsMarketVerifyDeals
	idx  int
}

func (itr *_MessageParamsMarketVerifyDeals__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 3 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MessageParamsMarketVerifyDeals_DealIDs
		v = &itr.n.DealIDs
	case 1:
		k = &fieldName__MessageParamsMarketVerifyDeals_SectorExpiry
		v = &itr.n.SectorExpiry
	case 2:
		k = &fieldName__MessageParamsMarketVerifyDeals_SectorStart
		v = &itr.n.SectorStart
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMarketVerifyDeals__MapItr) Done() bool {
	return itr.idx >= 3
}

func (MessageParamsMarketVerifyDeals) ListIterator() ipld.ListIterator {
	return nil
}
func (MessageParamsMarketVerifyDeals) Length() int {
	return 3
}
func (MessageParamsMarketVerifyDeals) IsAbsent() bool {
	return false
}
func (MessageParamsMarketVerifyDeals) IsNull() bool {
	return false
}
func (MessageParamsMarketVerifyDeals) AsBool() (bool, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.AsBool()
}
func (MessageParamsMarketVerifyDeals) AsInt() (int, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.AsInt()
}
func (MessageParamsMarketVerifyDeals) AsFloat() (float64, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.AsFloat()
}
func (MessageParamsMarketVerifyDeals) AsString() (string, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.AsString()
}
func (MessageParamsMarketVerifyDeals) AsBytes() ([]byte, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.AsBytes()
}
func (MessageParamsMarketVerifyDeals) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.MessageParamsMarketVerifyDeals"}.AsLink()
}
func (MessageParamsMarketVerifyDeals) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketVerifyDeals__Prototype{}
}
type _MessageParamsMarketVerifyDeals__Prototype struct{}

func (_MessageParamsMarketVerifyDeals__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMarketVerifyDeals__Builder
	nb.Reset()
	return &nb
}
type _MessageParamsMarketVerifyDeals__Builder struct {
	_MessageParamsMarketVerifyDeals__Assembler
}
func (nb *_MessageParamsMarketVerifyDeals__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMarketVerifyDeals__Builder) Reset() {
	var w _MessageParamsMarketVerifyDeals
	var m schema.Maybe
	*nb = _MessageParamsMarketVerifyDeals__Builder{_MessageParamsMarketVerifyDeals__Assembler{w: &w, m: &m}}
}
type _MessageParamsMarketVerifyDeals__Assembler struct {
	w *_MessageParamsMarketVerifyDeals
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_DealIDs _List__DealID__Assembler
	ca_SectorExpiry _ChainEpoch__Assembler
	ca_SectorStart _ChainEpoch__Assembler
	}

func (na *_MessageParamsMarketVerifyDeals__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_DealIDs.reset()
	na.ca_SectorExpiry.reset()
	na.ca_SectorStart.reset()
}

var (
	fieldBit__MessageParamsMarketVerifyDeals_DealIDs = 1 << 0
	fieldBit__MessageParamsMarketVerifyDeals_SectorExpiry = 1 << 1
	fieldBit__MessageParamsMarketVerifyDeals_SectorStart = 1 << 2
	fieldBits__MessageParamsMarketVerifyDeals_sufficient = 0 + 1 << 0 + 1 << 1 + 1 << 2
)
func (na *_MessageParamsMarketVerifyDeals__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMarketVerifyDeals{}
	}
	return na, nil
}
func (_MessageParamsMarketVerifyDeals__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.BeginList(0)
}
func (na *_MessageParamsMarketVerifyDeals__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMarketVerifyDeals__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignBool(false)
}
func (_MessageParamsMarketVerifyDeals__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignInt(0)
}
func (_MessageParamsMarketVerifyDeals__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignFloat(0)
}
func (_MessageParamsMarketVerifyDeals__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignString("")
}
func (_MessageParamsMarketVerifyDeals__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignBytes(nil)
}
func (_MessageParamsMarketVerifyDeals__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.MessageParamsMarketVerifyDeals"}.AssignLink(nil)
}
func (na *_MessageParamsMarketVerifyDeals__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMarketVerifyDeals); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMarketVerifyDeals", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MessageParamsMarketVerifyDeals__Assembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketVerifyDeals__Prototype{}
}
func (ma *_MessageParamsMarketVerifyDeals__Assembler) valueFinishTidy() bool {
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
			ma.ca_SectorExpiry.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 2:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_SectorStart.w = nil
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
func (ma *_MessageParamsMarketVerifyDeals__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
		if ma.s & fieldBit__MessageParamsMarketVerifyDeals_DealIDs != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketVerifyDeals_DealIDs}
		}
		ma.s += fieldBit__MessageParamsMarketVerifyDeals_DealIDs
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_DealIDs.w = &ma.w.DealIDs
		ma.ca_DealIDs.m = &ma.cm
		return &ma.ca_DealIDs, nil
	case "SectorExpiry":
		if ma.s & fieldBit__MessageParamsMarketVerifyDeals_SectorExpiry != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketVerifyDeals_SectorExpiry}
		}
		ma.s += fieldBit__MessageParamsMarketVerifyDeals_SectorExpiry
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_SectorExpiry.w = &ma.w.SectorExpiry
		ma.ca_SectorExpiry.m = &ma.cm
		return &ma.ca_SectorExpiry, nil
	case "SectorStart":
		if ma.s & fieldBit__MessageParamsMarketVerifyDeals_SectorStart != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketVerifyDeals_SectorStart}
		}
		ma.s += fieldBit__MessageParamsMarketVerifyDeals_SectorStart
		ma.state = maState_midValue
		ma.f = 2
		ma.ca_SectorStart.w = &ma.w.SectorStart
		ma.ca_SectorStart.m = &ma.cm
		return &ma.ca_SectorStart, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.MessageParamsMarketVerifyDeals", Key:&_String{k}}
	}
}
func (ma *_MessageParamsMarketVerifyDeals__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MessageParamsMarketVerifyDeals__KeyAssembler)(ma)
}
func (ma *_MessageParamsMarketVerifyDeals__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_SectorExpiry.w = &ma.w.SectorExpiry
		ma.ca_SectorExpiry.m = &ma.cm
		return &ma.ca_SectorExpiry
	case 2:
		ma.ca_SectorStart.w = &ma.w.SectorStart
		ma.ca_SectorStart.m = &ma.cm
		return &ma.ca_SectorStart
	default:
		panic("unreachable")
	}
}
func (ma *_MessageParamsMarketVerifyDeals__Assembler) Finish() error {
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
func (ma *_MessageParamsMarketVerifyDeals__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MessageParamsMarketVerifyDeals__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MessageParamsMarketVerifyDeals__KeyAssembler _MessageParamsMarketVerifyDeals__Assembler
func (_MessageParamsMarketVerifyDeals__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.BeginMap(0)
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.BeginList(0)
}
func (na *_MessageParamsMarketVerifyDeals__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.AssignNull()
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.AssignBool(false)
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.AssignInt(0)
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MessageParamsMarketVerifyDeals__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "DealIDs":
		if ka.s & fieldBit__MessageParamsMarketVerifyDeals_DealIDs != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketVerifyDeals_DealIDs}
		}
		ka.s += fieldBit__MessageParamsMarketVerifyDeals_DealIDs
		ka.state = maState_expectValue
		ka.f = 0
	case "SectorExpiry":
		if ka.s & fieldBit__MessageParamsMarketVerifyDeals_SectorExpiry != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketVerifyDeals_SectorExpiry}
		}
		ka.s += fieldBit__MessageParamsMarketVerifyDeals_SectorExpiry
		ka.state = maState_expectValue
		ka.f = 1
	case "SectorStart":
		if ka.s & fieldBit__MessageParamsMarketVerifyDeals_SectorStart != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketVerifyDeals_SectorStart}
		}
		ka.s += fieldBit__MessageParamsMarketVerifyDeals_SectorStart
		ka.state = maState_expectValue
		ka.f = 2
	default:
		return ipld.ErrInvalidKey{TypeName:"types.MessageParamsMarketVerifyDeals", Key:&_String{k}}
	}
	return nil
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.AssignBytes(nil)
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.MessageParamsMarketVerifyDeals.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MessageParamsMarketVerifyDeals__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MessageParamsMarketVerifyDeals__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MessageParamsMarketVerifyDeals) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MessageParamsMarketVerifyDeals) Representation() ipld.Node {
	return (*_MessageParamsMarketVerifyDeals__Repr)(n)
}
type _MessageParamsMarketVerifyDeals__Repr _MessageParamsMarketVerifyDeals
var _ ipld.Node = &_MessageParamsMarketVerifyDeals__Repr{}
func (_MessageParamsMarketVerifyDeals__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_MessageParamsMarketVerifyDeals__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.LookupByString("")
}
func (n *_MessageParamsMarketVerifyDeals__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_MessageParamsMarketVerifyDeals__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.DealIDs.Representation(), nil
	case 1:
		return n.SectorExpiry.Representation(), nil
	case 2:
		return n.SectorStart.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _MessageParamsMarketVerifyDeals__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.MessageParamsMarketVerifyDeals.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_MessageParamsMarketVerifyDeals__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_MessageParamsMarketVerifyDeals__Repr) ListIterator() ipld.ListIterator {
	return &_MessageParamsMarketVerifyDeals__ReprListItr{n, 0}
}

type _MessageParamsMarketVerifyDeals__ReprListItr struct {
	n   *_MessageParamsMarketVerifyDeals__Repr
	idx int
	
}

func (itr *_MessageParamsMarketVerifyDeals__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 3 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.DealIDs.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.SectorExpiry.Representation()
	case 2:
		idx = itr.idx
		v = itr.n.SectorStart.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMarketVerifyDeals__ReprListItr) Done() bool {
	return itr.idx >= 3
}

func (rn *_MessageParamsMarketVerifyDeals__Repr) Length() int {
	l := 3
	return l
}
func (_MessageParamsMarketVerifyDeals__Repr) IsAbsent() bool {
	return false
}
func (_MessageParamsMarketVerifyDeals__Repr) IsNull() bool {
	return false
}
func (_MessageParamsMarketVerifyDeals__Repr) AsBool() (bool, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.AsBool()
}
func (_MessageParamsMarketVerifyDeals__Repr) AsInt() (int, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.AsInt()
}
func (_MessageParamsMarketVerifyDeals__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.AsFloat()
}
func (_MessageParamsMarketVerifyDeals__Repr) AsString() (string, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.AsString()
}
func (_MessageParamsMarketVerifyDeals__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.AsBytes()
}
func (_MessageParamsMarketVerifyDeals__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.MessageParamsMarketVerifyDeals.Repr"}.AsLink()
}
func (_MessageParamsMarketVerifyDeals__Repr) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketVerifyDeals__ReprPrototype{}
}
type _MessageParamsMarketVerifyDeals__ReprPrototype struct{}

func (_MessageParamsMarketVerifyDeals__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMarketVerifyDeals__ReprBuilder
	nb.Reset()
	return &nb
}
type _MessageParamsMarketVerifyDeals__ReprBuilder struct {
	_MessageParamsMarketVerifyDeals__ReprAssembler
}
func (nb *_MessageParamsMarketVerifyDeals__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMarketVerifyDeals__ReprBuilder) Reset() {
	var w _MessageParamsMarketVerifyDeals
	var m schema.Maybe
	*nb = _MessageParamsMarketVerifyDeals__ReprBuilder{_MessageParamsMarketVerifyDeals__ReprAssembler{w: &w, m: &m}}
}
type _MessageParamsMarketVerifyDeals__ReprAssembler struct {
	w *_MessageParamsMarketVerifyDeals
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_DealIDs _List__DealID__ReprAssembler
	ca_SectorExpiry _ChainEpoch__ReprAssembler
	ca_SectorStart _ChainEpoch__ReprAssembler
	}

func (na *_MessageParamsMarketVerifyDeals__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_DealIDs.reset()
	na.ca_SectorExpiry.reset()
	na.ca_SectorStart.reset()
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.BeginMap(0)
}
func (na *_MessageParamsMarketVerifyDeals__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMarketVerifyDeals{}
	}
	return na, nil
}
func (na *_MessageParamsMarketVerifyDeals__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.AssignBool(false)
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.AssignInt(0)
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.AssignFloat(0)
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.AssignString("")
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.AssignBytes(nil)
}
func (_MessageParamsMarketVerifyDeals__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.MessageParamsMarketVerifyDeals.Repr"}.AssignLink(nil)
}
func (na *_MessageParamsMarketVerifyDeals__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMarketVerifyDeals); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMarketVerifyDeals.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_MessageParamsMarketVerifyDeals__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketVerifyDeals__ReprPrototype{}
}
func (la *_MessageParamsMarketVerifyDeals__ReprAssembler) valueFinishTidy() bool {
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
func (la *_MessageParamsMarketVerifyDeals__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
		la.ca_DealIDs.w = &la.w.DealIDs
		la.ca_DealIDs.m = &la.cm
		return &la.ca_DealIDs
	case 1:
		la.ca_SectorExpiry.w = &la.w.SectorExpiry
		la.ca_SectorExpiry.m = &la.cm
		return &la.ca_SectorExpiry
	case 2:
		la.ca_SectorStart.w = &la.w.SectorStart
		la.ca_SectorStart.m = &la.cm
		return &la.ca_SectorStart
	default:
		panic("unreachable")
	}
}
func (la *_MessageParamsMarketVerifyDeals__ReprAssembler) Finish() error {
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
func (la *_MessageParamsMarketVerifyDeals__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}