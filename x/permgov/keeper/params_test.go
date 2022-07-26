package keeper_test

import (
	"testing"

	testkeeper "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/x/permgov/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PermgovKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.Governor, k.Governor(ctx))
}
