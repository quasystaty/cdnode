package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/x/permgov/keeper"
	"github.com/cdbo/cdnode/x/permgov/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.PermgovKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
