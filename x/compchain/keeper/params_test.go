package keeper_test

import (
	"testing"

	testkeeper "github.com/arnabghose997/compchain/testutil/keeper"
	"github.com/arnabghose997/compchain/x/compchain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CompchainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
