package permgov_test

import (
	"testing"

	keepertest "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/testutil/nullify"
	"github.com/cdbo/cdnode/x/permgov"
	"github.com/cdbo/cdnode/x/permgov/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PermgovKeeper(t)
	permgov.InitGenesis(ctx, *k, genesisState)
	got := permgov.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
