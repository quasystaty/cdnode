package keeper_test

import (
	"testing"

	testkeeper "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/x/coinmaster/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.NewCoinmasterKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Minters, k.Minters(ctx))
}
