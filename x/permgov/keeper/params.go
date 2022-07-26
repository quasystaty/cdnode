package keeper

import (
	"github.com/cdbo/cdnode/x/permgov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.Governor(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// Governor returns the Governor param
func (k Keeper) Governor(ctx sdk.Context) (res string) {
	k.paramstore.Get(ctx, types.KeyGovernor, &res)
	return
}
