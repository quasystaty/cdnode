package keeper

import (
	"github.com/cdbo/cdnode/x/cdnode/types"
)

var _ types.QueryServer = Keeper{}
