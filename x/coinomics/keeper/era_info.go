package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/haqq-network/haqq/x/coinomics/types"
)

func (k Keeper) GetInflation(ctx sdk.Context) sdk.Dec {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixInflation)
	if len(bz) == 0 {
		return sdk.ZeroDec()
	}

	var inflationValue sdk.Dec
	err := inflationValue.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal inflationValue value: %w", err))
	}

	return inflationValue
}

func (k Keeper) SetInflation(ctx sdk.Context, inflation sdk.Dec) {
	binaryInfValue, err := inflation.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value: %w", err))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixInflation, binaryInfValue)
}

// GetEra gets current era
func (k Keeper) GetEra(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEra)
	if len(bz) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetEra stores the current era
func (k Keeper) SetEra(ctx sdk.Context, era uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEra, sdk.Uint64ToBigEndian(era))
}

// GetStartEraBlock gets current era start block number
func (k Keeper) GetEraStartedAtBlock(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEraStartedAtBlock)
	if len(bz) == 0 {
		return 0
	}

	return sdk.BigEndianToUint64(bz)
}

// SetStartEraBlock stores the start era block number
func (k Keeper) SetEraStartedAtBlock(ctx sdk.Context, block uint64) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEraStartedAtBlock, sdk.Uint64ToBigEndian(block))
}

func (k Keeper) GetEraTargetMint(ctx sdk.Context) sdk.Coin {
	params := k.GetParams(ctx)

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KetPrefixEraTargetMint)
	if len(bz) == 0 {
		return sdk.NewCoin(params.MintDenom, sdk.ZeroInt())
	}

	var eraTragetMintValue sdk.Coin
	err := eraTragetMintValue.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal eraTragetMintValue value: %w", err))
	}

	return eraTragetMintValue
}

func (k Keeper) SetEraTargetMint(ctx sdk.Context, eraMint sdk.Coin) {
	binaryEraTragetMintValue, err := eraMint.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value: %w", err))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KetPrefixEraTargetMint, binaryEraTragetMintValue)
}

func (k Keeper) GetEraTargetSupply(ctx sdk.Context) sdk.Coin {
	params := k.GetParams(ctx)

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixEraTargetSupply)
	if len(bz) == 0 {
		return sdk.NewCoin(params.MintDenom, sdk.ZeroInt())
	}

	var eraTarget sdk.Coin
	err := eraTarget.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal eraTarget value: %w", err))
	}

	return eraTarget
}

func (k Keeper) SetEraTargetSupply(ctx sdk.Context, eraTargetSupply sdk.Coin) {
	binaryEraSupply, err := eraTargetSupply.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value: %w", err))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixEraTargetSupply, binaryEraSupply)
}

func (k Keeper) GetTargetTotalSupply(ctx sdk.Context) sdk.Coin {
	params := k.GetParams(ctx)

	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.KeyPrefixTargetTotalSupply)
	if len(bz) == 0 {
		return sdk.NewCoin(params.MintDenom, sdk.ZeroInt())
	}

	var targetTotal sdk.Coin
	err := targetTotal.Unmarshal(bz)
	if err != nil {
		panic(fmt.Errorf("unable to unmarshal targetTotal value: %w", err))
	}

	return targetTotal
}

func (k Keeper) SetTargetTotalSupply(ctx sdk.Context, totalTarget sdk.Coin) {
	binaryTotalTarget, err := totalTarget.Marshal()
	if err != nil {
		panic(fmt.Errorf("unable to marshal amount value: %w", err))
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyPrefixTargetTotalSupply, binaryTotalTarget)
}
