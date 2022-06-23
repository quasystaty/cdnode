package keeper_test

import (
	"testing"

	testkeeper "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/x/cdnode/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CdnodeKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
