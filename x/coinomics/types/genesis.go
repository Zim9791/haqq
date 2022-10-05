package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewGenesisState(
	params Params,
	inflation sdk.Dec,
	era uint64,
	eraStartedAtBlock uint64,
	eraTargetMint sdk.Coin,
	eraTargetSupply sdk.Coin,
	targetTotalSupply sdk.Coin,
) GenesisState {
	return GenesisState{
		Params:            params,
		Inflation:         inflation,
		Era:               era,
		EraStartedAtBlock: eraStartedAtBlock,
		EraTargetMint:     eraTargetMint,
		EraTargetSupply:   eraTargetSupply,
		TargetTotalSupply: targetTotalSupply,
	}
}

func DefaultGenesisState() *GenesisState {
	params := DefaultParams()

	targetSupply := sdk.NewIntWithDecimal(100_000_000_000, 18)

	return &GenesisState{
		Params:            params,
		Inflation:         sdk.NewDec(0),
		Era:               uint64(0),
		EraStartedAtBlock: uint64(0),
		EraTargetMint:     sdk.NewCoin(params.MintDenom, sdk.NewInt(0)),
		EraTargetSupply:   sdk.NewCoin(params.MintDenom, sdk.NewInt(0)),
		TargetTotalSupply: sdk.NewCoin(params.MintDenom, targetSupply),
	}
}

// Validate genesis state
func (gs GenesisState) Validate() error {
	if err := validateInflationRate(gs.Inflation); err != nil {
		return err
	}

	if err := validateEraNumber(gs.Era); err != nil {
		return err
	}

	if err := validateEraStartedAtBlock(gs.EraStartedAtBlock); err != nil {
		return err
	}

	if err := validateEraTargetMint(gs.EraTargetMint); err != nil {
		return err
	}

	if err := validateEraTargetSupply(gs.EraTargetSupply); err != nil {
		return err
	}

	if err := validateTargetTotalSupply(gs.TargetTotalSupply); err != nil {
		return err
	}

	return gs.Params.Validate()
}

func validateInflationRate(i interface{}) error {
	_, ok := i.(sdk.Dec)

	if !ok {
		return fmt.Errorf("inflation rate: invalid genesis state type: %T", i)
	}

	return nil
}

func validateEraNumber(i interface{}) error {
	_, ok := i.(uint64)

	if !ok {
		return fmt.Errorf("era number: invalid genesis state type: %T", i)
	}

	return nil
}

func validateEraStartedAtBlock(i interface{}) error {
	_, ok := i.(uint64)

	if !ok {
		return fmt.Errorf("start era block: invalid genesis state type: %T", i)
	}

	return nil
}

func validateEraTargetMint(i interface{}) error {
	_, ok := i.(sdk.Coin)

	if !ok {
		return fmt.Errorf("era mint: invalid genesis state type: %T", i)
	}

	return nil
}

func validateEraTargetSupply(i interface{}) error {
	_, ok := i.(sdk.Coin)

	if !ok {
		return fmt.Errorf("era target supply: invalid genesis state type: %T", i)
	}

	return nil
}

func validateTargetTotalSupply(i interface{}) error {
	_, ok := i.(sdk.Coin)

	if !ok {
		return fmt.Errorf("target total supply: invalid genesis state type: %T", i)
	}

	return nil
}
