package types

// Code generated by go-ipld-prime gengo.  DO NOT EDIT.

import (
	ipld "github.com/ipld/go-ipld-prime"
	"github.com/ipld/go-ipld-prime/node/mixins"
	"github.com/ipld/go-ipld-prime/schema"
)

type _MessageParamsMarketPublishDeals struct {
	Deals _List__ClientDealProposal
}
type MessageParamsMarketPublishDeals = *_MessageParamsMarketPublishDeals

func (n _MessageParamsMarketPublishDeals) FieldDeals()	List__ClientDealProposal {
	return &n.Deals
}
type _MessageParamsMarketPublishDeals__Maybe struct {
	m schema.Maybe
	v MessageParamsMarketPublishDeals
}
type MaybeMessageParamsMarketPublishDeals = *_MessageParamsMarketPublishDeals__Maybe

func (m MaybeMessageParamsMarketPublishDeals) IsNull() bool {
	return m.m == schema.Maybe_Null
}
func (m MaybeMessageParamsMarketPublishDeals) IsAbsent() bool {
	return m.m == schema.Maybe_Absent
}
func (m MaybeMessageParamsMarketPublishDeals) Exists() bool {
	return m.m == schema.Maybe_Value
}
func (m MaybeMessageParamsMarketPublishDeals) AsNode() ipld.Node {
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
func (m MaybeMessageParamsMarketPublishDeals) Must() MessageParamsMarketPublishDeals {
	if !m.Exists() {
		panic("unbox of a maybe rejected")
	}
	return m.v
}
var (
	fieldName__MessageParamsMarketPublishDeals_Deals = _String{"Deals"}
)
var _ ipld.Node = (MessageParamsMarketPublishDeals)(&_MessageParamsMarketPublishDeals{})
var _ schema.TypedNode = (MessageParamsMarketPublishDeals)(&_MessageParamsMarketPublishDeals{})
func (MessageParamsMarketPublishDeals) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_Map
}
func (n MessageParamsMarketPublishDeals) LookupByString(key string) (ipld.Node, error) {
	switch key {
	case "Deals":
		return &n.Deals, nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfString(key)}
	}
}
func (n MessageParamsMarketPublishDeals) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ks, err := key.AsString()
	if err != nil {
		return nil, err
	}
	return n.LookupByString(ks)
}
func (MessageParamsMarketPublishDeals) LookupByIndex(idx int) (ipld.Node, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.LookupByIndex(0)
}
func (n MessageParamsMarketPublishDeals) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	return n.LookupByString(seg.String())
}
func (n MessageParamsMarketPublishDeals) MapIterator() ipld.MapIterator {
	return &_MessageParamsMarketPublishDeals__MapItr{n, 0}
}

type _MessageParamsMarketPublishDeals__MapItr struct {
	n MessageParamsMarketPublishDeals
	idx  int
}

func (itr *_MessageParamsMarketPublishDeals__MapItr) Next() (k ipld.Node, v ipld.Node, _ error) {
	if itr.idx >= 1 {
		return nil, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		k = &fieldName__MessageParamsMarketPublishDeals_Deals
		v = &itr.n.Deals
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMarketPublishDeals__MapItr) Done() bool {
	return itr.idx >= 1
}

func (MessageParamsMarketPublishDeals) ListIterator() ipld.ListIterator {
	return nil
}
func (MessageParamsMarketPublishDeals) Length() int {
	return 1
}
func (MessageParamsMarketPublishDeals) IsAbsent() bool {
	return false
}
func (MessageParamsMarketPublishDeals) IsNull() bool {
	return false
}
func (MessageParamsMarketPublishDeals) AsBool() (bool, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.AsBool()
}
func (MessageParamsMarketPublishDeals) AsInt() (int, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.AsInt()
}
func (MessageParamsMarketPublishDeals) AsFloat() (float64, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.AsFloat()
}
func (MessageParamsMarketPublishDeals) AsString() (string, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.AsString()
}
func (MessageParamsMarketPublishDeals) AsBytes() ([]byte, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.AsBytes()
}
func (MessageParamsMarketPublishDeals) AsLink() (ipld.Link, error) {
	return mixins.Map{"types.MessageParamsMarketPublishDeals"}.AsLink()
}
func (MessageParamsMarketPublishDeals) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketPublishDeals__Prototype{}
}
type _MessageParamsMarketPublishDeals__Prototype struct{}

func (_MessageParamsMarketPublishDeals__Prototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMarketPublishDeals__Builder
	nb.Reset()
	return &nb
}
type _MessageParamsMarketPublishDeals__Builder struct {
	_MessageParamsMarketPublishDeals__Assembler
}
func (nb *_MessageParamsMarketPublishDeals__Builder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMarketPublishDeals__Builder) Reset() {
	var w _MessageParamsMarketPublishDeals
	var m schema.Maybe
	*nb = _MessageParamsMarketPublishDeals__Builder{_MessageParamsMarketPublishDeals__Assembler{w: &w, m: &m}}
}
type _MessageParamsMarketPublishDeals__Assembler struct {
	w *_MessageParamsMarketPublishDeals
	m *schema.Maybe
	state maState
	s int
	f int

	cm schema.Maybe
	ca_Deals _List__ClientDealProposal__Assembler
	}

func (na *_MessageParamsMarketPublishDeals__Assembler) reset() {
	na.state = maState_initial
	na.s = 0
	na.ca_Deals.reset()
}

var (
	fieldBit__MessageParamsMarketPublishDeals_Deals = 1 << 0
	fieldBits__MessageParamsMarketPublishDeals_sufficient = 0 + 1 << 0
)
func (na *_MessageParamsMarketPublishDeals__Assembler) BeginMap(int) (ipld.MapAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMarketPublishDeals{}
	}
	return na, nil
}
func (_MessageParamsMarketPublishDeals__Assembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.BeginList(0)
}
func (na *_MessageParamsMarketPublishDeals__Assembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMarketPublishDeals__Assembler) AssignBool(bool) error {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignBool(false)
}
func (_MessageParamsMarketPublishDeals__Assembler) AssignInt(int) error {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignInt(0)
}
func (_MessageParamsMarketPublishDeals__Assembler) AssignFloat(float64) error {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignFloat(0)
}
func (_MessageParamsMarketPublishDeals__Assembler) AssignString(string) error {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignString("")
}
func (_MessageParamsMarketPublishDeals__Assembler) AssignBytes([]byte) error {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignBytes(nil)
}
func (_MessageParamsMarketPublishDeals__Assembler) AssignLink(ipld.Link) error {
	return mixins.MapAssembler{"types.MessageParamsMarketPublishDeals"}.AssignLink(nil)
}
func (na *_MessageParamsMarketPublishDeals__Assembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMarketPublishDeals); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMarketPublishDeals", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustMap, ActualKind: v.ReprKind()}
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
func (_MessageParamsMarketPublishDeals__Assembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketPublishDeals__Prototype{}
}
func (ma *_MessageParamsMarketPublishDeals__Assembler) valueFinishTidy() bool {
	switch ma.f {
	case 0:
		switch ma.cm {
		case schema.Maybe_Value:
			ma.ca_Deals.w = nil
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
func (ma *_MessageParamsMarketPublishDeals__Assembler) AssembleEntry(k string) (ipld.NodeAssembler, error) {
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
	case "Deals":
		if ma.s & fieldBit__MessageParamsMarketPublishDeals_Deals != 0 {
			return nil, ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketPublishDeals_Deals}
		}
		ma.s += fieldBit__MessageParamsMarketPublishDeals_Deals
		ma.state = maState_midValue
		ma.f = 0
		ma.ca_Deals.w = &ma.w.Deals
		ma.ca_Deals.m = &ma.cm
		return &ma.ca_Deals, nil
	default:
		return nil, ipld.ErrInvalidKey{TypeName:"types.MessageParamsMarketPublishDeals", Key:&_String{k}}
	}
}
func (ma *_MessageParamsMarketPublishDeals__Assembler) AssembleKey() ipld.NodeAssembler {
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
	return (*_MessageParamsMarketPublishDeals__KeyAssembler)(ma)
}
func (ma *_MessageParamsMarketPublishDeals__Assembler) AssembleValue() ipld.NodeAssembler {
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
		ma.ca_Deals.w = &ma.w.Deals
		ma.ca_Deals.m = &ma.cm
		return &ma.ca_Deals
	default:
		panic("unreachable")
	}
}
func (ma *_MessageParamsMarketPublishDeals__Assembler) Finish() error {
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
func (ma *_MessageParamsMarketPublishDeals__Assembler) KeyPrototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (ma *_MessageParamsMarketPublishDeals__Assembler) ValuePrototype(k string) ipld.NodePrototype {
	panic("todo structbuilder mapassembler valueprototype")
}
type _MessageParamsMarketPublishDeals__KeyAssembler _MessageParamsMarketPublishDeals__Assembler
func (_MessageParamsMarketPublishDeals__KeyAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.BeginMap(0)
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) BeginList(sizeHint int) (ipld.ListAssembler, error) {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.BeginList(0)
}
func (na *_MessageParamsMarketPublishDeals__KeyAssembler) AssignNull() error {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.AssignNull()
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) AssignBool(bool) error {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.AssignBool(false)
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) AssignInt(int) error {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.AssignInt(0)
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) AssignFloat(float64) error {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.AssignFloat(0)
}
func (ka *_MessageParamsMarketPublishDeals__KeyAssembler) AssignString(k string) error {
	if ka.state != maState_midKey {
		panic("misuse: KeyAssembler held beyond its valid lifetime")
	}
	switch k {
	case "Deals":
		if ka.s & fieldBit__MessageParamsMarketPublishDeals_Deals != 0 {
			return ipld.ErrRepeatedMapKey{&fieldName__MessageParamsMarketPublishDeals_Deals}
		}
		ka.s += fieldBit__MessageParamsMarketPublishDeals_Deals
		ka.state = maState_expectValue
		ka.f = 0
	default:
		return ipld.ErrInvalidKey{TypeName:"types.MessageParamsMarketPublishDeals", Key:&_String{k}}
	}
	return nil
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) AssignBytes([]byte) error {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.AssignBytes(nil)
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) AssignLink(ipld.Link) error {
	return mixins.StringAssembler{"types.MessageParamsMarketPublishDeals.KeyAssembler"}.AssignLink(nil)
}
func (ka *_MessageParamsMarketPublishDeals__KeyAssembler) AssignNode(v ipld.Node) error {
	if v2, err := v.AsString(); err != nil {
		return err
	} else {
		return ka.AssignString(v2)
	}
}
func (_MessageParamsMarketPublishDeals__KeyAssembler) Prototype() ipld.NodePrototype {
	return _String__Prototype{}
}
func (MessageParamsMarketPublishDeals) Type() schema.Type {
	return nil /*TODO:typelit*/
}
func (n MessageParamsMarketPublishDeals) Representation() ipld.Node {
	return (*_MessageParamsMarketPublishDeals__Repr)(n)
}
type _MessageParamsMarketPublishDeals__Repr _MessageParamsMarketPublishDeals
var _ ipld.Node = &_MessageParamsMarketPublishDeals__Repr{}
func (_MessageParamsMarketPublishDeals__Repr) ReprKind() ipld.ReprKind {
	return ipld.ReprKind_List
}
func (_MessageParamsMarketPublishDeals__Repr) LookupByString(string) (ipld.Node, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.LookupByString("")
}
func (n *_MessageParamsMarketPublishDeals__Repr) LookupByNode(key ipld.Node) (ipld.Node, error) {
	ki, err := key.AsInt()
	if err != nil {
		return nil, err
	}
	return n.LookupByIndex(ki)
}
func (n *_MessageParamsMarketPublishDeals__Repr) LookupByIndex(idx int) (ipld.Node, error) {
	switch idx {
	case 0:
		return n.Deals.Representation(), nil
	default:
		return nil, schema.ErrNoSuchField{Type: nil /*TODO*/, Field: ipld.PathSegmentOfInt(idx)}
	}
}
func (n _MessageParamsMarketPublishDeals__Repr) LookupBySegment(seg ipld.PathSegment) (ipld.Node, error) {
	i, err := seg.Index()
	if err != nil {
		return nil, ipld.ErrInvalidSegmentForList{TypeName: "types.MessageParamsMarketPublishDeals.Repr", TroubleSegment: seg, Reason: err}
	}
	return n.LookupByIndex(i)
}
func (_MessageParamsMarketPublishDeals__Repr) MapIterator() ipld.MapIterator {
	return nil
}
func (n *_MessageParamsMarketPublishDeals__Repr) ListIterator() ipld.ListIterator {
	return &_MessageParamsMarketPublishDeals__ReprListItr{n, 0}
}

type _MessageParamsMarketPublishDeals__ReprListItr struct {
	n   *_MessageParamsMarketPublishDeals__Repr
	idx int
	
}

func (itr *_MessageParamsMarketPublishDeals__ReprListItr) Next() (idx int, v ipld.Node, err error) {
	if itr.idx >= 1 {
		return -1, nil, ipld.ErrIteratorOverread{}
	}
	switch itr.idx {
	case 0:
		idx = itr.idx
		v = itr.n.Deals.Representation()
	default:
		panic("unreachable")
	}
	itr.idx++
	return
}
func (itr *_MessageParamsMarketPublishDeals__ReprListItr) Done() bool {
	return itr.idx >= 1
}

func (rn *_MessageParamsMarketPublishDeals__Repr) Length() int {
	l := 1
	return l
}
func (_MessageParamsMarketPublishDeals__Repr) IsAbsent() bool {
	return false
}
func (_MessageParamsMarketPublishDeals__Repr) IsNull() bool {
	return false
}
func (_MessageParamsMarketPublishDeals__Repr) AsBool() (bool, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.AsBool()
}
func (_MessageParamsMarketPublishDeals__Repr) AsInt() (int, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.AsInt()
}
func (_MessageParamsMarketPublishDeals__Repr) AsFloat() (float64, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.AsFloat()
}
func (_MessageParamsMarketPublishDeals__Repr) AsString() (string, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.AsString()
}
func (_MessageParamsMarketPublishDeals__Repr) AsBytes() ([]byte, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.AsBytes()
}
func (_MessageParamsMarketPublishDeals__Repr) AsLink() (ipld.Link, error) {
	return mixins.List{"types.MessageParamsMarketPublishDeals.Repr"}.AsLink()
}
func (_MessageParamsMarketPublishDeals__Repr) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketPublishDeals__ReprPrototype{}
}
type _MessageParamsMarketPublishDeals__ReprPrototype struct{}

func (_MessageParamsMarketPublishDeals__ReprPrototype) NewBuilder() ipld.NodeBuilder {
	var nb _MessageParamsMarketPublishDeals__ReprBuilder
	nb.Reset()
	return &nb
}
type _MessageParamsMarketPublishDeals__ReprBuilder struct {
	_MessageParamsMarketPublishDeals__ReprAssembler
}
func (nb *_MessageParamsMarketPublishDeals__ReprBuilder) Build() ipld.Node {
	if *nb.m != schema.Maybe_Value {
		panic("invalid state: cannot call Build on an assembler that's not finished")
	}
	return nb.w
}
func (nb *_MessageParamsMarketPublishDeals__ReprBuilder) Reset() {
	var w _MessageParamsMarketPublishDeals
	var m schema.Maybe
	*nb = _MessageParamsMarketPublishDeals__ReprBuilder{_MessageParamsMarketPublishDeals__ReprAssembler{w: &w, m: &m}}
}
type _MessageParamsMarketPublishDeals__ReprAssembler struct {
	w *_MessageParamsMarketPublishDeals
	m *schema.Maybe
	state laState
	f int

	cm schema.Maybe
	ca_Deals _List__ClientDealProposal__ReprAssembler
	}

func (na *_MessageParamsMarketPublishDeals__ReprAssembler) reset() {
	na.state = laState_initial
	na.f = 0
	na.ca_Deals.reset()
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) BeginMap(sizeHint int) (ipld.MapAssembler, error) {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.BeginMap(0)
}
func (na *_MessageParamsMarketPublishDeals__ReprAssembler) BeginList(int) (ipld.ListAssembler, error) {
	switch *na.m {
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: it makes no sense to 'begin' twice on the same assembler!")
	}
	*na.m = midvalue
	if na.w == nil {
		na.w = &_MessageParamsMarketPublishDeals{}
	}
	return na, nil
}
func (na *_MessageParamsMarketPublishDeals__ReprAssembler) AssignNull() error {
	switch *na.m {
	case allowNull:
		*na.m = schema.Maybe_Null
		return nil
	case schema.Maybe_Absent:
		return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr.Repr"}.AssignNull()
	case schema.Maybe_Value, schema.Maybe_Null:
		panic("invalid state: cannot assign into assembler that's already finished")
	case midvalue:
		panic("invalid state: cannot assign null into an assembler that's already begun working on recursive structures!")
	}
	panic("unreachable")
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) AssignBool(bool) error {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.AssignBool(false)
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) AssignInt(int) error {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.AssignInt(0)
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) AssignFloat(float64) error {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.AssignFloat(0)
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) AssignString(string) error {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.AssignString("")
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) AssignBytes([]byte) error {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.AssignBytes(nil)
}
func (_MessageParamsMarketPublishDeals__ReprAssembler) AssignLink(ipld.Link) error {
	return mixins.ListAssembler{"types.MessageParamsMarketPublishDeals.Repr"}.AssignLink(nil)
}
func (na *_MessageParamsMarketPublishDeals__ReprAssembler) AssignNode(v ipld.Node) error {
	if v.IsNull() {
		return na.AssignNull()
	}
	if v2, ok := v.(*_MessageParamsMarketPublishDeals); ok {
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
		return ipld.ErrWrongKind{TypeName: "types.MessageParamsMarketPublishDeals.Repr", MethodName: "AssignNode", AppropriateKind: ipld.ReprKindSet_JustList, ActualKind: v.ReprKind()}
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
func (_MessageParamsMarketPublishDeals__ReprAssembler) Prototype() ipld.NodePrototype {
	return _MessageParamsMarketPublishDeals__ReprPrototype{}
}
func (la *_MessageParamsMarketPublishDeals__ReprAssembler) valueFinishTidy() bool {
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
func (la *_MessageParamsMarketPublishDeals__ReprAssembler) AssembleValue() ipld.NodeAssembler {
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
		la.ca_Deals.w = &la.w.Deals
		la.ca_Deals.m = &la.cm
		return &la.ca_Deals
	default:
		panic("unreachable")
	}
}
func (la *_MessageParamsMarketPublishDeals__ReprAssembler) Finish() error {
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
func (la *_MessageParamsMarketPublishDeals__ReprAssembler) ValuePrototype(_ int) ipld.NodePrototype {
	panic("todo structbuilder tuplerepr valueprototype")
}