package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/arnabghose997/compchain/testutil/keeper"
	"github.com/arnabghose997/compchain/x/compchain/keeper"
	"github.com/arnabghose997/compchain/x/compchain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CompchainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
