package keeper_test

import (
	"testing"

	testkeeper "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/x/coinmaster/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func TestParamsQuery(t *testing.T) {
	k, ctx := testkeeper.NewCoinmasterKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	params := types.DefaultParams()
	k.SetParams(ctx, params)

	response, err := k.Params(wctx, &types.QueryParamsRequest{})
	require.NoError(t, err)
	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
}
