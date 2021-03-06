package bt

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	"github.com/libsv/go-bt/bscript"
)

/*
General format (inside a block) of each output of a transaction - Txout
Field	                        Description	                                Size
-----------------------------------------------------------------------------------------------------
value                         non negative integer giving the number of   8 bytes
                              Satoshis(BTC/10^8) to be transferred
Txout-script length           non negative integer                        1 - 9 bytes VI = VarInt
Txout-script / scriptPubKey   Script                                      <out-script length>-many bytes
(lockingScript)

*/

// Output is a representation of a transaction output
type Output struct {
	Satoshis      uint64
	LockingScript *bscript.Script
}

// LockingScriptHexString returns the locking script
// of an output encoded as a hex string.
func (o *Output) LockingScriptHexString() string {
	return hex.EncodeToString(*o.LockingScript)
}

func (o *Output) String() string {
	return fmt.Sprintf(`value:     %d
scriptLen: %d
script:    %x
`, o.Satoshis, len(*o.LockingScript), o.LockingScript)
}

// ToBytes encodes the Output into a byte array.
func (o *Output) ToBytes() []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, o.Satoshis)

	h := make([]byte, 0)
	h = append(h, b...)
	h = append(h, VarInt(uint64(len(*o.LockingScript)))...)
	h = append(h, *o.LockingScript...)

	return h
}

// BytesForSigHash returns the proper serialization
// of an output to be hashed and signed (sighash).
func (o *Output) BytesForSigHash() []byte {
	buf := make([]byte, 0)

	satoshis := make([]byte, 8)
	binary.LittleEndian.PutUint64(satoshis, o.Satoshis)
	buf = append(buf, satoshis...)

	buf = append(buf, VarInt(uint64(len(*o.LockingScript)))...)
	buf = append(buf, *o.LockingScript...)

	return buf
}
