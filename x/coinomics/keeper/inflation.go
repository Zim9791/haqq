package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/haqq-network/haqq/x/coinomics/types"
)

// NextPhase calculus
func (k Keeper) CountEraForBlock(ctx sdk.Context, params types.Params, currentEra uint64, currentBlock uint64) uint64 {
	if currentEra == 0 {
		return 1
	}

	params = k.GetParams(ctx)

	startedBlock := k.GetEraStartedAtBlock(ctx)
	nextEraBlock := params.BlocksPerEra + startedBlock

	fmt.Println("!!! CountEraForBlock startedBlock: ", startedBlock)
	fmt.Println("!!! CountEraForBlock nextEraBlock: ", nextEraBlock)
	fmt.Println("!!! CountEraForBlock currentBlock: ", currentBlock)

	if currentBlock < nextEraBlock {
		return currentEra
	}

	return currentEra + 1
}

func (k Keeper) CalcTargetMintForEra(ctx sdk.Context, eraNumber uint64) sdk.Coin {
	params := k.GetParams(ctx)

	eraCoef := sdk.NewDecWithPrec(95, 2) // 0.95

	if eraNumber == 1 {
		eraPeriod := uint64(2)
		currentTotalSupply := k.bankKeeper.GetSupply(ctx, "aISLM")
		targetTotalSupply := k.GetTargetTotalSupply(ctx)

		totalMintNeeded := targetTotalSupply.SubAmount(currentTotalSupply.Amount)

		// (1-era_coef)*total_mint_needed/(1-era_coef^(100/era_period))
		calc_part1 := (sdk.OneDec().Sub(eraCoef)).Mul(totalMintNeeded.Amount.ToDec())
		calc_part2 := sdk.OneDec().Sub(eraCoef.Power(100 / eraPeriod))

		target := calc_part1.Quo(calc_part2)

		return sdk.NewCoin(params.MintDenom, target.RoundInt())
	} else if eraNumber > 1 && eraNumber < 50 {
		prevTarget := k.GetEraTargetMint(ctx)
		target := prevTarget.Amount.ToDec().Mul(eraCoef)

		return sdk.NewCoin("aISLM", target.RoundInt())
	} else if eraNumber == 50 {
		currentTotalSupply := k.bankKeeper.GetSupply(ctx, "aISLM")
		targetTotalSupply := k.GetTargetTotalSupply(ctx)

		return targetTotalSupply.SubAmount(currentTotalSupply.Amount)
	} else {
		return sdk.NewCoin(params.MintDenom, sdk.NewInt(0))
	}
}

func (k Keeper) CalcInflation(ctx sdk.Context, era uint64, eraTargetSupply sdk.Coin, eraTargetMint sdk.Coin) sdk.Dec {
	if era > 50 {
		return sdk.NewDec(0)
	}

	return eraTargetMint.Amount.ToDec().
		Quo(eraTargetSupply.SubAmount(eraTargetMint.Amount).Amount.ToDec()).
		Mul(sdk.NewDec(100))
}

//

func (k Keeper) GetProportion(
	coin sdk.Coin,
	distribution sdk.Dec,
) sdk.Coin {
	return sdk.NewCoin(
		coin.Denom,
		coin.Amount.ToDec().Mul(distribution).TruncateInt(),
	)
}

func (k Keeper) MintAndAllocateInflation(ctx sdk.Context) error {
	params := k.GetParams(ctx)
	eraTargetMint := k.GetEraTargetMint(ctx)

	totalMintOnBlockInt := eraTargetMint.Amount.Quo(sdk.NewInt(int64(params.BlocksPerEra)))
	totalMintOnBlockCoin := sdk.NewCoin(params.MintDenom, totalMintOnBlockInt)

	// Mint coins to coinomics module
	k.MintCoins(ctx, totalMintOnBlockCoin)

	coinomicsModuleAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	// // Allocate first part of coins to EvergreenDAO
	// evergreen := sdk.NewCoins(k.GetProportion(totalMintOnBlockCoin, params.MintDistribution.EvergreenDao))
	// err := k.distrKeeper.FundCommunityPool(
	// 	ctx,
	// 	evergreen,
	// 	coinomicsModuleAddr,
	// )
	// if err != nil {
	// 	return err
	// }

	// Allocate remaining coinomics module balance to stacking rewards
	staking := k.bankKeeper.GetAllBalances(ctx, coinomicsModuleAddr)
	err := k.bankKeeper.SendCoinsFromModuleToModule(
		ctx,
		types.ModuleName,
		k.feeCollectorName,
		staking,
	)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) MintCoins(ctx sdk.Context, coin sdk.Coin) error {
	coins := sdk.NewCoins(coin)

	// skip as no coins need to be minted
	if coins.Empty() {
		return nil
	}

	return k.bankKeeper.MintCoins(ctx, types.ModuleName, coins)
}
