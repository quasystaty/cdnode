package keeper

import (
	"github.com/cdbo/cdnode/x/permgov/types"
)

var _ types.QueryServer = Keeper{}
