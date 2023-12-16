package keeper

import (
	"github.com/arnabghose997/compchain/x/compchain/types"
)

var _ types.QueryServer = Keeper{}
