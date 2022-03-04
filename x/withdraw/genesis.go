package incentives

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/tharsis/evmos/x/withdraw/keeper"
	"github.com/tharsis/evmos/x/withdraw/types"
)

// InitGenesis import module genesis
func InitGenesis(
	ctx sdk.Context,
	k keeper.Keeper,
	data types.GenesisState,
) {
	k.SetParams(ctx, data.Params)
}

// ExportGenesis export module status
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}