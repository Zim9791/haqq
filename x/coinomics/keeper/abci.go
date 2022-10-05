package keeper

import (
	// "strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	// sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	// vestexported "github.com/cosmos/cosmos-sdk/x/auth/vesting/exported"
	// "github.com/evmos/evmos/v7/x/claims/types"
)

func (k Keeper) EndBlocker(ctx sdk.Context) {
	params := k.GetParams(ctx)

	// NOTE: ignore end of block if coinomics is disabled
	if !params.EnableCoinomics {
		return
	}

	//

	currentBlock := uint64(ctx.BlockHeight())
	currentEra := k.GetEra(ctx)
	eraForBlock := k.CountEraForBlock(ctx, params, currentEra, currentBlock)

	if currentEra != eraForBlock {
		k.SetEra(ctx, eraForBlock)
		k.SetEraStartedAtBlock(ctx, currentBlock)

		nextEraTargetMint := k.CalcTargetMintForEra(ctx, eraForBlock)

		currentTotalSupply := k.bankKeeper.GetSupply(ctx, "aISLM")
		nextEraTargetSupply := currentTotalSupply.AddAmount(nextEraTargetMint.Amount)
		nextEraInflation := k.CalcInflation(ctx, eraForBlock, nextEraTargetSupply, nextEraTargetMint)

		k.SetEraTargetMint(ctx, nextEraTargetMint)
		k.SetEraTargetSupply(ctx, nextEraTargetSupply)
		k.SetInflation(ctx, nextEraInflation)
	}

	k.MintAndAllocateInflation(ctx)
}
