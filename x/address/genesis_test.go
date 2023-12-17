package compchain_test

import (
	"testing"

	keepertest "github.com/arnabghose997/compchain/testutil/keeper"
	"github.com/arnabghose997/compchain/testutil/nullify"
	"github.com/arnabghose997/compchain/x/address"
	"github.com/arnabghose997/compchain/x/address/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CompchainKeeper(t)
	compchain.InitGenesis(ctx, *k, genesisState)
	got := compchain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
