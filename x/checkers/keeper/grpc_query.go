package keeper

import (
	"github.com/kempy007/cosmos-checkers/x/checkers/types"
)

var _ types.QueryServer = Keeper{}
