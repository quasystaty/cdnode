package keeper

import (
	"github.com/cdbo/cdnode/x/coinmaster/types"
)

var _ types.QueryServer = Keeper{}
