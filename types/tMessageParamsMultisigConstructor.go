package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MessageParamsMultisigConstructor struct {
	Signers _List__Address
	NumApprovalsThreshold _Int
	UnlockDuration _ChainEpoch
	StartEpoch _ChainEpoch
}
type MessageParamsMultisigConstructor = *_MessageParamsMultisigConstructor

func (n _MessageParamsMultisigConstructor) FieldSigners()	List__Address {
	return &n.Signers
}
func (n _MessageParamsMultisigConstructor) FieldNumApprovalsThreshold()	Int {
	return &n.NumApprovalsThreshold
}
func (n _MessageParamsMultisigConstructor) FieldUnlockDuration()	ChainEpoch {
	return &n.UnlockDuration
}
func (n _MessageParamsMultisigConstructor) FieldStartEpoch()	ChainEpoch {
	return &n.StartEpoch
}
type _MessageParamsMultisigConstructor__Maybe struct {
	m schema.Maybe
	v MessageParamsMultisigConstructor
}
type MaybeMessageParamsMultisigConstructor = *_MessageParamsMultisigConstructor__Maybe

func (m MaybeMessageParamsMultisigConstructor) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMessageParamsMultisigConstructor) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMessageParamsMultisigConstructor) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMessageParamsMultisigConstructor) AsNode() ipld.Node {
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
func (m MaybeMessageParamsMultisigConstructor) Must() MessageParamsMultisigConstructor {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MessageParamsMultisigConstructor_Signers = _String{"Signers"}
	fieldName__MessageParamsMultisigConstructor_NumApprovalsThreshold = _String{"NumApprovalsThreshold"}
	fieldName__MessageParamsMultisigConstructor_UnlockDuration = _String{"UnlockDuration"}
	fieldName__MessageParamsMultisigConstructor_StartEpoch = _String{"StartEpoch"}
)
var _ ipld.Node = (MessageParamsMultisigConstructor)(&_MessageParamsMultisigConstructor{})
var _ schema.TypedNode = (MessageParamsMultisigConstructor)(&_MessageParamsMultisigConstructor{})
func (MessageParamsMultisigConstructor) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MessageParamsMultisigConstructor) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "Signers":
		return &n.Signers, nil
	case "NumApprovalsThreshold":
		return &n.NumApprovalsThreshold, nil
	case "UnlockDuration":
		return &n.UnlockDuration, nil
	case "StartEpoch":
		return &n.StartEpoch, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n MessageParamsMultisigConstructor) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MessageParamsMultisigConstructor) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.LookupByIndex(0)
}
func (n MessageParamsMultisigConstructor) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MessageParamsMultisigConstructor) MapIterator() ipld.MapIterator {
	return &_MessageParamsMultisigConstructor__MapItr{n, 0}
}

type _MessageParamsMultisigConstructor__MapItr struct {
	n MessageParamsMultisigConstructor
	idx  int
}

func (itr *_MessageParamsMultisigConstructor__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 4 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MessageParamsMultisigConstructor_Signers
		v = &itr.n.Signers
	case 1:
		k = &fieldName__MessageParamsMultisigConstructor_NumApprovalsThreshold
		v = &itr.n.NumApprovalsThreshold
	case 2:
		k = &fieldName__MessageParamsMultisigConstructor_UnlockDuration
		v = &itr.n.UnlockDuration
	case 3:
		k = &fieldName__MessageParamsMultisigConstructor_StartEpoch
		v = &itr.n.StartEpoch
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMultisigConstructor__MapItr) Done() bool {
	return itr.idx >= 4
}

func (MessageParamsMultisigConstructor) ListIterator() ipld.ListIterator {
	return nil
}
func (MessageParamsMultisigConstructor) Length() int {
	return 4
}
func (MessageParamsMultisigConstructor) IsAbsent() bool {
	return false
}
func (MessageParamsMultisigConstructor) IsNull() bool {
	return false
}
func (MessageParamsMultisigConstructor) AsBool() (bool, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.AsBool()
}
func (MessageParamsMultisigConstructor) AsInt() (int, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.AsInt()
}
func (MessageParamsMultisigConstructor) AsFloat() (float64, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.AsFloat()
}
func (MessageParamsMultisigConstructor) AsString() (string, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.AsString()
}
func (MessageParamsMultisigConstructor) AsBytes() ([]byte, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.AsBytes()
}
func (MessageParamsMultisigConstructor) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.MessageParamsMultisigConstructor"}.AsLink()
}
func (MessageParamsMultisigConstructor) Prototype() ipld.NodePrototype {
	return _MessageParamsMultisigConstructor__Prototype{}
}
type _MessageParamsMultisigConstructor__Prototype struct{}

func (_MessageParamsMultisigConstructor__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMultisigConstructor__Builder
	nb.Reset()
	return &nb
}
type _MessageParamsMultisigConstructor__Builder struct {
	_MessageParamsMultisigConstructor__Assembler
}
func (nb *_MessageParamsMultisigConstructor__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMultisigConstructor__Builder) Reset() {
	var w _MessageParamsMultisigConstructor
	var m schema.Maybe
	*nb = _MessageParamsMultisigConstructor__Builder{_MessageParamsMultisigConstructor__Assembler{w: &w, m: &m}}
}
type _MessageParamsMultisigConstructor__Assembler struct {
	w *_MessageParamsMultisigConstructor
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_Signers _List__Address__Assembler
	ca_NumApprovalsThreshold _Int__Assembler
	ca_UnlockDuration _ChainEpoch__Assembler
	ca_StartEpoch _ChainEpoch__Assembler
	}

func (na *_MessageParamsMultisigConstructor__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_Signers.reset()
	na.ca_NumApprovalsThreshold.reset()
	na.ca_UnlockDuration.reset()
	na.ca_StartEpoch.reset()
}

var (
	fieldBit__MessageParamsMultisigConstructor_Signers = 1 << 0
	fieldBit__MessageParamsMultisigConstructor_NumApprovalsThreshold = 1 << 1
	fieldBit__MessageParamsMultisigConstructor_UnlockDuration = 1 << 2
	fieldBit__MessageParamsMultisigConstructor_StartEpoch = 1 << 3
	fieldBits__MessageParamsMultisigConstructor_sufficient = 0 + 1 << 0 + 1 << 1 + 1 << 2 + 1 << 3
)
func (na *_MessageParamsMultisigConstructor__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMultisigConstructor{}
	}
	return na, nil
}
func (_MessageParamsMultisigConstructor__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.BeginList(0)
}
func (na *_MessageParamsMultisigConstructor__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMultisigConstructor__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignBool(false)
}
func (_MessageParamsMultisigConstructor__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignInt(0)
}
func (_MessageParamsMultisigConstructor__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignFloat(0)
}
func (_MessageParamsMultisigConstructor__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignString("")
}
func (_MessageParamsMultisigConstructor__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignBytes(nil)
}
func (_MessageParamsMultisigConstructor__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.MessageParamsMultisigConstructor"}.AssignLink(nil)
}
func (na *_MessageParamsMultisigConstructor__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMultisigConstructor); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMultisigConstructor", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MessageParamsMultisigConstructor__Assembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMultisigConstructor__Prototype{}
}
func (ma *_MessageParamsMultisigConstructor__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_Signers.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 1:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_NumApprovalsThreshold.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 2:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_UnlockDuration.w = nil
			ma.cm = schema.Maybe_Absent
			ma.state = maState_initial
			return true
		default:
			return false
		}
	case 3:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_StartEpoch.w = nil
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
func (ma *_MessageParamsMultisigConstructor__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "Signers":
		if ma.s & fieldBit__MessageParamsMultisigConstructor_Signers != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_Signers}
		}
		ma.s += fieldBit__MessageParamsMultisigConstructor_Signers
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_Signers.w = &ma.w.Signers
		ma.ca_Signers.m = &ma.cm
		return &ma.ca_Signers, nil
	case "NumApprovalsThreshold":
		if ma.s & fieldBit__MessageParamsMultisigConstructor_NumApprovalsThreshold != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_NumApprovalsThreshold}
		}
		ma.s += fieldBit__MessageParamsMultisigConstructor_NumApprovalsThreshold
		ma.state = maState_midValue
		ma.f = 1
		ma.ca_NumApprovalsThreshold.w = &ma.w.NumApprovalsThreshold
		ma.ca_NumApprovalsThreshold.m = &ma.cm
		return &ma.ca_NumApprovalsThreshold, nil
	case "UnlockDuration":
		if ma.s & fieldBit__MessageParamsMultisigConstructor_UnlockDuration != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_UnlockDuration}
		}
		ma.s += fieldBit__MessageParamsMultisigConstructor_UnlockDuration
		ma.state = maState_midValue
		ma.f = 2
		ma.ca_UnlockDuration.w = &ma.w.UnlockDuration
		ma.ca_UnlockDuration.m = &ma.cm
		return &ma.ca_UnlockDuration, nil
	case "StartEpoch":
		if ma.s & fieldBit__MessageParamsMultisigConstructor_StartEpoch != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_StartEpoch}
		}
		ma.s += fieldBit__MessageParamsMultisigConstructor_StartEpoch
		ma.state = maState_midValue
		ma.f = 3
		ma.ca_StartEpoch.w = &ma.w.StartEpoch
		ma.ca_StartEpoch.m = &ma.cm
		return &ma.ca_StartEpoch, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.MessageParamsMultisigConstructor", Key:&_String{k}}
	}
}
func (ma *_MessageParamsMultisigConstructor__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MessageParamsMultisigConstructor__KeyAssembler)(ma)
}
func (ma *_MessageParamsMultisigConstructor__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_Signers.w = &ma.w.Signers
		ma.ca_Signers.m = &ma.cm
		return &ma.ca_Signers
	case 1:
		ma.ca_NumApprovalsThreshold.w = &ma.w.NumApprovalsThreshold
		ma.ca_NumApprovalsThreshold.m = &ma.cm
		return &ma.ca_NumApprovalsThreshold
	case 2:
		ma.ca_UnlockDuration.w = &ma.w.UnlockDuration
		ma.ca_UnlockDuration.m = &ma.cm
		return &ma.ca_UnlockDuration
	case 3:
		ma.ca_StartEpoch.w = &ma.w.StartEpoch
		ma.ca_StartEpoch.m = &ma.cm
		return &ma.ca_StartEpoch
	default:
		panic("unreachable")
	}
}
func (ma *_MessageParamsMultisigConstructor__Assembler) Finish() error {
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
func (ma *_MessageParamsMultisigConstructor__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MessageParamsMultisigConstructor__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MessageParamsMultisigConstructor__KeyAssembler _MessageParamsMultisigConstructor__Assembler
func (_MessageParamsMultisigConstructor__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.BeginMap(0)
}
func (_MessageParamsMultisigConstructor__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.BeginList(0)
}
func (na *_MessageParamsMultisigConstructor__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.AssignNull()
}
func (_MessageParamsMultisigConstructor__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.AssignBool(false)
}
func (_MessageParamsMultisigConstructor__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.AssignInt(0)
}
func (_MessageParamsMultisigConstructor__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MessageParamsMultisigConstructor__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "Signers":
		if ka.s & fieldBit__MessageParamsMultisigConstructor_Signers != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_Signers}
		}
		ka.s += fieldBit__MessageParamsMultisigConstructor_Signers
		ka.state = maState_expectValue
		ka.f = 0
	case "NumApprovalsThreshold":
		if ka.s & fieldBit__MessageParamsMultisigConstructor_NumApprovalsThreshold != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_NumApprovalsThreshold}
		}
		ka.s += fieldBit__MessageParamsMultisigConstructor_NumApprovalsThreshold
		ka.state = maState_expectValue
		ka.f = 1
	case "UnlockDuration":
		if ka.s & fieldBit__MessageParamsMultisigConstructor_UnlockDuration != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_UnlockDuration}
		}
		ka.s += fieldBit__MessageParamsMultisigConstructor_UnlockDuration
		ka.state = maState_expectValue
		ka.f = 2
	case "StartEpoch":
		if ka.s & fieldBit__MessageParamsMultisigConstructor_StartEpoch != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMultisigConstructor_StartEpoch}
		}
		ka.s += fieldBit__MessageParamsMultisigConstructor_StartEpoch
		ka.state = maState_expectValue
		ka.f = 3
	default:
		return ipld.ErrInvalidKey{TypeName:"types.MessageParamsMultisigConstructor", Key:&_String{k}}
	}
	return nil
}
func (_MessageParamsMultisigConstructor__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.AssignBytes(nil)
}
func (_MessageParamsMultisigConstructor__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.MessageParamsMultisigConstructor.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MessageParamsMultisigConstructor__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MessageParamsMultisigConstructor__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MessageParamsMultisigConstructor) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MessageParamsMultisigConstructor) Representation() ipld.Node {
	return (*_MessageParamsMultisigConstructor__Repr)(n)
}
type _MessageParamsMultisigConstructor__Repr _MessageParamsMultisigConstructor
var _ ipld.Node = &_MessageParamsMultisigConstructor__Repr{}
func (_MessageParamsMultisigConstructor__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_MessageParamsMultisigConstructor__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.LookupByString("")
}
func (n *_MessageParamsMultisigConstructor__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_MessageParamsMultisigConstructor__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.Signers.Representation(), nil
	case 1:
		return n.NumApprovalsThreshold.Representation(), nil
	case 2:
		return n.UnlockDuration.Representation(), nil
	case 3:
		return n.StartEpoch.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _MessageParamsMultisigConstructor__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.MessageParamsMultisigConstructor.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_MessageParamsMultisigConstructor__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_MessageParamsMultisigConstructor__Repr) ListIterator() ipld.ListIterator {
	return &_MessageParamsMultisigConstructor__ReprListItr{n, 0}
}

type _MessageParamsMultisigConstructor__ReprListItr struct {
	n   *_MessageParamsMultisigConstructor__Repr
	idx int
	
}

func (itr *_MessageParamsMultisigConstructor__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 4 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.Signers.Representation()
	case 1:
		idx = itr.idx
		v = itr.n.NumApprovalsThreshold.Representation()
	case 2:
		idx = itr.idx
		v = itr.n.UnlockDuration.Representation()
	case 3:
		idx = itr.idx
		v = itr.n.StartEpoch.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMultisigConstructor__ReprListItr) Done() bool {
	return itr.idx >= 4
}

func (rn *_MessageParamsMultisigConstructor__Repr) Length() int {
	l := 4
	return l
}
func (_MessageParamsMultisigConstructor__Repr) IsAbsent() bool {
	return false
}
func (_MessageParamsMultisigConstructor__Repr) IsNull() bool {
	return false
}
func (_MessageParamsMultisigConstructor__Repr) AsBool() (bool, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.AsBool()
}
func (_MessageParamsMultisigConstructor__Repr) AsInt() (int, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.AsInt()
}
func (_MessageParamsMultisigConstructor__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.AsFloat()
}
func (_MessageParamsMultisigConstructor__Repr) AsString() (string, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.AsString()
}
func (_MessageParamsMultisigConstructor__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.AsBytes()
}
func (_MessageParamsMultisigConstructor__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.MessageParamsMultisigConstructor.Repr"}.AsLink()
}
func (_MessageParamsMultisigConstructor__Repr) Prototype() ipld.NodePrototype {
	return _MessageParamsMultisigConstructor__ReprPrototype{}
}
type _MessageParamsMultisigConstructor__ReprPrototype struct{}

func (_MessageParamsMultisigConstructor__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMultisigConstructor__ReprBuilder
	nb.Reset()
	return &nb
}
type _MessageParamsMultisigConstructor__ReprBuilder struct {
	_MessageParamsMultisigConstructor__ReprAssembler
}
func (nb *_MessageParamsMultisigConstructor__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMultisigConstructor__ReprBuilder) Reset() {
	var w _MessageParamsMultisigConstructor
	var m schema.Maybe
	*nb = _MessageParamsMultisigConstructor__ReprBuilder{_MessageParamsMultisigConstructor__ReprAssembler{w: &w, m: &m}}
}
type _MessageParamsMultisigConstructor__ReprAssembler struct {
	w *_MessageParamsMultisigConstructor
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_Signers _List__Address__ReprAssembler
	ca_NumApprovalsThreshold _Int__ReprAssembler
	ca_UnlockDuration _ChainEpoch__ReprAssembler
	ca_StartEpoch _ChainEpoch__ReprAssembler
	}

func (na *_MessageParamsMultisigConstructor__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_Signers.reset()
	na.ca_NumApprovalsThreshold.reset()
	na.ca_UnlockDuration.reset()
	na.ca_StartEpoch.reset()
}
func (_MessageParamsMultisigConstructor__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.BeginMap(0)
}
func (na *_MessageParamsMultisigConstructor__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMultisigConstructor{}
	}
	return na, nil
}
func (na *_MessageParamsMultisigConstructor__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMultisigConstructor__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.AssignBool(false)
}
func (_MessageParamsMultisigConstructor__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.AssignInt(0)
}
func (_MessageParamsMultisigConstructor__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.AssignFloat(0)
}
func (_MessageParamsMultisigConstructor__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.AssignString("")
}
func (_MessageParamsMultisigConstructor__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.AssignBytes(nil)
}
func (_MessageParamsMultisigConstructor__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.MessageParamsMultisigConstructor.Repr"}.AssignLink(nil)
}
func (na *_MessageParamsMultisigConstructor__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMultisigConstructor); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMultisigConstructor.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_MessageParamsMultisigConstructor__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMultisigConstructor__ReprPrototype{}
}
func (la *_MessageParamsMultisigConstructor__ReprAssembler) valueFinishTidy() bool {
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
	case 3:
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
func (la *_MessageParamsMultisigConstructor__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
	if la.f >= 4 {
		return _ErrorThunkAssembler{schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(4)}}
	}
	la.state = laState_midValue
	switch la.f {
	case 0:
		la.ca_Signers.w = &la.w.Signers
		la.ca_Signers.m = &la.cm
		return &la.ca_Signers
	case 1:
		la.ca_NumApprovalsThreshold.w = &la.w.NumApprovalsThreshold
		la.ca_NumApprovalsThreshold.m = &la.cm
		return &la.ca_NumApprovalsThreshold
	case 2:
		la.ca_UnlockDuration.w = &la.w.UnlockDuration
		la.ca_UnlockDuration.m = &la.cm
		return &la.ca_UnlockDuration
	case 3:
		la.ca_StartEpoch.w = &la.w.StartEpoch
		la.ca_StartEpoch.m = &la.cm
		return &la.ca_StartEpoch
	default:
		panic("unreachable")
	}
}
func (la *_MessageParamsMultisigConstructor__ReprAssembler) Finish() error {
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
func (la *_MessageParamsMultisigConstructor__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}