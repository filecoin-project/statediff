package types

// Type is a struct embeding a NodePrototype/Type for every Node implementation in this package.
// One of its major uses is to start the construction of a value.
// You can use it like this:
//
// 		types.Type.YourTypeName.NewBuilder().BeginMap() //...
//
// and:
//
// 		types.Type.OtherTypeName.NewBuilder().AssignString("x") // ...
//
var Type typeSlab

type typeSlab struct {
	AccountV0State       _AccountV0State__Prototype
	AccountV0State__Repr _AccountV0State__ReprPrototype
	Address       _Address__Prototype
	Address__Repr _Address__ReprPrototype
	Any       _Any__Prototype
	Any__Repr _Any__ReprPrototype
	Bool       _Bool__Prototype
	Bool__Repr _Bool__ReprPrototype
	Bytes       _Bytes__Prototype
	Bytes__Repr _Bytes__ReprPrototype
	Float       _Float__Prototype
	Float__Repr _Float__ReprPrototype
	Int       _Int__Prototype
	Int__Repr _Int__ReprPrototype
	Link       _Link__Prototype
	Link__Repr _Link__ReprPrototype
	List       _List__Prototype
	List__Repr _List__ReprPrototype
	Map       _Map__Prototype
	Map__Repr _Map__ReprPrototype
	String       _String__Prototype
	String__Repr _String__ReprPrototype
}
