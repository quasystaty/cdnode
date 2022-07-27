package keeper

import (
	"github.com/cdbo/cdnode/x/coinmaster/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Minters(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Minters returns the Minters param
func (k Keeper) Minters(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyMinters, &res)
	return
}
