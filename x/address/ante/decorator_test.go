package ante

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/stretchr/testify/require"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

func TestMsgSend(t *testing.T) {
	s := SetupTestSuite(t, true)
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	affr := NewAddressFundReceiveRestriction()
	anteHandler := sdk.ChainAnteDecorators(affr)

	// Create sender account
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// Addresses ending with letter 's' are barred from receiving funds
	blockedAddr, err := sdk.AccAddressFromBech32("cosmos1utvd62ft4sgkzk9k2ukxcyvt599euay9588j0s")
	require.NoError(t, err)

	// Create MsgSend
	coinsToSend := sdk.NewCoins(sdk.NewInt64Coin("token", 10000000))
	s.txBuilder.SetMsgs(banktypes.NewMsgSend(addr1, blockedAddr, coinsToSend))
	s.txBuilder.SetGasLimit(200000000)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	txInvalid, err := s.CreateTestTx(privs, accNums, accSeqs, s.ctx.ChainID())
	require.NoError(t, err)

	_, err = anteHandler(s.ctx, txInvalid, false)
	require.Error(t, err)
}

func TestMsgMultiSend(t *testing.T) {
	s := SetupTestSuite(t, true)
	s.txBuilder = s.clientCtx.TxConfig.NewTxBuilder()

	affr := NewAddressFundReceiveRestriction()
	anteHandler := sdk.ChainAnteDecorators(affr)

	// Create sender account
	priv1, _, addr1 := testdata.KeyTestPubAddr()

	// Addresses ending with letter 's' are barred from receiving funds
	blockedAddr, err := sdk.AccAddressFromBech32("cosmos1utvd62ft4sgkzk9k2ukxcyvt599euay9588j0s")
	require.NoError(t, err)

	// Create MsgMultiSend
	coinsToSend := sdk.NewCoins(sdk.NewInt64Coin("token", 10000000))
	s.txBuilder.SetMsgs(
		banktypes.NewMsgMultiSend(
			[]banktypes.Input{{
				Address: addr1.String(),
				Coins:   coinsToSend,
			}},
			[]banktypes.Output{
				{
					Address: blockedAddr.String(),
					Coins:   coinsToSend,
				},
			},
		),
	)
	s.txBuilder.SetGasLimit(200000000)

	privs, accNums, accSeqs := []cryptotypes.PrivKey{priv1}, []uint64{0}, []uint64{0}
	txInvalid, err := s.CreateTestTx(privs, accNums, accSeqs, s.ctx.ChainID())
	require.NoError(t, err)

	_, err = anteHandler(s.ctx, txInvalid, false)
	require.Error(t, err)
}
