package bt_test

import (
	"testing"

	"github.com/libsv/go-bt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetDataFee(t *testing.T) {
	t.Run("get valid data fee", func(t *testing.T) {
		fees := []*bt.Fee{bt.DefaultDataFee()}
		fee, err := bt.GetDataFee(fees)
		require.NoError(t, err)
		assert.NotNil(t, fee)
		assert.Equal(t, bt.FeeTypeData, fee.FeeType)
		assert.Equal(t, 5, fee.MiningFee.Satoshis)
		assert.Equal(t, 10, fee.MiningFee.Bytes)
		assert.Equal(t, 5, fee.RelayFee.Satoshis)
		assert.Equal(t, 10, fee.RelayFee.Bytes)
	})

	t.Run("no data fee found", func(t *testing.T) {
		wrongFee := bt.DefaultDataFee()
		wrongFee.FeeType = "unknown"
		fees := []*bt.Fee{wrongFee}
		fee, err := bt.GetDataFee(fees)
		require.Error(t, err)
		assert.Nil(t, fee)
	})
}

func TestGetStandardFee(t *testing.T) {
	t.Run("get valid standard fee", func(t *testing.T) {
		fees := []*bt.Fee{bt.DefaultStandardFee()}
		fee, err := bt.GetStandardFee(fees)
		require.NoError(t, err)
		assert.NotNil(t, fee)
		assert.Equal(t, bt.FeeTypeStandard, fee.FeeType)
		assert.Equal(t, 5, fee.MiningFee.Satoshis)
		assert.Equal(t, 10, fee.MiningFee.Bytes)
		assert.Equal(t, 5, fee.RelayFee.Satoshis)
		assert.Equal(t, 10, fee.RelayFee.Bytes)
	})

	t.Run("no standard fee found", func(t *testing.T) {
		wrongFee := bt.DefaultStandardFee()
		wrongFee.FeeType = "unknown"
		fees := []*bt.Fee{wrongFee}
		fee, err := bt.GetStandardFee(fees)
		require.Error(t, err)
		assert.Nil(t, fee)
	})
}

func TestDefaultFees(t *testing.T) {
	fees := bt.DefaultFees()
	assert.Len(t, fees, 2)

	fee, err := bt.GetDataFee(fees)
	require.NoError(t, err)
	assert.NotNil(t, fee)
	assert.Equal(t, bt.FeeTypeData, fee.FeeType)

	fee, err = bt.GetStandardFee(fees)
	require.NoError(t, err)
	assert.NotNil(t, fee)
	assert.Equal(t, bt.FeeTypeStandard, fee.FeeType)
}
