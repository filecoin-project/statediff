package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MessageParamsMinerWithdrawBalance struct {
	AmountRequested _BigInt
}
type MessageParamsMinerWithdrawBalance = *_MessageParamsMinerWithdrawBalance

func (n _MessageParamsMinerWithdrawBalance) FieldAmountRequested()	BigInt {
	return &n.AmountRequested
}
type _MessageParamsMinerWithdrawBalance__Maybe struct {
	m schema.Maybe
	v MessageParamsMinerWithdrawBalance
}
type MaybeMessageParamsMinerWithdrawBalance = *_MessageParamsMinerWithdrawBalance__Maybe

func (m MaybeMessageParamsMinerWithdrawBalance) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMessageParamsMinerWithdrawBalance) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMessageParamsMinerWithdrawBalance) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMessageParamsMinerWithdrawBalance) AsNode() ipld.Node {
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
func (m MaybeMessageParamsMinerWithdrawBalance) Must() MessageParamsMinerWithdrawBalance {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MessageParamsMinerWithdrawBalance_AmountRequested = _String{"AmountRequested"}
)
var _ ipld.Node = (MessageParamsMinerWithdrawBalance)(&_MessageParamsMinerWithdrawBalance{})
var _ schema.TypedNode = (MessageParamsMinerWithdrawBalance)(&_MessageParamsMinerWithdrawBalance{})
func (MessageParamsMinerWithdrawBalance) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MessageParamsMinerWithdrawBalance) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "AmountRequested":
		return &n.AmountRequested, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n MessageParamsMinerWithdrawBalance) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MessageParamsMinerWithdrawBalance) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.LookupByIndex(0)
}
func (n MessageParamsMinerWithdrawBalance) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MessageParamsMinerWithdrawBalance) MapIterator() ipld.MapIterator {
	return &_MessageParamsMinerWithdrawBalance__MapItr{n, 0}
}

type _MessageParamsMinerWithdrawBalance__MapItr struct {
	n MessageParamsMinerWithdrawBalance
	idx  int
}

func (itr *_MessageParamsMinerWithdrawBalance__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 1 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MessageParamsMinerWithdrawBalance_AmountRequested
		v = &itr.n.AmountRequested
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMinerWithdrawBalance__MapItr) Done() bool {
	return itr.idx >= 1
}

func (MessageParamsMinerWithdrawBalance) ListIterator() ipld.ListIterator {
	return nil
}
func (MessageParamsMinerWithdrawBalance) Length() int {
	return 1
}
func (MessageParamsMinerWithdrawBalance) IsAbsent() bool {
	return false
}
func (MessageParamsMinerWithdrawBalance) IsNull() bool {
	return false
}
func (MessageParamsMinerWithdrawBalance) AsBool() (bool, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.AsBool()
}
func (MessageParamsMinerWithdrawBalance) AsInt() (int, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.AsInt()
}
func (MessageParamsMinerWithdrawBalance) AsFloat() (float64, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.AsFloat()
}
func (MessageParamsMinerWithdrawBalance) AsString() (string, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.AsString()
}
func (MessageParamsMinerWithdrawBalance) AsBytes() ([]byte, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.AsBytes()
}
func (MessageParamsMinerWithdrawBalance) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.MessageParamsMinerWithdrawBalance"}.AsLink()
}
func (MessageParamsMinerWithdrawBalance) Prototype() ipld.NodePrototype {
	return _MessageParamsMinerWithdrawBalance__Prototype{}
}
type _MessageParamsMinerWithdrawBalance__Prototype struct{}

func (_MessageParamsMinerWithdrawBalance__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMinerWithdrawBalance__Builder
	nb.Reset()
	return &nb
}
type _MessageParamsMinerWithdrawBalance__Builder struct {
	_MessageParamsMinerWithdrawBalance__Assembler
}
func (nb *_MessageParamsMinerWithdrawBalance__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMinerWithdrawBalance__Builder) Reset() {
	var w _MessageParamsMinerWithdrawBalance
	var m schema.Maybe
	*nb = _MessageParamsMinerWithdrawBalance__Builder{_MessageParamsMinerWithdrawBalance__Assembler{w: &w, m: &m}}
}
type _MessageParamsMinerWithdrawBalance__Assembler struct {
	w *_MessageParamsMinerWithdrawBalance
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_AmountRequested _BigInt__Assembler
	}

func (na *_MessageParamsMinerWithdrawBalance__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_AmountRequested.reset()
}

var (
	fieldBit__MessageParamsMinerWithdrawBalance_AmountRequested = 1 << 0
	fieldBits__MessageParamsMinerWithdrawBalance_sufficient = 0 + 1 << 0
)
func (na *_MessageParamsMinerWithdrawBalance__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMinerWithdrawBalance{}
	}
	return na, nil
}
func (_MessageParamsMinerWithdrawBalance__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.BeginList(0)
}
func (na *_MessageParamsMinerWithdrawBalance__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMinerWithdrawBalance__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignBool(false)
}
func (_MessageParamsMinerWithdrawBalance__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignInt(0)
}
func (_MessageParamsMinerWithdrawBalance__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignFloat(0)
}
func (_MessageParamsMinerWithdrawBalance__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignString("")
}
func (_MessageParamsMinerWithdrawBalance__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignBytes(nil)
}
func (_MessageParamsMinerWithdrawBalance__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.MessageParamsMinerWithdrawBalance"}.AssignLink(nil)
}
func (na *_MessageParamsMinerWithdrawBalance__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMinerWithdrawBalance); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMinerWithdrawBalance", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MessageParamsMinerWithdrawBalance__Assembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMinerWithdrawBalance__Prototype{}
}
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_AmountRequested.w = nil
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
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "AmountRequested":
		if ma.s & fieldBit__MessageParamsMinerWithdrawBalance_AmountRequested != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMinerWithdrawBalance_AmountRequested}
		}
		ma.s += fieldBit__MessageParamsMinerWithdrawBalance_AmountRequested
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_AmountRequested.w = &ma.w.AmountRequested
		ma.ca_AmountRequested.m = &ma.cm
		return &ma.ca_AmountRequested, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.MessageParamsMinerWithdrawBalance", Key:&_String{k}}
	}
}
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MessageParamsMinerWithdrawBalance__KeyAssembler)(ma)
}
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_AmountRequested.w = &ma.w.AmountRequested
		ma.ca_AmountRequested.m = &ma.cm
		return &ma.ca_AmountRequested
	default:
		panic("unreachable")
	}
}
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) Finish() error {
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
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MessageParamsMinerWithdrawBalance__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MessageParamsMinerWithdrawBalance__KeyAssembler _MessageParamsMinerWithdrawBalance__Assembler
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.BeginMap(0)
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.BeginList(0)
}
func (na *_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.AssignNull()
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.AssignBool(false)
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.AssignInt(0)
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "AmountRequested":
		if ka.s & fieldBit__MessageParamsMinerWithdrawBalance_AmountRequested != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMinerWithdrawBalance_AmountRequested}
		}
		ka.s += fieldBit__MessageParamsMinerWithdrawBalance_AmountRequested
		ka.state = maState_expectValue
		ka.f = 0
	default:
		return ipld.ErrInvalidKey{TypeName:"types.MessageParamsMinerWithdrawBalance", Key:&_String{k}}
	}
	return nil
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.AssignBytes(nil)
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.MessageParamsMinerWithdrawBalance.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MessageParamsMinerWithdrawBalance__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MessageParamsMinerWithdrawBalance__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MessageParamsMinerWithdrawBalance) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MessageParamsMinerWithdrawBalance) Representation() ipld.Node {
	return (*_MessageParamsMinerWithdrawBalance__Repr)(n)
}
type _MessageParamsMinerWithdrawBalance__Repr _MessageParamsMinerWithdrawBalance
var _ ipld.Node = &_MessageParamsMinerWithdrawBalance__Repr{}
func (_MessageParamsMinerWithdrawBalance__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_MessageParamsMinerWithdrawBalance__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.LookupByString("")
}
func (n *_MessageParamsMinerWithdrawBalance__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_MessageParamsMinerWithdrawBalance__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.AmountRequested.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _MessageParamsMinerWithdrawBalance__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.MessageParamsMinerWithdrawBalance.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_MessageParamsMinerWithdrawBalance__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_MessageParamsMinerWithdrawBalance__Repr) ListIterator() ipld.ListIterator {
	return &_MessageParamsMinerWithdrawBalance__ReprListItr{n, 0}
}

type _MessageParamsMinerWithdrawBalance__ReprListItr struct {
	n   *_MessageParamsMinerWithdrawBalance__Repr
	idx int
	
}

func (itr *_MessageParamsMinerWithdrawBalance__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 1 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.AmountRequested.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMinerWithdrawBalance__ReprListItr) Done() bool {
	return itr.idx >= 1
}

func (rn *_MessageParamsMinerWithdrawBalance__Repr) Length() int {
	l := 1
	return l
}
func (_MessageParamsMinerWithdrawBalance__Repr) IsAbsent() bool {
	return false
}
func (_MessageParamsMinerWithdrawBalance__Repr) IsNull() bool {
	return false
}
func (_MessageParamsMinerWithdrawBalance__Repr) AsBool() (bool, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.AsBool()
}
func (_MessageParamsMinerWithdrawBalance__Repr) AsInt() (int, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.AsInt()
}
func (_MessageParamsMinerWithdrawBalance__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.AsFloat()
}
func (_MessageParamsMinerWithdrawBalance__Repr) AsString() (string, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.AsString()
}
func (_MessageParamsMinerWithdrawBalance__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.AsBytes()
}
func (_MessageParamsMinerWithdrawBalance__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.MessageParamsMinerWithdrawBalance.Repr"}.AsLink()
}
func (_MessageParamsMinerWithdrawBalance__Repr) Prototype() ipld.NodePrototype {
	return _MessageParamsMinerWithdrawBalance__ReprPrototype{}
}
type _MessageParamsMinerWithdrawBalance__ReprPrototype struct{}

func (_MessageParamsMinerWithdrawBalance__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMinerWithdrawBalance__ReprBuilder
	nb.Reset()
	return &nb
}
type _MessageParamsMinerWithdrawBalance__ReprBuilder struct {
	_MessageParamsMinerWithdrawBalance__ReprAssembler
}
func (nb *_MessageParamsMinerWithdrawBalance__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMinerWithdrawBalance__ReprBuilder) Reset() {
	var w _MessageParamsMinerWithdrawBalance
	var m schema.Maybe
	*nb = _MessageParamsMinerWithdrawBalance__ReprBuilder{_MessageParamsMinerWithdrawBalance__ReprAssembler{w: &w, m: &m}}
}
type _MessageParamsMinerWithdrawBalance__ReprAssembler struct {
	w *_MessageParamsMinerWithdrawBalance
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_AmountRequested _BigInt__ReprAssembler
	}

func (na *_MessageParamsMinerWithdrawBalance__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_AmountRequested.reset()
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.BeginMap(0)
}
func (na *_MessageParamsMinerWithdrawBalance__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMinerWithdrawBalance{}
	}
	return na, nil
}
func (na *_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.AssignBool(false)
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.AssignInt(0)
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.AssignFloat(0)
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.AssignString("")
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.AssignBytes(nil)
}
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.MessageParamsMinerWithdrawBalance.Repr"}.AssignLink(nil)
}
func (na *_MessageParamsMinerWithdrawBalance__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMinerWithdrawBalance); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMinerWithdrawBalance.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_MessageParamsMinerWithdrawBalance__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMinerWithdrawBalance__ReprPrototype{}
}
func (la *_MessageParamsMinerWithdrawBalance__ReprAssembler) valueFinishTidy() bool {
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
	default:
		panic("unreachable")
	}
}
func (la *_MessageParamsMinerWithdrawBalance__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	if la.f >= 1 {
		return _ErrorThunkAssembler{schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(1)}}
	}
	la.state = laState_midValue
	switch la.f {
	case 0:
		la.ca_AmountRequested.w = &la.w.AmountRequested
		la.ca_AmountRequested.m = &la.cm
		return &la.ca_AmountRequested
	default:
		panic("unreachable")
	}
}
func (la *_MessageParamsMinerWithdrawBalance__ReprAssembler) Finish() error {
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
func (la *_MessageParamsMinerWithdrawBalance__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}