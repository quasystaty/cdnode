package keeper_test

import (
	"fmt"
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
	creator := "noria197hfvr6nfd2228xhlyykc234yhwm6tps2drjx8"
	denom := "unoria"
	wCtx := sdk.WrapSDKContext(ctx)

	mintAmount, _ := sdk.ParseCoinNormalized(fmt.Sprintf("100000%v", denom))

	// Set SDK account prefix
	sdk.GetConfig().SetBech32PrefixForAccount("noria", "noria")
	creatorAccAddress, _ := sdk.AccAddressFromBech32(creator)

	// Set single minter and single denom
	k.SetParams(ctx, types.NewParams(creator, denom))

	// BankKeeper expectations
	bankKeeperMock.EXPECT().MintCoins(ctx, "coinmaster", sdk.NewCoins(mintAmount)).Return(nil).Once()
	bankKeeperMock.EXPECT().SendCoinsFromModuleToAccount(ctx, "coinmaster", creatorAccAddress, sdk.NewCoins(mintAmount)).Return(nil).Once()

	// Base case
	resp, err := msgServer.Mint(wCtx, &types.MsgMint{
		Creator: creator,
		Amount:  mintAmount,
	})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
}
