package evm

import (
	"context"
	"encoding/hex"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/libocr/gethwrappers2/ocr2aggregator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	evmclimocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/client/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	lpmocks "github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller/mocks"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/txmgr"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

var sampleAddress = testutils.NewAddress()

type mockTransmitter struct{}

func (mockTransmitter) CreateEthTransaction(ctx context.Context, toAddress gethcommon.Address, payload []byte, _ *txmgr.EthTxMeta) error {
	return nil
}
func (mockTransmitter) FromAddress() gethcommon.Address { return sampleAddress }

func TestContractTransmitter(t *testing.T) {
	t.Parallel()

	lggr := logger.TestLogger(t)
	c := evmclimocks.NewClient(t)
	lp := lpmocks.NewLogPoller(t)
	// scanLogs = false
	digestAndEpochDontScanLogs, _ := hex.DecodeString(
		"0000000000000000000000000000000000000000000000000000000000000000" + // false
			"000130da6b9315bd59af6b0a3f5463c0d0a39e92eaa34cbcbdbace7b3bfcc776" + // config digest
			"0000000000000000000000000000000000000000000000000000000000000002") // epoch
	c.On("CallContract", mock.Anything, mock.Anything, mock.Anything).Return(digestAndEpochDontScanLogs, nil).Once()
	contractABI, _ := abi.JSON(strings.NewReader(ocr2aggregator.OCR2AggregatorABI))
	lp.On("RegisterFilter", mock.Anything).Return(nil)
	ot, err := NewOCRContractTransmitter(gethcommon.Address{}, c, contractABI, mockTransmitter{}, lp, lggr, func(b []byte) (*txmgr.EthTxMeta, error) {
		return &txmgr.EthTxMeta{}, nil
	})
	require.NoError(t, err)
	digest, epoch, err := ot.LatestConfigDigestAndEpoch(testutils.Context(t))
	require.NoError(t, err)
	assert.Equal(t, "000130da6b9315bd59af6b0a3f5463c0d0a39e92eaa34cbcbdbace7b3bfcc776", hex.EncodeToString(digest[:]))
	assert.Equal(t, uint32(2), epoch)

	// scanLogs = true
	digestAndEpochScanLogs, _ := hex.DecodeString(
		"0000000000000000000000000000000000000000000000000000000000000001" + // false
			"000130da6b9315bd59af6b0a3f5463c0d0a39e92eaa34cbcbdbace7b3bfcc776" + // config digest
			"0000000000000000000000000000000000000000000000000000000000000002") // epoch
	c.On("CallContract", mock.Anything, mock.Anything, mock.Anything).Return(digestAndEpochScanLogs, nil).Once()
	transmitted2, _ := hex.DecodeString(
		"000130da6b9315bd59af6b0a3f5463c0d0a39e92eaa34cbcbdbace7b3bfcc777" + // config digest
			"0000000000000000000000000000000000000000000000000000000000000002") // epoch
	lp.On("LatestLogByEventSigWithConfs",
		mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&logpoller.Log{
		Data: transmitted2,
	}, nil)
	digest, epoch, err = ot.LatestConfigDigestAndEpoch(testutils.Context(t))
	require.NoError(t, err)
	assert.Equal(t, "000130da6b9315bd59af6b0a3f5463c0d0a39e92eaa34cbcbdbace7b3bfcc777", hex.EncodeToString(digest[:]))
	assert.Equal(t, uint32(2), epoch)
	from, err := ot.FromAccount()
	require.NoError(t, err)
	assert.Equal(t, sampleAddress.String(), string(from))
}
