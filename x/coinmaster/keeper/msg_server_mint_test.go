package keeper_test

import (
	"testing"

	"github.com/cdbo/cdnode/mocks"
	keepertest "github.com/cdbo/cdnode/testutil/keeper"
	"github.com/cdbo/cdnode/x/coinmaster/keeper"
	"github.com/cdbo/cdnode/x/coinmaster/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"
)

func Test_Mint(t *testing.T) {

	bankKeeperMock := mocks.NewBankKeeper(t)
	k, ctx := keepertest.NewCoinmasterKeeperWithBankKeeper(t, bankKeeperMock)
	msgServer := keeper.NewMsgServerImpl(*k)

	mintAmount, _ := sdk.ParseCoinNormalized("100000unoria")

	// Base case
	resp, err := msgServer.Mint(ctx.Context(), &types.MsgMint{
		Creator: "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8",
		Amount:  mintAmount,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)

}
