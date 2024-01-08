package bscript

import (
	"testing"

	"github.com/bitcoinsv/bsvutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewP2PKHUnlockingScript(t *testing.T) {

	t.Run("unlock script with valid pubkey", func(t *testing.T) {

		wif, err := bsvutil.DecodeWIF("KznvCNc6Yf4iztSThoMH6oHWzH9EgjfodKxmeuUGPq5DEX5maspS")
		require.NoError(t, err)
		assert.NotNil(t, wif)

		var script *Script
		script, err = NewP2PKHUnlockingScript(wif.SerializePubKey(), []byte("some-signature"), 0)
		require.NoError(t, err)
		assert.NotNil(t, script)
		assert.Equal(t, "0f736f6d652d7369676e6174757265002102798913bc057b344de675dac34faafe3dc2f312c758cd9068209f810877306d66", script.ToString())
	})

}
