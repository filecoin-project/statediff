package blst

type Scalar struct{}
type Fp struct{}
type Fp2 struct{}
type Fp6 struct{}
type Fp12 struct{}
type P1 struct{}
type P2 struct{}
type P1Affine struct{}
type P2Affine struct{}
type Message = []byte
type Pairing = []uint64
type SecretKey = Scalar

func KeyGen(ikm []byte, optional ...[]byte) *SecretKey {
	return nil
}

func PairingCtx(hash_or_encode bool, DST []byte) Pairing {
	return Pairing([]uint64{})
}

func PairingAggregatePkInG1(ctx Pairing, PK *P1Affine, sig *P2Affine,
	msg []byte,
	optional ...[]byte) int {
	return 0
}

func PairingAggregatePkInG2(ctx Pairing, PK *P2Affine, sig *P1Affine,
	msg []byte,
	optional ...[]byte) int {
	return 0
}

func PairingMulNAggregatePkInG1(ctx Pairing, PK *P1Affine, sig *P2Affine,
	rand *Scalar, randBits int, msg []byte,
	optional ...[]byte) int {
	return 0
}

func PairingMulNAggregatePkInG2(ctx Pairing, PK *P2Affine, sig *P1Affine,
	rand *Scalar, randBits int, msg []byte,
	optional ...[]byte) int {
	return 0
}

func PairingCommit(ctx Pairing) {}

func PairingMerge(ctx Pairing, ctx1 Pairing) int {
	return 0
}

func PairingFinalVerify(ctx Pairing, optional ...*Fp12) bool {
	return false
}

func Fp12One() Fp12 {
	return Fp12{}
}

func (pk *P1Affine) From(s *Scalar) *P1Affine {
	return nil
}

func (sig *P2Affine) Sign(sk *SecretKey, msg []byte, dst []byte,
	optional ...interface{}) *P2Affine {
	return nil
}

func (sig *P2Affine) Verify(pk *P1Affine, msg Message, dst []byte,
	optional ...interface{}) bool { // useHash bool, aug []byte
	return false
}

func (dummy *P2Affine) VerifyCompressed(sig []byte, pk []byte,
	msg Message, dst []byte,
	optional ...bool) bool {
	return false
}

func (sig *P2Affine) AggregateVerify(pks []*P1Affine, msgs []Message,
	dst []byte,
	optional ...interface{}) bool {
	return false
}

func (dummy *P2Affine) AggregateVerifyCompressed(sig []byte, pks [][]byte,
	msgs []Message, dst []byte,
	optional ...bool) bool {
	return false
}

func (sig *P2Affine) FastAggregateVerify(pks []*P1Affine, msg Message,
	dst []byte,
	optional ...interface{}) bool {
	return false
}

func (dummy *P2Affine) MultipleAggregateVerify(sigs []*P2Affine,
	pks []*P1Affine, msgs []Message, dst []byte, randFn func(*Scalar),
	randBits int,
	optional ...interface{}) bool {
	return false
}

type P2Aggregate struct {
	v *P2
}

func (agg *P2Aggregate) Aggregate(elmts []*P2Affine) *P2Aggregate {
	return nil
}

func (agg *P2Aggregate) AggregateCompressed(elmts [][]byte) *P2Aggregate {
	return nil
}

func (agg *P2Aggregate) AddAggregate(other *P2Aggregate) *P2Aggregate {
	return nil
}

func (agg *P2Aggregate) Add(elmt *P2Affine) *P2Aggregate {
	return nil
}

func (agg *P2Aggregate) ToAffine() *P2Affine {
	return nil
}

func (pk *P2Affine) From(s *Scalar) *P2Affine {
	return nil
}

func (sig *P1Affine) Sign(sk *SecretKey, msg []byte, dst []byte,
	optional ...interface{}) *P1Affine {
	return nil
}

func (sig *P1Affine) Verify(pk *P2Affine, msg Message, dst []byte,
	optional ...interface{}) bool {
	return false
}

func (dummy *P1Affine) VerifyCompressed(sig []byte, pk []byte,
	msg Message, dst []byte,
	optional ...bool) bool {
	return false
}

func (sig *P1Affine) AggregateVerify(pks []*P2Affine, msgs []Message,
	dst []byte,
	optional ...interface{}) bool {
	return false
}

func (dummy *P1Affine) AggregateVerifyCompressed(sig []byte, pks [][]byte,
	msgs []Message, dst []byte,
	optional ...bool) bool {
	return false
}

func (sig *P1Affine) FastAggregateVerify(pks []*P2Affine, msg Message,
	dst []byte,
	optional ...interface{}) bool {
	return false
}

func (dummy *P1Affine) MultipleAggregateVerify(sigs []*P1Affine,
	pks []*P2Affine, msgs []Message, dst []byte, randFn func(*Scalar),
	randBits int,
	optional ...interface{}) bool {
	return false
}

type P1Aggregate struct {
	v *P1
}

func (agg *P1Aggregate) Aggregate(elmts []*P1Affine) *P1Aggregate {
	return nil
}

func (agg *P1Aggregate) AggregateCompressed(elmts [][]byte) *P1Aggregate {
	return nil
}

func (agg *P1Aggregate) AddAggregate(other *P1Aggregate) *P1Aggregate {
	return nil
}

func (agg *P1Aggregate) Add(elmt *P1Affine) *P1Aggregate {
	return nil
}

func (agg *P1Aggregate) ToAffine() *P1Affine {
	return nil
}

func (p1 *P1Affine) Serialize() []byte {
	return []byte{}
}

func (p1 *P1Affine) Deserialize(in []byte) *P1Affine {
	return nil
}

func (p1 *P1Affine) Compress() []byte {
	return []byte{}
}

func (p1 *P1Affine) Uncompress(in []byte) *P1Affine {
	return nil
}

func (p1 *P1Affine) InG1() bool {
	return false
}

func (dummy *P1Affine) BatchUncompress(in [][]byte) []*P1Affine {
	return nil
}

func (p1 *P1) Serialize() []byte {
	return []byte{}
}

func (p1 *P1) Compress() []byte {
	return []byte{}
}

func (p *P1) ToAffine() *P1Affine {
	return nil
}

func HashToG1(msg []byte, dst []byte,
	optional ...[]byte) *P1 {
	return nil
}

func EncodeToG1(msg []byte, dst []byte,
	optional ...[]byte) *P1 {
	return nil
}

func (p2 *P2Affine) Serialize() []byte {
	return []byte{}
}

func (p2 *P2Affine) Deserialize(in []byte) *P2Affine {
	return nil
}
func (p2 *P2Affine) Compress() []byte {
	return []byte{}
}

func (p2 *P2Affine) Uncompress(in []byte) *P2Affine {
	return nil
}

func (p2 *P2Affine) InG2() bool {
	return false
}

func (dummy *P2Affine) BatchUncompress(in [][]byte) []*P2Affine {
	return nil
}

func (p2 *P2) Serialize() []byte {
	return []byte{}
}

func (p2 *P2) Compress() []byte {
	return []byte{}
}

func (p *P2) ToAffine() *P2Affine {
	return nil
}

func HashToG2(msg []byte, dst []byte,
	optional ...[]byte) *P2 {
	return nil
}

func EncodeToG2(msg []byte, dst []byte,
	optional ...[]byte) *P2 {
	return nil
}

func (s *Scalar) Serialize() []byte {
	return []byte{}
}

func (s *Scalar) Deserialize(in []byte) *Scalar {
	return nil
}

func (s *Scalar) Valid() bool {
	return false
}

func (fr *Scalar) ToLEndian() []byte {
	return []byte{}
}

func (fp *Fp) ToLEndian() []byte {
	return []byte{}
}

func (fr *Scalar) FromLEndian(arr []byte) *Scalar {
	return nil
}

func (fp *Fp) FromLEndian(arr []byte) *Fp {
	return nil
}

func (fr *Scalar) ToBEndian() []byte {
	return []byte{}
}

func (fp *Fp) ToBEndian() []byte {
	return []byte{}
}

func (fr *Scalar) FromBEndian(arr []byte) *Scalar {
	return nil
}

func (fp *Fp) FromBEndian(arr []byte) *Fp {
	return nil
}

func PrintBytes(val []byte, name string) {}

func (s *Scalar) Print(name string) {}

func (p *P1Affine) Print(name string) {}

func (p *P1) Print(name string) {}

func (f *Fp2) Print(name string) {}

func (p *P2Affine) Print(name string) {}

func (p *P2) Print(name string) {}

func (s1 *Scalar) Equals(s2 *Scalar) bool     { return false }
func (e1 *Fp) Equals(e2 *Fp) bool             { return false }
func (e1 *Fp2) Equals(e2 *Fp2) bool           { return false }
func (e1 *P1Affine) Equals(e2 *P1Affine) bool { return false }
func (e1 *P1) Equals(e2 *P1) bool             { return false }
func (e1 *P2Affine) Equals(e2 *P2Affine) bool { return false }
func (e1 *P2) Equals(e2 *P2) bool             { return false }
