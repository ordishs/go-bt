package bt_test

import (
	"bytes"
	"encoding/hex"
	"testing"

	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/chainhash"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	txStr                    = "010000000110ee96aa946338cfd0b2ed0603259cfe2f5458c32ee4bd7b88b583769c6b046e010000006b483045022100e5e4749d539a163039769f52e1ebc8e6f62e39387d61e1a305bd722116cded6c022014924b745dd02194fe6b5cb8ac88ee8e9a2aede89e680dcea6169ea696e24d52012102b4b754609b46b5d09644c2161f1767b72b93847ce8154d795f95d31031a08aa2ffffffff028098f34c010000001976a914a134408afa258a50ed7a1d9817f26b63cc9002cc88ac8028bb13010000001976a914fec5b1145596b35f59f8be1daf169f375942143388ac00000000"
	txBytes, _               = hex.DecodeString(txStr)
	inputBytes, _            = hex.DecodeString("10ee96aa946338cfd0b2ed0603259cfe2f5458c32ee4bd7b88b583769c6b046e010000006b483045022100e5e4749d539a163039769f52e1ebc8e6f62e39387d61e1a305bd722116cded6c022014924b745dd02194fe6b5cb8ac88ee8e9a2aede89e680dcea6169ea696e24d52012102b4b754609b46b5d09644c2161f1767b72b93847ce8154d795f95d31031a08aa2ffffffff")
	previousTxIDBytes, _     = hex.DecodeString("10ee96aa946338cfd0b2ed0603259cfe2f5458c32ee4bd7b88b583769c6b046e")
	previousTxIDStr          = "6e046b9c7683b5887bbde42ec358542ffe9c250306edb2d0cf386394aa96ee10"
	previousTxIDChainHash, _ = chainhash.NewHashFromStr(previousTxIDStr)
)

func TestValidTransaction(t *testing.T) {
	tx, err := bt.NewTxFromBytes(txBytes)
	require.NoError(t, err)

	assert.Equal(t, 1, len(tx.Inputs))
	assert.Equal(t, 2, len(tx.Outputs))
	assert.Equal(t, previousTxIDStr, tx.Inputs[0].PreviousTxIDStr())
	assert.Equal(t, previousTxIDStr, tx.Inputs[0].PreviousTxIDChainHash().String())
	// t.Logf("%x", tx.Inputs[0].PreviousTxIDChainHash().CloneBytes())
}

func TestValidTransaction2(t *testing.T) {
	tx, err := bt.NewTxFromString(txStr)
	require.NoError(t, err)

	assert.Equal(t, 1, len(tx.Inputs))
	assert.Equal(t, 2, len(tx.Outputs))
	assert.Equal(t, previousTxIDStr, tx.Inputs[0].PreviousTxIDStr())
	assert.Equal(t, previousTxIDStr, tx.Inputs[0].PreviousTxIDChainHash().String())
	// t.Logf("%x", tx.Inputs[0].PreviousTxIDChainHash().CloneBytes())
}

func TestPreviousTxId(t *testing.T) {
	// PreviousTxIDAdd
	// PreviousTxIDAddStr
	// PreviousTxID
	// PreviousTxIDStr
	// PreviousTxIDChainHash

}

func TestReadFrom(t *testing.T) {
	assert.Equal(t, 148, len(inputBytes))

	r := bytes.NewReader(inputBytes)

	i := &bt.Input{}
	n, err := i.ReadFrom(r)
	require.NoError(t, err)

	assert.Equal(t, int64(148), n)

	assert.Equal(t, previousTxIDStr, i.PreviousTxIDStr())
	assert.Equal(t, previousTxIDStr, i.PreviousTxIDChainHash().String())
}

// The following test will fail because the bytes are incorrectly not reversed in the input
// func TestPreviousTxIDAdd(t *testing.T) {
// 	i := &bt.Input{}
// 	err := i.PreviousTxIDAdd(previousTxIDBytes)
// 	require.NoError(t, err)

// 	assert.Equal(t, previousTxIDStr, i.PreviousTxIDStr())
// 	assert.Equal(t, previousTxIDStr, i.PreviousTxIDChainHash().String())
// }

func TestPreviousTxIDAddStr(t *testing.T) {
	i := &bt.Input{}
	err := i.PreviousTxIDAddStr(previousTxIDStr)
	require.NoError(t, err)

	assert.Equal(t, previousTxIDStr, i.PreviousTxIDStr())
	assert.Equal(t, previousTxIDStr, i.PreviousTxIDChainHash().String())
}
